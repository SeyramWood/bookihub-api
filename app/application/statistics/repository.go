package statistics

import (
	"context"
	"log"

	"entgo.io/ent/dialect/sql"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/transaction"
)

var months = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

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

	var users []struct {
		ID      int
		Name    string
		Product string
		Amount  float64
	}
	err := r.db.Company.Query().WithTransactions().
		GroupBy(company.FieldID, company.FieldName).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(transaction.Table)
			s.Join(t).On(s.C(company.FieldID), t.C(company.TransactionsColumn))
			s.GroupBy(transaction.FieldProduct)
			sql.As(t.C(transaction.FieldProduct), "product")
			return sql.As(sql.Sum(t.C(transaction.FieldAmount)), "amount")
		}).
		Scan(r.ctx, &users)

	// err := r.db.Transaction.Query().
	// 	GroupBy(transaction.FieldProduct).
	// 	Aggregate(func(s *sql.Selector) string {
	// 		t := sql.Table(company.Table)
	// 		s.Join(t).On(s.C(transaction.FieldID), t.C(transaction.CompanyColumn))
	// 		return ""
	// 		// return sql.As(sql.Sum(t.C(transaction.FieldAmount)), "amount")
	// 	}).
	// 	Scan(r.ctx, &users)

	if err != nil {
		return nil, err
	}
	log.Println("users: ", users)
	panic("unimplemented")
	// count := query.CountX(r.ctx)
	// results, err := query.
	// 	Limit(limit).
	// 	Offset(offset).
	// 	Order(
	// 		trip.ByRouteField(
	// 			route.FieldPopularity,
	// 			sql.OrderDesc(),
	// 		),
	// 	).
	// 	WithFromTerminal().
	// 	WithToTerminal().
	// 	WithVehicle(func(vq *ent.VehicleQuery) {
	// 		vq.WithImages()
	// 	}).
	// 	WithRoute().
	// 	WithStops().
	// 	WithDriver().
	// 	WithCompany().
	// 	WithBookings(func(bq *ent.BookingQuery) {
	// 		bq.WithPassengers()
	// 		bq.WithLuggages()
	// 		bq.WithContact()
	// 		bq.WithCustomer(func(cq *ent.CustomerQuery) {
	// 			cq.WithProfile()
	// 		})
	// 	}).
	// 	WithParcels(func(pq *ent.ParcelQuery) {
	// 		pq.WithImages()
	// 	}).
	// 	WithIncidents(func(iq *ent.IncidentQuery) {
	// 		iq.WithImages()
	// 	}).
	// 	All(r.ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// return application.Paginate(count, results)

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
func (r *repository) ReadAdminRevenueOverview(filter string) ([]*presenters.AdminOverview, error) {
	panic("unimplemented")
}
