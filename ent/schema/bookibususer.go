package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BookibusUser holds the schema definition for the BookibusUser entity.
type BookibusUser struct {
	ent.Schema
}

// Time mixin.
func (BookibusUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the BookibusUser.
func (BookibusUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").NotEmpty(),
		field.String("other_name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("other_phone").Optional(),
		field.Enum("role").Values("super_admin", "admin").Default("super_admin"),
	}
}

// Edges of the BookibusUser.
func (BookibusUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", User.Type).Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("notifications", Notification.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
