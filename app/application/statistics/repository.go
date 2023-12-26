package statistics

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/customer"
	"github.com/SeyramWood/bookibus/ent/customercontact"
	"github.com/SeyramWood/bookibus/ent/incident"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/terminal"
	"github.com/SeyramWood/bookibus/ent/transaction"
	"github.com/SeyramWood/bookibus/ent/trip"
	"github.com/SeyramWood/bookibus/ent/user"
	"github.com/SeyramWood/bookibus/ent/vehicle"
)

var months = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var days = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

type (
	adminOverview struct {
		Month  string  `json:"month"`
		Amount float64 `json:"amount"`
	}
	companyOverview struct {
		Day    int     `json:"day"`
		Amount float64 `json:"amount"`
	}
)
type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.StatisticsRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// ReadAdminBestSelling implements gateways.StatisticsRepo.
func (r *repository) ReadAdminBestSelling(limit int, offset int, minDate string, maxDate string) (*presenters.PaginationResponse, error) {
	var bs []presenters.AdminBestSelling
	err := r.db.Company.Query().
		Limit(limit).
		Offset(offset).
		GroupBy(company.FieldID, company.FieldName).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(transaction.Table)
			s.Join(t).On(s.C(company.FieldID), t.C(company.TransactionsColumn))
			s.Select(
				sql.As(s.C(company.FieldID), "id"),
				sql.As(s.C(company.FieldName), "company"),
				sql.As(t.C(transaction.FieldProduct), "product"),
				sql.As(sql.Sum(t.C(transaction.FieldAmount)), "amount"),
			)
			s.GroupBy(transaction.FieldProduct)
			s.OrderBy(sql.Desc(t.C(transaction.FieldAmount)))
			return ""
		}).
		Scan(r.ctx, &bs)

	if err != nil {
		return nil, err
	}
	return application.Paginate(len(bs), bs)
}

// ReadAdminRevenue implements gateways.StatisticsRepo.
func (r *repository) ReadAdminRevenue() *presenters.AdminRevenue {
	return &presenters.AdminRevenue{
		Total: func() float64 {
			result, err := r.db.Transaction.Query().Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
			if err != nil {
				return 0
			}
			return result
		}(),
		Trip: func() float64 {
			result, err := r.db.Transaction.Query().Where(transaction.ProductIn(transaction.ProductTrip)).Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
			if err != nil {
				return 0
			}
			return result
		}(),
		Delivery: func() float64 {
			result, err := r.db.Transaction.Query().Where(transaction.ProductIn(transaction.ProductDelivery)).Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
			if err != nil {
				return 0
			}
			return result
		}(),
	}
}

// ReadAdminRevenueOverview implements gateways.StatisticsRepo.
func (r *repository) ReadAdminRevenueOverview(filter string) ([]presenters.AdminOverview, error) {
	var overview []adminOverview
	err := r.db.Transaction.Query().Where(r.filterTransactionPredicate(filter)).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As("DATE_FORMAT(created_at, '%b')", "month"),
				sql.As(sql.Sum("amount"), "amount"),
			).
				GroupBy("MONTH(created_at)", "YEAR(created_at)").
				OrderBy(sql.Asc("MONTH(created_at)"))
		}).
		Scan(r.ctx, &overview)
	if err != nil {
		return nil, err
	}
	return r.formatOverview(overview), nil
}

// ReadAdminUserOverview implements gateways.StatisticsRepo.
func (r *repository) ReadAdminUserOverview(filter string) ([]presenters.AdminUserOverview, error) {
	var overview []presenters.AdminUserOverview
	err := r.db.User.Query().Where(r.filterUserPredicate(filter)).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As("DATE_FORMAT(created_at, '%b')", "month"),
				sql.As(sql.Count("*"), "user"),
			).
				GroupBy("MONTH(created_at)", "YEAR(created_at)").
				OrderBy(sql.Asc("MONTH(created_at)"))
		}).
		Scan(r.ctx, &overview)
	if err != nil {
		return nil, err
	}
	return r.formatUserOverview(overview), nil
}

