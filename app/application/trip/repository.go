package trip

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/route"
	"github.com/SeyramWood/bookibus/ent/trip"
	"github.com/SeyramWood/bookibus/ent/vehicle"
)

const (
	Exterior           string = "exterior"
	Interior           string = "interior"
	EngineCompartment  string = "engine-compartment"
	BrakeAndSteering   string = "brake-and-steering"
	EmergencyEquipment string = "emergency-equipment"
	FuelAndFluid       string = "fuel-and-fluid"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.TripRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// Delete implements gateways.TripRepo.
func (r *repository) Delete(id int) error {
	return r.db.Trip.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.TripRepo.
func (r *repository) Insert(companyId int, request *requeststructs.TripRequest) (*ent.Trip, error) {
	if request.ReturnDate != "" {
		result, err := r.db.Trip.Create().
			SetDepartureDate(application.ParseRFC3339Datetime(request.DepartureDate)).
			SetArrivalDate(application.ParseRFC3339Datetime(request.ArrivalDate)).
			SetReturnDate(application.ParseRFC3339Datetime(request.ReturnDate)).
			SetType(trip.Type(request.TripType)).
			SetScheduled(request.Schedule).
			SetSeatLeft(r.db.Vehicle.GetX(r.ctx, request.VehicleID).Seat).
			SetRate(request.Rate).
			SetDiscount(request.Discount).
			SetFromTerminalID(request.FromTerminalID).
			SetToTerminalID(request.ToTerminalID).
			SetVehicleID(request.VehicleID).
			SetRouteID(request.RouteID).
			AddStopIDs(request.Stops...).
			SetDriverID(request.DriverID).
			SetCompanyID(companyId).
			Save(r.ctx)
		if err != nil {
			return nil, err
		}
		return r.Read(result.ID)
	}
	result, err := r.db.Trip.Create().
		SetDepartureDate(application.ParseRFC3339Datetime(request.DepartureDate)).
		SetArrivalDate(application.ParseRFC3339Datetime(request.ArrivalDate)).
		SetType(trip.Type(request.TripType)).
		SetScheduled(request.Schedule).
		SetSeatLeft(r.db.Vehicle.GetX(r.ctx, request.VehicleID).Seat).
		SetRate(request.Rate).
		SetDiscount(request.Discount).
		SetFromTerminalID(request.FromTerminalID).
		SetToTerminalID(request.ToTerminalID).
		SetVehicleID(request.VehicleID).
		SetRouteID(request.RouteID).
		AddStopIDs(request.Stops...).
		SetDriverID(request.DriverID).
		SetCompanyID(companyId).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(result.ID)
}

// Read implements gateways.TripRepo.
func (r *repository) Read(id int) (*ent.Trip, error) {
	result, err := r.db.Trip.Query().Where(trip.ID(id)).
		WithFromTerminal().
		WithToTerminal().
		WithVehicle(func(vq *ent.VehicleQuery) {
			vq.WithImages()
		}).
		WithRoute().
		WithStops().
		WithDriver().
		WithCompany().
		WithBookings(func(bq *ent.BookingQuery) {
			bq.WithPassengers()
			bq.WithLuggages()
			bq.WithContact()
			bq.WithCustomer(func(cq *ent.CustomerQuery) {
				cq.WithProfile()
			})
			bq.WithTrip(func(tq *ent.TripQuery) {
				tq.WithVehicle(func(vq *ent.VehicleQuery) {
					vq.WithImages()
				})
				tq.WithRoute()
				tq.WithStops()
				tq.WithDriver()
				tq.WithCompany()
			})
		}).
		WithParcels().
		WithIncidents().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.TripRepo.
func (r *repository) ReadAll(limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) && r.compareFilter(fm[com[9]]) {
				query := r.db.Trip.Query().Where(trip.And(r.filterPredicate(fm, com)...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}

		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-5) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-6) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-7) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-8) && r.compareFilter(fm[com[1]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 9) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
	}
	if filter.From == "" && filter.To == "" && filter.Datetime == "" && !filter.Today && !filter.Scheduled && !filter.Completed {
		return r.filterTrip(r.db.Trip.Query(), limit, offset)
	}
	return application.Paginate(0, []*ent.Trip{})
}

