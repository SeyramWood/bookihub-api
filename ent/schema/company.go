package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type (
	BankAccount struct {
		AccountName   string `json:"accountName"`
		AccountNumber string `json:"accountNumber"`
		Bank          string `json:"bank"`
		Branch        string `json:"branch"`
	}
	ContactPerson struct {
		Name     string `json:"name"`
		Position string `json:"position"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
	}
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Time mixin.
func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("certificate").Optional(),
		field.JSON("bank_account", &BankAccount{}).Optional(),
		field.JSON("contact_person", &ContactPerson{}).Optional(),
		field.String("logo").Optional(),
		field.Enum("onboarding_status").Values("pending", "approved", "rejected").Default("pending"),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", CompanyUser.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("terminals", Terminal.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("vehicles", Vehicle.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("routes", Route.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("trips", Trip.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("bookings", Booking.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("incidents", Incident.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("parcels", Parcel.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("transactions", Transaction.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("notifications", Notification.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
