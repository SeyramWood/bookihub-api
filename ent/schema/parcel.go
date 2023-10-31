package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Parcel holds the schema definition for the Parcel entity.
type Parcel struct {
	ent.Schema
}

// Time mixin.
func (Parcel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Parcel.
func (Parcel) Fields() []ent.Field {
	return []ent.Field{
		field.String("parcel_code").NotEmpty(),
		field.String("sender_name").NotEmpty(),
		field.String("sender_phone").NotEmpty(),
		field.String("recipient_name").NotEmpty(),
		field.String("recipient_phone").NotEmpty(),
		field.String("recipient_location").NotEmpty(),
		field.Float("amount").Default(0.00),
		field.Time("paid_at").Optional(),
		field.Enum("tans_type").Values("momo", "card", "cash").Default("cash"),
		field.Enum("status").Values("outgoing", "delivered").Default("outgoing"),
	}
}

// Edges of the Parcel.
func (Parcel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", ParcelImage.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("trip", Trip.Type).
			Ref("parcels").
			Unique(),
		edge.From("company", Company.Type).
			Ref("parcels").
			Unique(),
		edge.From("driver", CompanyUser.Type).
			Ref("parcels").
			Unique(),
	}
}