// ReadAllSearch implements gateways.TripRepo.
func (r *repository) ReadAllSearch(searchKey string, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	searchPredicate := trip.Or(
		trip.HasBookingsWith(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) LIKE ?", booking.FieldBookingNumber), "%"+strings.ToLower(searchKey)+"%"))
		}),
		trip.HasVehicleWith(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) LIKE ?", vehicle.FieldRegistrationNumber), "%"+strings.ToLower(searchKey)+"%"))
		}),
		trip.HasDriverWith(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) LIKE ? OR LOWER(%s) LIKE ?", companyuser.FieldOtherName, companyuser.FieldLastName), "%"+strings.ToLower(searchKey)+"%", "%"+strings.ToLower(searchKey)+"%"))
		}),
	)
	if filter.From == "" && filter.To == "" && filter.Datetime == "" && !filter.Today && !filter.Scheduled && !filter.Completed {
		return r.filterTrip(r.db.Trip.Query().Where(searchPredicate), limit, offset)
	}
	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) && r.compareFilter(fm[com[9]]) {
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(r.filterPredicate(fm, com)...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}

		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-5) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-6) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-7) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-8) && r.compareFilter(fm[com[1]]) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 9) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					searchPredicate, trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTrip(query, limit, offset)
			}
		}
	}
	return application.Paginate(0, []*ent.Trip{})
}

// ReadAllByCompany implements gateways.TripRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	if filter.From == "" && filter.To == "" && filter.Datetime == "" && filter.ReturnDate == "" && filter.TripType == "" && !filter.Today && !filter.Scheduled && !filter.Completed {
		return r.filterTrip(r.db.Trip.Query().Where(trip.And(trip.HasCompanyWith(company.ID(companyId)))), limit, offset)
	}

	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) && r.compareFilter(fm[com[9]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 5) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 6) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 7) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 8) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 9) {
			if r.compareFilter(fm[com[0]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
	}
	return application.Paginate(0, []*ent.Trip{})
}

// ReadAllSearchByCompany implements gateways.TripRepo.
func (r *repository) ReadAllSearchByCompany(searchKey string, companyId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	searchPredicate := trip.Or(
		trip.HasBookingsWith(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) LIKE ?", booking.FieldBookingNumber), "%"+strings.ToLower(searchKey)+"%"))
		}),
		trip.HasVehicleWith(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) LIKE ?", vehicle.FieldRegistrationNumber), "%"+strings.ToLower(searchKey)+"%"))
		}),
		trip.HasDriverWith(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) LIKE ? OR LOWER(%s) LIKE ?", companyuser.FieldOtherName, companyuser.FieldLastName), "%"+strings.ToLower(searchKey)+"%", "%"+strings.ToLower(searchKey)+"%"))
		}),
	)
	if filter.From == "" && filter.To == "" && filter.Datetime == "" && filter.ReturnDate == "" && filter.TripType == "" && !filter.Today && !filter.Scheduled && !filter.Completed {
		return r.filterTrip(r.db.Trip.Query().Where(searchPredicate, trip.And(trip.HasCompanyWith(company.ID(companyId)))), limit, offset)
	}
	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) && r.compareFilter(fm[com[9]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 5) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 6) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 7) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 8) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 9) {
			if r.compareFilter(fm[com[0]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasCompanyWith(company.ID(companyId)))
				query := r.db.Trip.Query().Where(searchPredicate, trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
	}
	return application.Paginate(0, []*ent.Trip{})
}

// ReadAllByDriver implements gateways.TripRepo.
func (r *repository) ReadAllByDriver(driverId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	fm := application.ConvertStructToMap(*(filter))
	driverID := r.db.User.GetX(r.ctx, driverId).QueryCompanyUser().OnlyIDX(r.ctx)
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) && r.compareFilter(fm[com[9]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) && r.compareFilter(fm[com[8]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) && r.compareFilter(fm[com[7]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}

		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-5) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) {
			if r.compareFilter(fm[com[0]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-6) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) {
			if r.compareFilter(fm[com[0]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-7) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) {
			if r.compareFilter(fm[com[0]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys())-8) && r.compareFilter(fm[com[1]]) {
			if r.compareFilter(fm[com[0]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 9) {
			if r.compareFilter(fm[com[0]]) {
				predicates := r.filterPredicate(fm, com)
				predicates = append(predicates, trip.HasDriverWith(companyuser.ID(driverID)))
				query := r.db.Trip.Query().Where(trip.And(predicates...))
				return r.filterTrip(query, limit, offset)
			}
		}
	}
	return application.Paginate(0, []*ent.Trip{})
}

// ReadAllCustomer implements gateways.TripRepo.
func (r *repository) ReadAllCustomer(limit int, offset int, filter *requeststructs.CustomerTripFilterRequest) (*presenters.PaginationResponse, error) {
	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.customerFilterKeys()) {
		if len(com) == len(r.customerFilterKeys()) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.customerFilterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.customerFilterKeys()) - 1) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.customerFilterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.customerFilterKeys()) - 2) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.customerFilterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.customerFilterKeys()) - 3) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.customerFilterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.customerFilterKeys()) - 4) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.customerFilterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.customerFilterKeys()) - 5) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.customerFilterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.customerFilterKeys()) - 6) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.customerFilterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
	}
	return r.filterTripByPopularity(r.db.Trip.Query(), limit, offset)
}

