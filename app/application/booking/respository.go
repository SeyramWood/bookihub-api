package booking

import (
	"fmt"
	"log"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/net/context"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/domain"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/configuration"
	"github.com/SeyramWood/bookibus/ent/customer"
	"github.com/SeyramWood/bookibus/ent/customercontact"
	"github.com/SeyramWood/bookibus/ent/customerluggage"
	"github.com/SeyramWood/bookibus/ent/passenger"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/schema"
	"github.com/SeyramWood/bookibus/ent/transaction"
	"github.com/SeyramWood/bookibus/ent/trip"
	"github.com/SeyramWood/bookibus/ent/user"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.BookingRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// CancelBooking implements gateways.BookingRepo.
func (r *repository) CancelBooking(id int, request *requeststructs.BookingCancelRequest) (*ent.Booking, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	_, err = tx.Booking.UpdateOneID(id).SetStatus(booking.StatusCanceled).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed updating booking: %w", err))
	}
	_, err = tx.Transaction.Update().
		Where(transaction.HasBookingWith(booking.ID(id))).
		SetCancellationFee(domain.ComputeCancellationFee(r.readCharge().TripCancellationFee, request.Amount)).
		SetCanceledAt(time.Now()).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed updating transaction: %w", err))
	}
	return r.Read(id)
}

// Delete implements gateways.BookingRepo.
func (r *repository) Delete(id int) error {
	return r.db.Booking.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.BookingRepo.
func (r *repository) Insert(request *requeststructs.BookingRequest, refResponse *requeststructs.PaymentReferenceResponse) (*ent.Booking, error) {
	seatLeft := r.db.Trip.GetX(r.ctx, request.TripID).SeatLeft
	if seatLeft == 0 {
		return nil, fmt.Errorf("no available seats")
	}
	if len(request.Passenger) > seatLeft {
		return nil, fmt.Errorf("only %d seats available", seatLeft)
	}
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	var result *ent.Booking
	if request.CustomerID != 0 {

		res, err := tx.Booking.Create().
			SetBookingNumber(domain.ComputeUniqueCode("T")).
			SetSmsNotification(request.SMSNotification).
			SetStatus(booking.StatusSuccessful).
			SetTripID(request.TripID).
			SetCompanyID(request.CompanyID).
			SetCustomerID(r.db.User.Query().Where(user.ID(request.CustomerID)).QueryCustomer().OnlyIDX(r.ctx)).
			Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating booking: %w", err))
		}
		result = res
	} else {
		res, err := tx.Booking.Create().
			SetBookingNumber(domain.ComputeUniqueCode("T")).
			SetSmsNotification(request.SMSNotification).
			SetStatus(booking.StatusSuccessful).
			SetTripID(request.TripID).
			SetCompanyID(request.CompanyID).
			Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating booking: %w", err))
		}
		_, err = tx.CustomerContact.Create().
			SetFullName(request.Contact.FullName).
			SetPhone(request.Contact.Phone).
			SetEmail(request.Contact.Email).
			SetBooking(res).
			Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating customer contact: %w", err))
		}
		result = res
	}
	_, err = tx.Passenger.MapCreateBulk(request.Passenger, func(create *ent.PassengerCreate, i int) {
		create.SetFullName(request.Passenger[i].FullName).
			SetAmount(request.Passenger[i].Amount).
			SetMaturity(passenger.Maturity(request.Passenger[i].Maturity)).
			SetGender(passenger.Gender(request.Passenger[i].Gender)).
			SetBooking(result)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating passenger: %w", err))
	}
	if len(request.Luggage) > 0 {
		_, err = tx.CustomerLuggage.MapCreateBulk(request.Luggage, func(create *ent.CustomerLuggageCreate, i int) {
			create.
				SetBaggage(customerluggage.Baggage(request.Luggage[i].Baggage)).
				SetQuantity(request.Luggage[i].Quantity).
				SetAmount(request.Luggage[i].Amount).
				SetBooking(result)
		}).Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating customer luggage: %w", err))
		}
	}
	_, err = tx.Transaction.Create().
		SetReference(request.Reference).
		SetAmount(request.Amount).
		SetVat(request.VAT).
		SetTransactionFee(domain.ComputeTransactionFee(r.readCharge().PaymentGatewayServiceFee, r.readCharge().TripServiceFee, request.Amount)).
		SetPaidAt(application.ParseRFC3339Datetime(refResponse.PaidAt)).
		SetChannel(transaction.Channel(refResponse.TransType)).
		SetBooking(result).
		SetCompanyID(request.CompanyID).
		SetProduct(transaction.ProductTrip).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating transaction: %w", err))
	}
	_, err = tx.Trip.UpdateOneID(request.TripID).SetSeatLeft(seatLeft - len(request.Passenger)).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed updating booking on seat left update: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing booking creation transaction: %w", err)
	}
	return r.Read(result.ID)
}

