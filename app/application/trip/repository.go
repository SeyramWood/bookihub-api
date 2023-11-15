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
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/route"
	"github.com/SeyramWood/bookibus/ent/trip"
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
			SetFromTerminalID(request.FromTerminalID).
			SetToTerminalID(request.ToTerminalID).
			SetVehicleID(request.VehicleID).
			SetRouteID(request.RouteID).
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
		SetFromTerminalID(request.FromTerminalID).
		SetToTerminalID(request.ToTerminalID).
		SetVehicleID(request.VehicleID).
		SetRouteID(request.RouteID).
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
		WithRoute(func(rq *ent.RouteQuery) {
			rq.WithStops()
		}).
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
				tq.WithRoute(func(rq *ent.RouteQuery) {
					rq.WithStops()
				})
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
	if filter.Today {
		query := r.db.Trip.Query().Where(trip.And(
			trip.Scheduled(false),
			func(s *sql.Selector) {
				s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
			},
		))

		return r.filterTrip(query, limit, offset)
	}
	if filter.Scheduled {
		query := r.db.Trip.Query().Where(trip.And(
			trip.Scheduled(true),
			func(s *sql.Selector) {
				s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
			},
		))
		return r.filterTrip(query, limit, offset)
	}
	if filter.Completed {
		query := r.db.Trip.Query().Where(trip.And(
			trip.StatusEQ(trip.StatusEnded),
		))
		return r.filterTrip(query, limit, offset)
	}

	return r.filterTrip(r.db.Trip.Query(), limit, offset)
}

// ReadAllByCompany implements gateways.TripRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {

	if filter.Today {
		query := r.db.Trip.Query().Where(trip.And(
			trip.HasCompanyWith(company.ID(companyId)),
			trip.Scheduled(false),
			func(s *sql.Selector) {
				s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
			},
		))

		return r.filterTrip(query, limit, offset)
	}
	if filter.Scheduled {
		query := r.db.Trip.Query().Where(trip.And(
			trip.HasCompanyWith(company.ID(companyId)),
			trip.Scheduled(true),
			func(s *sql.Selector) {
				s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
			},
		))
		return r.filterTrip(query, limit, offset)
	}
	if filter.Completed {
		query := r.db.Trip.Query().Where(trip.And(
			trip.HasCompanyWith(company.ID(companyId)),
			trip.StatusEQ(trip.StatusEnded),
		))
		return r.filterTrip(query, limit, offset)
	}

	return r.filterTrip(r.db.Trip.Query().Where(trip.HasCompanyWith(company.ID(companyId))), limit, offset)
}

// ReadAllByDriver implements gateways.TripRepo.
func (r *repository) ReadAllByDriver(driverId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	if filter.Today {
		query := r.db.Trip.Query().Where(trip.And(
			trip.HasDriverWith(companyuser.ID(driverId)),
			trip.Scheduled(false),
			func(s *sql.Selector) {
				s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
			},
		))
		return r.filterTrip(query, limit, offset)
	}
	if filter.Scheduled {
		query := r.db.Trip.Query().Where(trip.And(
			trip.HasDriverWith(companyuser.ID(driverId)),
			trip.Scheduled(true),
			func(s *sql.Selector) {
				s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = CURDATE()", trip.FieldDepartureDate)))
			},
		))
		return r.filterTrip(query, limit, offset)
	}
	query := r.db.Trip.Query().Where(trip.And(
		trip.HasDriverWith(companyuser.ID(driverId)),
		trip.StatusEQ(trip.StatusEnded),
	))
	return r.filterTrip(query, limit, offset)
}

// ReadAllCustomer implements gateways.TripRepo.
func (r *repository) ReadAllCustomer(limit int, offset int, filter *requeststructs.CustomerTripFilterRequest) (*presenters.PaginationResponse, error) {
	fm := application.ConvertStructToMap(*(filter))
	for _, com := range application.FilterCombinations(r.filterKeys()) {
		if len(com) == len(r.filterKeys()) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) && r.compareFilter(fm[com[6]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 1) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) && r.compareFilter(fm[com[5]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 2) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) && r.compareFilter(fm[com[4]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 3) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) && r.compareFilter(fm[com[3]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 4) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) && r.compareFilter(fm[com[2]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 5) {
			if r.compareFilter(fm[com[0]]) && r.compareFilter(fm[com[1]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
				)
				return r.filterTripByPopularity(query, limit, offset)
			}
		}
		if len(com) == (len(r.filterKeys()) - 6) {
			if r.compareFilter(fm[com[0]]) {
				query := r.db.Trip.Query().Where(
					trip.And(r.filterPredicate(fm, com)...),
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
			SetVehicleID(request.VehicleID).
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
		ClearReturnDate().
		SetType(trip.Type(request.TripType)).
		SetVehicleID(request.VehicleID).
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
		Order(ent.Desc(trip.FieldCreatedAt)).
		WithFromTerminal().
		WithToTerminal().
		WithVehicle(func(vq *ent.VehicleQuery) {
			vq.WithImages()
		}).
		WithRoute(func(rq *ent.RouteQuery) {
			rq.WithStops()
		}).
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
			trip.ByRouteField(route.FieldPopularity,
				sql.OrderDesc(),
			),
		).
		WithFromTerminal().
		WithToTerminal().
		WithVehicle(func(vq *ent.VehicleQuery) {
			vq.WithImages()
		}).
		WithRoute(func(rq *ent.RouteQuery) {
			rq.WithStops()
		}).
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
	}
	return false
}
func (r *repository) filterKeys() []string {
	return []string{"CompanyID", "TripType", "From", "To", "DepartureDate", "ReturnDate", "Passengers"}
}
func (r *repository) filterPredicate(data map[string]any, combinations []string) []predicate.Trip {
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
				results = append(results, trip.DepartureDate(application.ParseRFC3339Datetime(v.(string))))
				break
			}
			if combination == k && combination == "ReturnDate" {
				results = append(results, trip.ReturnDate(application.ParseRFC3339Datetime(v.(string))))
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
