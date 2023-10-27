package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VehicleImage holds the schema definition for the VehicleImage entity.
type VehicleImage struct {
	ent.Schema
}

// Time mixin.
func (VehicleImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the VehicleImage.
func (VehicleImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("image").NotEmpty(),
	}
}

// Edges of the VehicleImage.
func (VehicleImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("vehicle", Vehicle.Type).
			Ref("images").
			Unique(),
	}
}
