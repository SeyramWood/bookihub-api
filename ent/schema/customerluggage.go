package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomerLuggage holds the schema definition for the CustomerLuggage entity.
type CustomerLuggage struct {
	ent.Schema
}

// Time mixin.
func (CustomerLuggage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the CustomerLuggage.
func (CustomerLuggage) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("baggage").Values("excess", "bulky").Optional(),
		field.Int("quantity").Default(0),
		field.Float("amount").Default(0.00),
	}

}

// Edges of the CustomerLuggage.
func (CustomerLuggage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking", Booking.Type).
			Ref("luggages").
			Unique(),
	}
}
