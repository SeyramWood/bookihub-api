package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Time mixin.
func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("reference").NotEmpty(),
		field.Float("amount").Default(0.00),
		field.Float("vat").Default(0.00),
		field.Float("transaction_fee").Default(0.00),
		field.Float("cancellation_fee").Default(0.00),
		field.Time("paid_at").Optional(),
		field.Time("canceled_at").Optional(),
		field.Enum("channel").Values("momo", "card", "bank", "cash").Default("cash"),
		field.Enum("tans_kind").Values("payment", "payout").Default("payment"),
		field.Enum("product").Values("trip", "delivery").Default("trip"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking", Booking.Type).
			Ref("transaction").
			Unique(),
		edge.From("parcel", Parcel.Type).
			Ref("transaction").
			Unique(),
		edge.From("company", Company.Type).
			Ref("transactions").
			Required().
			Unique(),
	}
}
