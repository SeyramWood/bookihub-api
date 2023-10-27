package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Passenger holds the schema definition for the Passenger entity.
type Passenger struct {
	ent.Schema
}

// Time mixin.
func (Passenger) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Passenger.
func (Passenger) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").NotEmpty(),
		field.Float("amount").Default(0.00),
		field.Enum("maturity").Values("adult", "child").Default("adult"),
		field.Enum("gender").Values("male", "female", "other").Optional(),
	}
}

// Edges of the Passenger.
func (Passenger) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking", Booking.Type).
			Ref("passengers").
			Unique().
			Required(),
	}
}