// ReadAllPopular implements gateways.TripRepo.
func (r *repository) ReadAllPopular(limit int, offset int) (*presenters.PaginationResponse, error) {
	return r.filterTripByPopularity(r.db.Trip.Query(), limit, offset)
}

// Update implements gateways.TripRepo.
func (r *repository) Update(id int, request *requeststructs.TripUpdateRequest) (*ent.Trip, error) {
	if request.ReturnDate != "" {
		result, err := r.db.Trip.UpdateOneID(id).
			SetDepartureDate(application.ParseRFC3339Datetime(request.DepartureDate)).
			SetArrivalDate(application.ParseRFC3339Datetime(request.ArrivalDate)).
			SetReturnDate(application.ParseRFC3339Datetime(request.ReturnDate)).
			SetType(trip.Type(request.TripType)).
			SetRate(request.Rate).
			SetDiscount(request.Discount).
			SetVehicleID(request.VehicleID).
			SetRouteID(request.RouteID).
			SetFromTerminalID(request.FromTerminalID).
			SetToTerminalID(request.ToTerminalID).
			ClearStops().
			AddStopIDs(request.Stops...).
			SetDriverID(request.DriverID).
			Save(r.ctx)
		if err != nil {
			return nil, err
		}
		return r.Read(result.ID)
	}
	result, err := r.db.Trip.UpdateOneID(id).
		SetDepartureDate(application.ParseRFC3339Datetime(request.DepartureDate)).
		SetArrivalDate(application.ParseRFC3339Datetime(request.ArrivalDate)).
		SetRate(request.Rate).
		SetDiscount(request.Discount).
		ClearReturnDate().
		SetType(trip.Type(request.TripType)).
		SetVehicleID(request.VehicleID).
		SetRouteID(request.RouteID).
		SetFromTerminalID(request.FromTerminalID).
		SetToTerminalID(request.ToTerminalID).
		ClearStops().
		AddStopIDs(request.Stops...).
		SetDriverID(request.DriverID).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(result.ID)

}