// ReadAdminCompanyOverview implements gateways.StatisticsRepo.
func (r *repository) ReadAdminCompanyOverview(companyId int, filter string) (*presenters.AdminCompanyOverview, error) {
	return &presenters.AdminCompanyOverview{
		Revenue: &presenters.AdminRevenue{
			Total: func() float64 {
				result, err := r.db.Transaction.Query().Where(
					transaction.And(
						transaction.HasCompanyWith(company.ID(companyId)),
						r.filterTransactionPredicate(filter),
					),
				).Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
				if err != nil {
					return 0
				}
				return result
			}(),
			Trip: func() float64 {
				result, err := r.db.Transaction.Query().Where(
					transaction.And(
						transaction.HasCompanyWith(company.ID(companyId)),
						transaction.ProductIn(transaction.ProductTrip),
						r.filterTransactionPredicate(filter),
					),
				).Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
				if err != nil {
					return 0
				}
				return result
			}(),
			Delivery: func() float64 {
				result, err := r.db.Transaction.Query().Where(
					transaction.And(
						transaction.HasCompanyWith(company.ID(companyId)),
						transaction.ProductIn(transaction.ProductDelivery),
						r.filterTransactionPredicate(filter),
					),
				).Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
				if err != nil {
					return 0
				}
				return result
			}(),
		},
		Staff:    r.db.CompanyUser.Query().Where(companyuser.And(companyuser.HasCompanyWith(company.ID(companyId)), r.filterCompanyUserPredicate(filter))).CountX(r.ctx),
		Terminal: r.db.Terminal.Query().Where(terminal.And(terminal.HasCompanyWith(company.ID(companyId)), r.filterTerminalPredicate(filter))).CountX(r.ctx),
		Fleet:    r.db.Vehicle.Query().Where(vehicle.And(vehicle.HasCompanyWith(company.ID(companyId)), r.filterVehiclePredicate(filter))).CountX(r.ctx),
		Customer: func() int {
			rc := r.db.Booking.Query().Where(booking.HasCompanyWith(company.ID(companyId))).QueryCustomer().Where(customer.And(r.filterCustomerPredicate(filter))).CountX(r.ctx)
			contacts := r.db.Booking.Query().Where(booking.HasCompanyWith(company.ID(companyId))).QueryContact().Where(customercontact.And(r.filterCustomerContactPredicate(filter))).Modify(func(s *sql.Selector) {
				s.Select(fmt.Sprintf("DISTINCT %s", customercontact.FieldPhone))
			}).CountX(r.ctx)
			return rc + contacts
		}(),
	}, nil
}

// ReadCompanyIncidentOverview implements gateways.StatisticsRepo.
func (r *repository) ReadCompanyIncidentOverview(companyId int, filter string) (*presenters.CompanyIncidentOverview, error) {
	return &presenters.CompanyIncidentOverview{
		Total: r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				r.filterIncidentPredicate(filter),
			),
		).CountX(r.ctx),
		Pending: r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				incident.StatusIn(incident.StatusPending),
				r.filterIncidentPredicate(filter),
			),
		).CountX(r.ctx),
		InProgress: r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				incident.StatusIn(incident.StatusInProgress),
				r.filterIncidentPredicate(filter),
			),
		).CountX(r.ctx),
		Resolved: r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				incident.StatusIn(incident.StatusResolved),
				r.filterIncidentPredicate(filter),
			),
		).CountX(r.ctx),
		Accident: r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				incident.TypeIn("accident", "accidents", "Accident", "Accidents"),
				r.filterIncidentPredicate(filter),
			),
		).CountX(r.ctx),
		Mechanical: r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				incident.TypeIn("mechanical", "Mechanical"),
				r.filterIncidentPredicate(filter),
			),
		).CountX(r.ctx),
	}, nil
}

// ReadCompanyMonthRevenue implements gateways.StatisticsRepo.
func (r *repository) ReadCompanyMonthRevenue(companyId int) (*presenters.CompanyMonthRevenue, error) {
	return &presenters.CompanyMonthRevenue{
		CurrentMonth: func() float64 {
			result, err := r.db.Transaction.Query().Where(transaction.And(
				transaction.HasCompanyWith(company.ID(companyId)),
				r.filterTransactionPredicate("this-month"),
			)).Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
			if err != nil {
				return 0
			}
			return result
		}(),
		LastMonth: func() float64 {
			result, err := r.db.Transaction.Query().Where(transaction.And(
				transaction.HasCompanyWith(company.ID(companyId)),
				r.filterTransactionPredicate("last-month"),
			)).Aggregate(ent.Sum(transaction.FieldAmount)).Float64(r.ctx)
			if err != nil {
				return 0
			}
			return result
		}(),
	}, nil
}

