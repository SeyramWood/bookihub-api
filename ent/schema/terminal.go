package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Terminal holds the schema definition for the Terminal entity.
type Terminal struct {
	ent.Schema
}

// Time mixin.
func (Terminal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Terminal.
func (Terminal) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Optional(),
		field.Float("latitude").Optional(),
		field.Float("longitude").Optional(),
	}
}

// Edges of the Terminal.
func (Terminal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("terminals").
			Unique(),
		edge.To("from", Trip.Type),
		edge.To("to", Trip.Type),
	}
}