// Read implements gateways.BookingRepo.
func (r *repository) Read(id int) (*ent.Booking, error) {
	result, err := r.db.Booking.Query().Where(booking.ID(id)).
		WithPassengers().
		WithLuggages().
		WithContact().
		WithCustomer(func(cq *ent.CustomerQuery) {
			cq.WithProfile()
		}).
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute()
			tq.WithStops()
			tq.WithDriver()
			tq.WithCompany()
			tq.WithFromTerminal()
			tq.WithToTerminal()
		}).
		WithTransaction().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadByBookingNumber implements gateways.BookingRepo.
func (r *repository) ReadByBookingNumber(id string) (*ent.Booking, error) {
	result, err := r.db.Booking.Query().Where(booking.BookingNumber(id)).
		WithPassengers().
		WithLuggages().
		WithContact().
		WithCustomer(func(cq *ent.CustomerQuery) {
			cq.WithProfile()
		}).
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute()
			tq.WithStops()
			tq.WithDriver()
			tq.WithCompany()
			tq.WithFromTerminal()
			tq.WithToTerminal()
		}).
		WithTransaction().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.BookingRepo.
func (r *repository) ReadAll(limit int, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error) {
	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) && application.CompareFilter(fm[com[4]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if application.CompareFilter(fm[com[0]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
	}

	return r.filterBooking(r.db.Booking.Query(), limit, offset)
}

// ReadAllByCompany implements gateways.BookingRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error) {
	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) && application.CompareFilter(fm[com[4]]) {
				query := r.db.Booking.Query().Where(
					booking.HasCompanyWith(company.ID(companyId)),
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) {
				query := r.db.Booking.Query().Where(
					booking.HasCompanyWith(company.ID(companyId)),
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) {
				query := r.db.Booking.Query().Where(
					booking.HasCompanyWith(company.ID(companyId)),
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) {
				query := r.db.Booking.Query().Where(
					booking.HasCompanyWith(company.ID(companyId)),
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if application.CompareFilter(fm[com[0]]) {
				query := r.db.Booking.Query().Where(
					booking.HasCompanyWith(company.ID(companyId)),
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
	}

	return r.filterBooking(r.db.Booking.Query(), limit, offset)
}

// ReadAllCustomer implements gateways.BookingRepo.
func (r *repository) ReadAllCustomer(limit int, offset int, filter *requeststructs.BookingFilterRequest, customerId ...int) (*presenters.PaginationResponse, error) {
	log.Println(filter)
	fm := application.ConvertStructToMap(*(filter))
	if customerId != nil && customerId[0] != 0 {
		for _, com := range application.FilterCombinations(r.filterKeys()) {
			if len(com) == len(r.filterKeys()) {
				if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) && application.CompareFilter(fm[com[4]]) {
					query := r.db.Booking.Query().Where(
						booking.And(r.filterPredicate(fm, com)...),
					)
					return r.filterBooking(query, limit, offset)
				}
			}
			if len(com) == (len(r.filterKeys()) - 1) {
				if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) {
					query := r.db.Booking.Query().Where(
						booking.And(r.filterPredicate(fm, com)...),
					)
					return r.filterBooking(query, limit, offset)
				}
			}
			if len(com) == (len(r.filterKeys()) - 2) {
				if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) {
					query := r.db.Booking.Query().Where(
						booking.And(r.filterPredicate(fm, com)...),
					)
					return r.filterBooking(query, limit, offset)
				}
			}
			if len(com) == (len(r.filterKeys()) - 3) {
				if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) {
					query := r.db.Booking.Query().Where(
						booking.And(r.filterPredicate(fm, com)...),
					)
					return r.filterBooking(query, limit, offset)
				}
			}
			if len(com) == (len(r.filterKeys()) - 4) {
				if application.CompareFilter(fm[com[0]]) {
					query := r.db.Booking.Query().Where(
						booking.And(r.filterPredicate(fm, com)...),
					)
					return r.filterBooking(query, limit, offset)
				}
			}
		}
		query := r.db.Booking.Query().Where(
			booking.HasCustomerWith(customer.ID(r.db.User.Query().Where(user.ID(customerId[0])).QueryCustomer().OnlyIDX(r.ctx))),
		)
		return r.filterBooking(query, limit, offset)
	}
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) && application.CompareFilter(fm[com[4]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) && application.CompareFilter(fm[com[3]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) && application.CompareFilter(fm[com[2]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if application.CompareFilter(fm[com[0]]) && application.CompareFilter(fm[com[1]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if application.CompareFilter(fm[com[0]]) {
				query := r.db.Booking.Query().Where(
					booking.And(r.filterPredicate(fm, com)...),
				)
				return r.filterBooking(query, limit, offset)
			}
		}
	}

	return application.Paginate(0, []*ent.Booking{})
}

// Update implements gateways.BookingRepo.
func (r *repository) Update(id int, request *requeststructs.BookingUpdateRequest) (*ent.Booking, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	result, err := tx.Booking.UpdateOneID(id).
		SetSmsNotification(request.SMSNotification).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed update booking: %w", err))
	}
	_, err = tx.Transaction.UpdateOneID(id).
		SetAmount(request.Amount).
		SetChannel(transaction.Channel(request.TransactionType)).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed update booking: %w", err))
	}
	if cId, err := result.QueryContact().OnlyID(r.ctx); err == nil {
		_, err = tx.CustomerContact.UpdateOneID(cId).
			SetFullName(request.Contact.FullName).
			SetPhone(request.Contact.Phone).
			SetEmail(request.Contact.Email).
			Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed update customer contact: %w", err))
		}
	}
	for _, p := range request.Passenger {
		_, err = tx.Passenger.UpdateOneID(p.ID).
			SetFullName(p.FullName).
			SetAmount(p.Amount).
			SetMaturity(passenger.Maturity(p.Maturity)).
			SetGender(passenger.Gender(p.Gender)).
			Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed updating customer passenger: %w", err))
		}
	}
	if len(request.Luggage) > 0 {
		for _, l := range request.Luggage {
			_, err = tx.CustomerLuggage.UpdateOneID(l.ID).
				SetBaggage(customerluggage.Baggage(l.Baggage)).
				SetQuantity(l.Quantity).
				SetAmount(l.Amount).
				Save(r.ctx)
			if err != nil {
				return nil, application.Rollback(tx, fmt.Errorf("failed update customer luggage: %w", err))
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing booking update transaction: %w", err)
	}

	return r.Read(result.ID)
}

func (r *repository) readCharge() *schema.Charge {
	return r.db.Configuration.Query().Select(configuration.FieldCharge).AllX(r.ctx)[0].Charge
}

func (r *repository) filterBooking(query *ent.BookingQuery, limit, offset int) (*presenters.PaginationResponse, error) {
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(booking.FieldCreatedAt)).
		WithPassengers().
		WithLuggages().
		WithContact().
		WithCustomer(func(cq *ent.CustomerQuery) {
			cq.WithProfile()
		}).
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute()
			tq.WithStops()
			tq.WithDriver()
			tq.WithCompany()
			tq.WithFromTerminal()
			tq.WithToTerminal()
		}).
		WithTransaction().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

func (r *repository) filterKeys() []string {
	return []string{"BookingNumber", "Phone", "Active", "Completed", "Canceled"}
}
func (r *repository) filterPredicate(data map[string]any, combinations []string) []predicate.Booking {
	results := make([]predicate.Booking, 0, len(combinations))
	for _, combination := range combinations {
		for k, v := range data {
			if combination == k && combination == "BookingNumber" {
				results = append(results, booking.BookingNumberEQ(v.(string)))
				break
			}
			if combination == k && combination == "Phone" {
				results = append(results, booking.Or(
					booking.HasCustomerWith(customer.PhoneEQ(strings.Replace("+"+v.(string), " ", "", -1))),
					booking.HasContactWith(customercontact.PhoneEQ(strings.Replace("+"+v.(string), " ", "", -1))),
				))
				break
			}
			if combination == k && combination == "Active" {
				results = append(results, booking.HasTripWith(func(s *sql.Selector) {
					s.Where(sql.ExprP(fmt.Sprintf("%s = ? OR %s IS NULL", trip.FieldStatus, trip.FieldStatus), trip.StatusStarted))
				}))
				break
			}
			if combination == k && combination == "Completed" {
				results = append(results, booking.And(
					booking.StatusEQ(booking.StatusSuccessful),
					booking.HasTripWith(trip.StatusEQ(trip.StatusEnded)),
				))
				break
			}
			if combination == k && combination == "Canceled" {
				results = append(results, booking.StatusEQ(booking.StatusCanceled))
				break
			}
		}
	}
	return results
}