// ReadCompanyRevenueOverview implements gateways.StatisticsRepo.
func (r *repository) ReadCompanyRevenueOverview(companyId int, filter string) ([]presenters.CompanyRevenueOverview, error) {
	var trip []companyOverview
	err := r.db.Transaction.Query().Where(
		transaction.And(
			transaction.HasCompanyWith(company.ID(companyId)),
			transaction.ProductIn(transaction.ProductTrip),
			r.filterTransactionPredicate(filter),
		),
	).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As("DATE_FORMAT(created_at, '%w')", "day"),
				sql.As(sql.Sum("amount"), "amount"),
			).
				GroupBy("MONTH(created_at)", "YEAR(created_at)").
				OrderBy(sql.Asc("MONTH(created_at)"))
		}).
		Scan(r.ctx, &trip)
	if err != nil {
		return nil, err
	}

	var delivery []companyOverview
	err = r.db.Transaction.Query().Where(
		transaction.And(
			transaction.HasCompanyWith(company.ID(companyId)),
			transaction.ProductIn(transaction.ProductDelivery),
			r.filterTransactionPredicate(filter),
		),
	).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As("DATE_FORMAT(created_at, '%w')", "day"),
				sql.As(sql.Sum("amount"), "amount"),
			).
				GroupBy("MONTH(created_at)", "YEAR(created_at)").
				OrderBy(sql.Asc("MONTH(created_at)"))
		}).
		Scan(r.ctx, &delivery)
	if err != nil {
		return nil, err
	}
	return r.formatDayOverview(trip, delivery), nil
}

// ReadCompanyTripOverview implements gateways.StatisticsRepo.
func (r *repository) ReadCompanyTripOverview(companyId int, filter string) (*presenters.CompanyTripOverview, error) {
	return &presenters.CompanyTripOverview{
		Trip: r.db.Trip.Query().Where(
			trip.And(
				trip.HasCompanyWith(company.ID(companyId)),
				r.filterTripPredicate(filter),
			),
		).CountX(r.ctx),
		Customer: func() int {
			rc := r.db.Booking.Query().Where(booking.HasCompanyWith(company.ID(companyId))).QueryCustomer().Where(customer.And(r.filterCustomerPredicate(filter))).CountX(r.ctx)
			contacts := r.db.Booking.Query().Where(booking.HasCompanyWith(company.ID(companyId))).QueryContact().Where(customercontact.And(r.filterCustomerContactPredicate(filter))).Modify(func(s *sql.Selector) {
				s.Select(fmt.Sprintf("DISTINCT %s", customercontact.FieldPhone))
			}).CountX(r.ctx)
			return rc + contacts
		}(),
		Package: r.db.Parcel.Query().Where(
			parcel.And(
				parcel.HasCompanyWith(company.ID(companyId)),
				r.filterParcelPredicate(filter),
			),
		).CountX(r.ctx),
		NewCustomer: func() int {
			rc := r.db.Booking.Query().Where(booking.HasCompanyWith(company.ID(companyId))).QueryCustomer().Where(customer.And(r.filterCustomerPredicate("this-week"), r.filterCustomerPredicate(filter))).CountX(r.ctx)
			contacts := r.db.Booking.Query().Where(booking.HasCompanyWith(company.ID(companyId))).QueryContact().Where(customercontact.And(r.filterCustomerContactPredicate("this-week"), r.filterCustomerContactPredicate(filter))).Modify(func(s *sql.Selector) {
				s.Select(fmt.Sprintf("DISTINCT %s", customercontact.FieldPhone))
			}).CountX(r.ctx)
			return rc + contacts
		}(),
		Incident: r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				r.filterIncidentPredicate(filter),
			),
		).CountX(r.ctx),
	}, nil
}

func (r *repository) formatOverview(overview []adminOverview) []presenters.AdminOverview {
	dataset := make([]presenters.AdminOverview, 0, 12)
	if len(overview) == 0 {
		for _, m := range months {
			var o presenters.AdminOverview
			o.Month = m
			o.Amount = 0
			dataset = append(dataset, o)
		}
		return dataset
	}
	for _, m := range months {
		var o presenters.AdminOverview
		for _, mc := range overview {
			if m == mc.Month {
				o.Month = mc.Month
				o.Amount = mc.Amount
			} else {
				o.Month = m
			}
		}
		dataset = append(dataset, o)
	}
	return dataset
}

