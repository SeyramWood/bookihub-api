// Code generated by ent, DO NOT EDIT.

package companyuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldLastName, v))
}

// OtherName applies equality check predicate on the "other_name" field. It's identical to OtherNameEQ.
func OtherName(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldOtherName, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldPhone, v))
}

// OtherPhone applies equality check predicate on the "other_phone" field. It's identical to OtherPhoneEQ.
func OtherPhone(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldOtherPhone, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContainsFold(FieldLastName, v))
}

// OtherNameEQ applies the EQ predicate on the "other_name" field.
func OtherNameEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldOtherName, v))
}

// OtherNameNEQ applies the NEQ predicate on the "other_name" field.
func OtherNameNEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldOtherName, v))
}

// OtherNameIn applies the In predicate on the "other_name" field.
func OtherNameIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldOtherName, vs...))
}

// OtherNameNotIn applies the NotIn predicate on the "other_name" field.
func OtherNameNotIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldOtherName, vs...))
}

// OtherNameGT applies the GT predicate on the "other_name" field.
func OtherNameGT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGT(FieldOtherName, v))
}

// OtherNameGTE applies the GTE predicate on the "other_name" field.
func OtherNameGTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGTE(FieldOtherName, v))
}

// OtherNameLT applies the LT predicate on the "other_name" field.
func OtherNameLT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLT(FieldOtherName, v))
}

// OtherNameLTE applies the LTE predicate on the "other_name" field.
func OtherNameLTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLTE(FieldOtherName, v))
}

// OtherNameContains applies the Contains predicate on the "other_name" field.
func OtherNameContains(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContains(FieldOtherName, v))
}

// OtherNameHasPrefix applies the HasPrefix predicate on the "other_name" field.
func OtherNameHasPrefix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasPrefix(FieldOtherName, v))
}

// OtherNameHasSuffix applies the HasSuffix predicate on the "other_name" field.
func OtherNameHasSuffix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasSuffix(FieldOtherName, v))
}

// OtherNameIsNil applies the IsNil predicate on the "other_name" field.
func OtherNameIsNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIsNull(FieldOtherName))
}

// OtherNameNotNil applies the NotNil predicate on the "other_name" field.
func OtherNameNotNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotNull(FieldOtherName))
}

// OtherNameEqualFold applies the EqualFold predicate on the "other_name" field.
func OtherNameEqualFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEqualFold(FieldOtherName, v))
}

// OtherNameContainsFold applies the ContainsFold predicate on the "other_name" field.
func OtherNameContainsFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContainsFold(FieldOtherName, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContainsFold(FieldPhone, v))
}

// OtherPhoneEQ applies the EQ predicate on the "other_phone" field.
func OtherPhoneEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldOtherPhone, v))
}

// OtherPhoneNEQ applies the NEQ predicate on the "other_phone" field.
func OtherPhoneNEQ(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldOtherPhone, v))
}

// OtherPhoneIn applies the In predicate on the "other_phone" field.
func OtherPhoneIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldOtherPhone, vs...))
}

// OtherPhoneNotIn applies the NotIn predicate on the "other_phone" field.
func OtherPhoneNotIn(vs ...string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldOtherPhone, vs...))
}

// OtherPhoneGT applies the GT predicate on the "other_phone" field.
func OtherPhoneGT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGT(FieldOtherPhone, v))
}

// OtherPhoneGTE applies the GTE predicate on the "other_phone" field.
func OtherPhoneGTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldGTE(FieldOtherPhone, v))
}

// OtherPhoneLT applies the LT predicate on the "other_phone" field.
func OtherPhoneLT(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLT(FieldOtherPhone, v))
}

// OtherPhoneLTE applies the LTE predicate on the "other_phone" field.
func OtherPhoneLTE(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldLTE(FieldOtherPhone, v))
}

// OtherPhoneContains applies the Contains predicate on the "other_phone" field.
func OtherPhoneContains(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContains(FieldOtherPhone, v))
}

// OtherPhoneHasPrefix applies the HasPrefix predicate on the "other_phone" field.
func OtherPhoneHasPrefix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasPrefix(FieldOtherPhone, v))
}

// OtherPhoneHasSuffix applies the HasSuffix predicate on the "other_phone" field.
func OtherPhoneHasSuffix(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldHasSuffix(FieldOtherPhone, v))
}

// OtherPhoneIsNil applies the IsNil predicate on the "other_phone" field.
func OtherPhoneIsNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIsNull(FieldOtherPhone))
}

// OtherPhoneNotNil applies the NotNil predicate on the "other_phone" field.
func OtherPhoneNotNil() predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotNull(FieldOtherPhone))
}

// OtherPhoneEqualFold applies the EqualFold predicate on the "other_phone" field.
func OtherPhoneEqualFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEqualFold(FieldOtherPhone, v))
}

// OtherPhoneContainsFold applies the ContainsFold predicate on the "other_phone" field.
func OtherPhoneContainsFold(v string) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldContainsFold(FieldOtherPhone, v))
}

// UserRoleEQ applies the EQ predicate on the "user_role" field.
func UserRoleEQ(v UserRole) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldEQ(FieldUserRole, v))
}

// UserRoleNEQ applies the NEQ predicate on the "user_role" field.
func UserRoleNEQ(v UserRole) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNEQ(FieldUserRole, v))
}

// UserRoleIn applies the In predicate on the "user_role" field.
func UserRoleIn(vs ...UserRole) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldIn(FieldUserRole, vs...))
}

// UserRoleNotIn applies the NotIn predicate on the "user_role" field.
func UserRoleNotIn(vs ...UserRole) predicate.CompanyUser {
	return predicate.CompanyUser(sql.FieldNotIn(FieldUserRole, vs...))
}

// HasProfile applies the HasEdge predicate on the "profile" edge.
func HasProfile() predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfileWith applies the HasEdge predicate on the "profile" edge with a given conditions (other predicates).
func HasProfileWith(preds ...predicate.User) predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := newProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTrips applies the HasEdge predicate on the "trips" edge.
func HasTrips() predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TripsTable, TripsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTripsWith applies the HasEdge predicate on the "trips" edge with a given conditions (other predicates).
func HasTripsWith(preds ...predicate.Trip) predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := newTripsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIncidents applies the HasEdge predicate on the "incidents" edge.
func HasIncidents() predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IncidentsTable, IncidentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIncidentsWith applies the HasEdge predicate on the "incidents" edge with a given conditions (other predicates).
func HasIncidentsWith(preds ...predicate.Incident) predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := newIncidentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParcels applies the HasEdge predicate on the "parcels" edge.
func HasParcels() predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParcelsTable, ParcelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParcelsWith applies the HasEdge predicate on the "parcels" edge with a given conditions (other predicates).
func HasParcelsWith(preds ...predicate.Parcel) predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := newParcelsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifications applies the HasEdge predicate on the "notifications" edge.
func HasNotifications() predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NotificationsTable, NotificationsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationsWith applies the HasEdge predicate on the "notifications" edge with a given conditions (other predicates).
func HasNotificationsWith(preds ...predicate.Notification) predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := newNotificationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.CompanyUser {
	return predicate.CompanyUser(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CompanyUser) predicate.CompanyUser {
	return predicate.CompanyUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CompanyUser) predicate.CompanyUser {
	return predicate.CompanyUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CompanyUser) predicate.CompanyUser {
	return predicate.CompanyUser(sql.NotPredicates(p))
}
