package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomerContact holds the schema definition for the CustomerContact entity.
type CustomerContact struct {
	ent.Schema
}

// Time mixin.
func (CustomerContact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the CustomerContact.
func (CustomerContact) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").NotEmpty(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
	}

}

// Edges of the CustomerContact.
func (CustomerContact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking", Booking.Type).
			Ref("contact").
			Unique(),
	}
}
