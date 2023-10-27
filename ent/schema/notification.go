package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type NotificationRead struct {
	ID     int    `json:"id"`
	ReadAt string `json:"readAt"`
}

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("event").NotEmpty(),
		field.String("activity").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("subject_type").NotEmpty(),
		field.Int("subject_id").Optional(),
		field.String("creator_type").NotEmpty(),
		field.Int("creator_id").Optional(),
		field.String("customer_read_at").Optional(),
		field.JSON(
			"bookibus_read_at", []*NotificationRead{},
		).Optional(),
		field.JSON(
			"company_read_at", []*NotificationRead{},
		).Optional(),
		field.JSON(
			"data", &struct {
				Data any `json:"data"`
			}{},
		).Optional(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bookibus_user", BookibusUser.Type).
			Ref("notifications"),
		edge.From("company_user", CompanyUser.Type).
			Ref("notifications"),
		edge.From("customer", Customer.Type).
			Ref("notifications"),
		edge.From("company", Company.Type).
			Ref("notifications").
			Unique(),
	}
}
