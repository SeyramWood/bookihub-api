package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Trip holds the schema definition for the Trip entity.
type Trip struct {
	ent.Schema
}

// Time mixin.
func (Trip) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Trip.
func (Trip) Fields() []ent.Field {
	return []ent.Field{
		field.Time("departure_date").Optional(),
		field.Time("arrival_date").Optional(),
		field.Time("return_date").Optional(),
		field.Enum("type").Values("oneway", "round").Default("oneway"),
		field.Bool("exterior_inspected").Default(false),
		field.Bool("interior_inspected").Default(false),
		field.Bool("engine_compartment_inspected").Default(false),
		field.Bool("brake_and_steering_inspected").Default(false),
		field.Bool("emergency_equipment_inspected").Default(false),
		field.Bool("fuel_and_fluids_inspected").Default(false),
		field.Bool("scheduled").Default(false),
		field.Int("seat_left").Default(0),
		field.Float("rate").Default(0.00),
		field.Float32("discount").Default(0.00),
		field.Enum("status").Values("started", "ended").Optional(),
	}
}

// Edges of the Trip.
func (Trip) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("trips").
			Unique(),
		edge.From("driver", CompanyUser.Type).
			Ref("trips").
			Unique(),
		edge.From("from_terminal", Terminal.Type).
			Ref("from").
			Unique(),
		edge.From("to_terminal", Terminal.Type).
			Ref("to").
			Unique(),
		edge.From("vehicle", Vehicle.Type).
			Ref("trips").
			Unique(),
		edge.From("route", Route.Type).
			Ref("trips").
			Unique(),
		edge.To("bookings", Booking.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("incidents", Incident.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("parcels", Parcel.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
