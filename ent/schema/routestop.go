package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RouteStop holds the schema definition for the RouteStop entity.
type RouteStop struct {
	ent.Schema
}

// Time mixin.
func (RouteStop) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the RouteStop.
func (RouteStop) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Optional(),
		field.Float("latitude").Optional(),
		field.Float("longitude").Optional(),
	}
}

// Edges of the RouteStop.
func (RouteStop) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("stops").
			Unique(),
		edge.From("trip", Trip.Type).
			Ref("stops"),
	}
}
