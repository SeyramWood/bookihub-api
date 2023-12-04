package configuration

import (
	"context"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/configuration"
	"github.com/SeyramWood/bookibus/ent/schema"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.ConfigurationRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// InsertCharge implements gateways.ConfigurationRepo.
func (r *repository) InsertCharge(request *requeststructs.TransactionChargeRequest) (*ent.Configuration, error) {
	if results := r.db.Configuration.Query().AllX(r.ctx); len(results) > 0 {
		result, err := r.db.Configuration.UpdateOne(results[0]).
			SetCharge(&schema.Charge{
				PaymentGatewayServiceFee: request.PaymentGatewayServiceFee,
				TripServiceFee:           request.TripServiceFee,
				ParcelServiceFee:         request.ParcelServiceFee,
				TripCancellationFee:      request.TripCancellationFee,
			}).
			Save(r.ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	result, err := r.db.Configuration.Create().
		SetCharge(&schema.Charge{
			PaymentGatewayServiceFee: request.PaymentGatewayServiceFee,
			TripServiceFee:           request.TripServiceFee,
			ParcelServiceFee:         request.ParcelServiceFee,
			TripCancellationFee:      request.TripCancellationFee,
		}).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Read implements gateways.ConfigurationRepo.
func (r *repository) Read() (*ent.Configuration, error) {
	results, err := r.db.Configuration.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results[0], nil
}

// ReadCharge implements gateways.ConfigurationRepo.
func (r *repository) ReadCharge() (*ent.Configuration, error) {
	results, err := r.db.Configuration.Query().Select(configuration.FieldID, configuration.FieldCharge).All(r.ctx)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results[0], nil
}

// UpdateCharge implements gateways.ConfigurationRepo.
func (r *repository) UpdateCharge(id int, request *requeststructs.TransactionChargeRequest) (*ent.Configuration, error) {
	_, err := r.db.Configuration.UpdateOneID(id).SetCharge(&schema.Charge{
		PaymentGatewayServiceFee: request.PaymentGatewayServiceFee,
		TripServiceFee:           request.TripServiceFee,
		ParcelServiceFee:         request.ParcelServiceFee,
		TripCancellationFee:      request.TripCancellationFee,
	}).Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.ReadCharge()
}
