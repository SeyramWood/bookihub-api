package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IncidentImage holds the schema definition for the IncidentImage entity.
type IncidentImage struct {
	ent.Schema
}

// Time mixin.
func (IncidentImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the IncidentImage.
func (IncidentImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("image").Optional(),
	}
}

// Edges of the IncidentImage.
func (IncidentImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("incident", Incident.Type).
			Ref("images").
			Unique(),
	}
}