// UpdateInspection implements gateways.TripRepo.
func (r *repository) UpdateInspection(id int, request *requeststructs.TripInspectionStatusRequest) (*ent.Trip, error) {
	result, err := r.db.Trip.UpdateOneID(id).
		SetExteriorInspected(request.Exterior).
		SetInteriorInspected(request.Interior).
		SetEngineCompartmentInspected(request.EngineCompartment).
		SetBrakeAndSteeringInspected(request.BrakeAndSteering).
		SetEmergencyEquipmentInspected(request.EmergencyEquipment).
		SetFuelAndFluidsInspected(request.FuelAndFluid).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(result.ID)
}

// UpdateSchedule implements gateways.TripRepo.
func (r *repository) UpdateSchedule(id int, status bool) (*ent.Trip, error) {
	result, err := r.db.Trip.UpdateOneID(id).
		SetScheduled(status).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(result.ID)
}

// UpdateStatus implements gateways.TripRepo.
func (r *repository) UpdateStatus(id int, status string) (*ent.Trip, error) {
	t := r.db.Trip.GetX(r.ctx, id)
	if trip.Status(status) == trip.StatusStarted && !t.Scheduled {
		return nil, fmt.Errorf("can not start trip")
	}
	if trip.Status(status) == trip.StatusEnded && !t.Scheduled {
		return nil, fmt.Errorf("can not end trip")
	}
	result, err := r.db.Trip.UpdateOneID(id).
		SetStatus(trip.Status(status)).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	go func(r *repository, t *ent.Trip) {
		if result.Status == "ended" {
			result, err := r.Read(result.ID)
			if err != nil {
				return
			}
			_ = r.db.Route.UpdateOne(result.Edges.Route).AddPopularity(1).SaveX(r.ctx)
		}
	}(r, result)
	return r.Read(result.ID)
}

func (r *repository) filterTrip(query *ent.TripQuery, limit, offset int) (*presenters.PaginationResponse, error) {

	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Asc(trip.FieldDepartureDate)).
		WithFromTerminal().
		WithToTerminal().
		WithVehicle(func(vq *ent.VehicleQuery) {
			vq.WithImages()
		}).
		WithRoute().
		WithStops().
		WithDriver().
		WithCompany().
		WithBookings(func(bq *ent.BookingQuery) {
			bq.WithPassengers()
			bq.WithLuggages()
			bq.WithContact()
			bq.WithCustomer(func(cq *ent.CustomerQuery) {
				cq.WithProfile()
			})
		}).
		WithParcels(func(pq *ent.ParcelQuery) {
			pq.WithImages()
		}).
		WithIncidents(func(iq *ent.IncidentQuery) {
			iq.WithImages()
		}).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

func (r *repository) filterTripByPopularity(query *ent.TripQuery, limit, offset int) (*presenters.PaginationResponse, error) {
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(
			trip.ByRouteField(
				route.FieldPopularity,
				sql.OrderDesc(),
			),
		).
		WithFromTerminal().
		WithToTerminal().
		WithVehicle(func(vq *ent.VehicleQuery) {
			vq.WithImages()
		}).
		WithRoute().
		WithStops().
		WithDriver().
		WithCompany().
		WithBookings(func(bq *ent.BookingQuery) {
			bq.WithPassengers()
			bq.WithLuggages()
			bq.WithContact()
			bq.WithCustomer(func(cq *ent.CustomerQuery) {
				cq.WithProfile()
			})
		}).
		WithParcels(func(pq *ent.ParcelQuery) {
			pq.WithImages()
		}).
		WithIncidents(func(iq *ent.IncidentQuery) {
			iq.WithImages()
		}).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

func (r *repository) compareFilter(value any) bool {
	switch value.(type) {
	case int:
		if value != 0 {
			return true
		}
	case string:
		if value != "" {
			return true
		}
	case bool:
		if value == true {
			return true
		}
	}
	return false
}
func (r *repository) customerFilterKeys() []string {
	return []string{"CompanyID", "TripType", "From", "To", "DepartureDate", "ReturnDate", "Passengers"}
}
func (r *repository) customerFilterPredicate(data map[string]any, combinations []string) []predicate.Trip {
	results := make([]predicate.Trip, 0, len(combinations))
	for _, combination := range combinations {
		for k, v := range data {
			if combination == k && combination == "CompanyID" {
				results = append(results, trip.HasCompanyWith(company.ID(v.(int))))
				break
			}
			if combination == k && combination == "TripType" {
				results = append(results, trip.TypeEQ(trip.Type(v.(string))))
				break
			}
			if combination == k && combination == "From" {
				results = append(results, trip.HasRouteWith(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) = ?", route.FieldFromLocation), strings.ToLower(v.(string))))
					}))
				break
			}
			if combination == k && combination == "To" {
				results = append(results, trip.HasRouteWith(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) = ?", route.FieldToLocation), strings.ToLower(v.(string))))
					}))
				break
			}
			if combination == k && combination == "DepartureDate" {
				results = append(results, trip.And(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", trip.FieldDepartureDate), application.ParseRFC3339MYSQLDatetime(v.(string))))
					},
				))
				break
			}
			if combination == k && combination == "ReturnDate" {
				results = append(results, trip.And(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", trip.FieldReturnDate), application.ParseRFC3339MYSQLDatetime(v.(string))))
					},
				))
				break
			}
			if combination == k && combination == "Passengers" {
				results = append(results, trip.SeatLeftGTE(v.(int)))
				break
			}
		}
	}
	return results
}

