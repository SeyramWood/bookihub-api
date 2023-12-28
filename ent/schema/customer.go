package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Time mixin.
func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", User.Type).Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("bookings", Booking.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("notifications", Notification.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
