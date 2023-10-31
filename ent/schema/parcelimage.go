package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParcelImage holds the schema definition for the ParcelImage entity.
type ParcelImage struct {
	ent.Schema
}

// Time mixin.
func (ParcelImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the ParcelImage.
func (ParcelImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("image").Optional(),
		field.Enum("kind").Values("parcel", "recipient").Default("parcel"),
	}
}

// Edges of the ParcelImage.
func (ParcelImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parcel", Parcel.Type).
			Ref("images").
			Unique(),
	}
}
