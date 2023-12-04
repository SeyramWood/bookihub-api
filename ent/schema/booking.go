package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Time mixin.
func (Booking) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.String("booking_number").NotEmpty(),
		field.Bool("sms_notification").Default(false),
		field.Enum("status").Values("successful", "canceled").Default("successful"),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("passengers", Passenger.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("luggages", CustomerLuggage.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("contact", CustomerContact.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).Unique(),
		edge.To("transaction", Transaction.Type).Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("trip", Trip.Type).
			Ref("bookings").
			Unique(),
		edge.From("company", Company.Type).
			Ref("bookings").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("bookings").
			Unique(),
	}
}