func (r *repository) formatUserOverview(overview []presenters.AdminUserOverview) []presenters.AdminUserOverview {
	dataset := make([]presenters.AdminUserOverview, 0, 12)
	if len(overview) == 0 {
		for _, m := range months {
			var o presenters.AdminUserOverview
			o.Month = m
			o.User = 0
			dataset = append(dataset, o)
		}
		return dataset
	}
	for _, m := range months {
		var o presenters.AdminUserOverview
		for _, mc := range overview {

			if m == mc.Month {
				o.Month = mc.Month
				o.User = mc.User
			} else {
				o.Month = m
			}
		}
		dataset = append(dataset, o)
	}

	return dataset
}

func (r *repository) formatDayOverview(trip, delivery []companyOverview) []presenters.CompanyRevenueOverview {
	dataset := make([]presenters.CompanyRevenueOverview, 0, 7)
	for i, m := range days {
		var o presenters.CompanyRevenueOverview
		for _, t := range trip {
			o.Day = m
			if (i + 1) == t.Day {
				o.Trip = t.Amount
			}
		}
		for _, p := range delivery {
			o.Day = m
			if (i + 1) == p.Day {
				o.Package = p.Amount
			}
		}
		dataset = append(dataset, o)
	}
	return dataset
}

func (r *repository) filterTransactionPredicate(filter string) predicate.Transaction {
	switch strings.ToLower(filter) {
	case "today":
		return transaction.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return transaction.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return transaction.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return transaction.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return transaction.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return transaction.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return transaction.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterCompanyUserPredicate(filter string) predicate.CompanyUser {
	switch strings.ToLower(filter) {
	case "today":
		return companyuser.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return companyuser.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return companyuser.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return companyuser.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return companyuser.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return companyuser.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return companyuser.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterTerminalPredicate(filter string) predicate.Terminal {
	switch strings.ToLower(filter) {
	case "today":
		return terminal.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return terminal.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return terminal.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return terminal.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return terminal.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return terminal.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return terminal.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterVehiclePredicate(filter string) predicate.Vehicle {
	switch strings.ToLower(filter) {
	case "today":
		return vehicle.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return vehicle.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return vehicle.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return vehicle.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return vehicle.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return vehicle.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return vehicle.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterCustomerPredicate(filter string) predicate.Customer {
	switch strings.ToLower(filter) {
	case "today":
		return customer.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return customer.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return customer.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return customer.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return customer.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return customer.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return customer.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterCustomerContactPredicate(filter string) predicate.CustomerContact {
	switch strings.ToLower(filter) {
	case "today":
		return customercontact.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return customercontact.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return customercontact.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return customercontact.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return customercontact.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return customercontact.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return customercontact.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterIncidentPredicate(filter string) predicate.Incident {
	switch strings.ToLower(filter) {
	case "today":
		return incident.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return incident.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return incident.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return incident.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return incident.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return incident.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return incident.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterTripPredicate(filter string) predicate.Trip {
	switch strings.ToLower(filter) {
	case "today":
		return trip.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return trip.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return trip.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return trip.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return trip.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return trip.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return trip.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterParcelPredicate(filter string) predicate.Parcel {
	switch strings.ToLower(filter) {
	case "today":
		return parcel.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return parcel.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return parcel.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return parcel.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return parcel.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return parcel.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return parcel.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}

func (r *repository) filterUserPredicate(filter string) predicate.User {
	switch strings.ToLower(filter) {
	case "today":
		return user.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY"))
			},
		)
	case "this-week":
		return user.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE())"))
			},
		)
	case "last-week":
		return user.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("WEEKOFYEAR(created_at) = WEEKOFYEAR(CURDATE()) - 1"))
			},
		)
	case "this-month":
		return user.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH)  AND CURDATE()"))
			},
		)
	case "last-month":
		return user.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("MONTH(created_at) = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 MONTH))"))
			},
		)
	case "last-year":
		return user.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 YEAR))"))
			},
		)
	default:
		return user.And(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("YEAR(created_at) = YEAR(CURDATE())"))
			},
		)
	}
}
