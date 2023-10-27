package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Time mixin.
func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", CompanyUser.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("vehicles", Vehicle.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("routes", Route.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("trips", Trip.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("bookings", Booking.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("notifications", Notification.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
