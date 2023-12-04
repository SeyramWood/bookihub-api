package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Route holds the schema definition for the Route entity.
type Route struct {
	ent.Schema
}

// Time mixin.
func (Route) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Route.
func (Route) Fields() []ent.Field {
	return []ent.Field{
		field.String("from_location").NotEmpty(),
		field.String("to_location").NotEmpty(),
		field.Float("from_latitude").Optional(),
		field.Float("from_longitude").Optional(),
		field.Float("to_latitude").Optional(),
		field.Float("to_longitude").Optional(),
		field.Int("popularity").Default(0),
	}
}

// Edges of the Route.
func (Route) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("routes").
			Unique(),
		edge.To("stops", RouteStop.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("trips", Trip.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
