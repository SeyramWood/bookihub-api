package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Time mixin.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.Enum("type").Values("bookibus", "company", "customer"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bookibus_user", BookibusUser.Type).
			Ref("profile").
			Unique(),
		edge.From("company_user", CompanyUser.Type).
			Ref("profile").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("profile").
			Unique(),
	}
}
