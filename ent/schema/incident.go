package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Incident holds the schema definition for the Incident entity.
type Incident struct {
	ent.Schema
}

// Time mixin.
func (Incident) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Incident.
func (Incident) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time").Optional(),
		field.String("location").NotEmpty(),
		field.Text("description").NotEmpty(),
		field.String("audio").Optional(),
	}
}

// Edges of the Incident.
func (Incident) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", IncidentImage.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("trip", Trip.Type).
			Ref("incidents").
			Unique(),
		edge.From("company", Company.Type).
			Ref("incidents").
			Unique(),
		edge.From("driver", CompanyUser.Type).
			Ref("incidents").
			Unique(),
	}
}
