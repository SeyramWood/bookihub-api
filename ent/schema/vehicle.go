package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Time mixin.
func (Vehicle) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.String("registration_number").NotEmpty(),
		field.String("model").NotEmpty(),
		field.Int("seat").Default(0),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("vehicles").
			Unique(),
		edge.To("images", VehicleImage.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("trips", Trip.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
