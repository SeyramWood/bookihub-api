package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type (
	Charge struct {
		PaymentGatewayServiceFee float32 `json:"paymentGatewayServiceFee"`
		TripServiceFee           float32 `json:"tripServiceFee"`
		ParcelServiceFee         float32 `json:"deliveryServiceFee"`
		TripCancellationFee      float32 `json:"tripCancellationFee"`
	}
)

// Configuration holds the schema definition for the Configuration entity.
type Configuration struct {
	ent.Schema
}

func (Configuration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Configuration.
func (Configuration) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("charge", &Charge{}).Default(&Charge{
			PaymentGatewayServiceFee: 1.95,
			TripServiceFee:           2.50,
			ParcelServiceFee:         2.00,
			TripCancellationFee:      4.45,
		}),
	}
}

// Edges of the Configuration.
func (Configuration) Edges() []ent.Edge {
	return nil
}
