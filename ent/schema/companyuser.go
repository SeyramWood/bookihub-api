package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyUser holds the schema definition for the CompanyUser entity.
type CompanyUser struct {
	ent.Schema
}

// Time mixin.
func (CompanyUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the CompanyUser.
func (CompanyUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").Optional(),
		field.String("other_name").Optional(),
		field.String("phone").Optional(),
		field.String("other_phone").Optional(),
		field.Enum("role").Values("admin", "manager", "teller", "driver").Default("admin"),
	}
}

// Edges of the CompanyUser.
func (CompanyUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", User.Type).Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("trips", Trip.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("incidents", Incident.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("parcels", Parcel.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("notifications", Notification.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("company", Company.Type).
			Ref("profile").
			Unique(),
	}
}