func (r *repository) filterKeys() []string {
	return []string{"From", "To", "Datetime", "ReturnDate", "TripType", "Today", "Scheduled", "Completed", "Passengers", "TimeRange"}
}
func (r *repository) filterPredicate(data map[string]any, combinations []string) []predicate.Trip {
	results := make([]predicate.Trip, 0, len(combinations))
	for _, combination := range combinations {
		for k, v := range data {
			if combination == k && combination == "From" {
				results = append(results, trip.HasRouteWith(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) = ?", route.FieldFromLocation), strings.ToLower(v.(string))))
					}))
				break
			}
			if combination == k && combination == "To" {
				results = append(results, trip.HasRouteWith(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("LOWER(%s) = ?", route.FieldToLocation), strings.ToLower(v.(string))))
					}))
				break
			}
			if combination == k && combination == "Datetime" {
				results = append(results, trip.And(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", trip.FieldDepartureDate), application.ParseRFC3339MYSQLDatetime(v.(string))))
					},
				))

				break
			}
			if combination == k && combination == "ReturnDate" {
				results = append(results, trip.And(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", trip.FieldReturnDate), application.ParseRFC3339MYSQLDatetime(v.(string))))
					},
				))
				break
			}
			if combination == k && combination == "TripType" {
				results = append(results, trip.TypeEQ(trip.Type(v.(string))))
				break
			}
			if combination == k && combination == "Today" {
				results = append(results, trip.And(
					trip.Scheduled(false),
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
					},
				))
				break
			}
			if combination == k && combination == "Scheduled" {
				results = append(results, trip.And(
					trip.Or(
						trip.StatusNotIn(trip.StatusEnded),
						func(s *sql.Selector) {
							s.Where(sql.ExprP(fmt.Sprintf("%s IS NULL", trip.FieldStatus)))
						},
					),
					trip.Scheduled(true),
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
					},
				))
				break
			}
			if combination == k && combination == "Completed" {
				results = append(results, trip.And(
					trip.Scheduled(true),
					trip.StatusEQ(trip.StatusEnded),
				))
				break
			}
			if combination == k && combination == "Passengers" {
				results = append(results, trip.SeatLeftGTE(v.(int)))
				break
			}
			if combination == k && combination == "TimeRange" {
				timeRage := strings.SplitN(v.(string), "_", 2)
				results = append(results, trip.And(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(fmt.Sprintf("TIME(%s) BETWEEN ? AND ?", trip.FieldDepartureDate), timeRage[0], timeRage[1]))
					},
				))
				break
			}
		}
	}
	return results
}
