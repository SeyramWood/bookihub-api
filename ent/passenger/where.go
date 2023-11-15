// Code generated by ent, DO NOT EDIT.

package passenger

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Passenger {
	return predicate.Passenger(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Passenger {
	return predicate.Passenger(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Passenger {
	return predicate.Passenger(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Passenger {
	return predicate.Passenger(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Passenger {
	return predicate.Passenger(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Passenger {
	return predicate.Passenger(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Passenger {
	return predicate.Passenger(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldUpdatedAt, v))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldFullName, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldAmount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Passenger {
	return predicate.Passenger(sql.FieldLTE(FieldUpdatedAt, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.Passenger {
	return predicate.Passenger(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.Passenger {
	return predicate.Passenger(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.Passenger {
	return predicate.Passenger(sql.FieldContainsFold(FieldFullName, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Passenger {
	return predicate.Passenger(sql.FieldLTE(FieldAmount, v))
}

// MaturityEQ applies the EQ predicate on the "maturity" field.
func MaturityEQ(v Maturity) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldMaturity, v))
}

// MaturityNEQ applies the NEQ predicate on the "maturity" field.
func MaturityNEQ(v Maturity) predicate.Passenger {
	return predicate.Passenger(sql.FieldNEQ(FieldMaturity, v))
}

// MaturityIn applies the In predicate on the "maturity" field.
func MaturityIn(vs ...Maturity) predicate.Passenger {
	return predicate.Passenger(sql.FieldIn(FieldMaturity, vs...))
}

// MaturityNotIn applies the NotIn predicate on the "maturity" field.
func MaturityNotIn(vs ...Maturity) predicate.Passenger {
	return predicate.Passenger(sql.FieldNotIn(FieldMaturity, vs...))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v Gender) predicate.Passenger {
	return predicate.Passenger(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v Gender) predicate.Passenger {
	return predicate.Passenger(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...Gender) predicate.Passenger {
	return predicate.Passenger(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...Gender) predicate.Passenger {
	return predicate.Passenger(sql.FieldNotIn(FieldGender, vs...))
}

// GenderIsNil applies the IsNil predicate on the "gender" field.
func GenderIsNil() predicate.Passenger {
	return predicate.Passenger(sql.FieldIsNull(FieldGender))
}

// GenderNotNil applies the NotNil predicate on the "gender" field.
func GenderNotNil() predicate.Passenger {
	return predicate.Passenger(sql.FieldNotNull(FieldGender))
}

// HasBooking applies the HasEdge predicate on the "booking" edge.
func HasBooking() predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BookingTable, BookingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingWith applies the HasEdge predicate on the "booking" edge with a given conditions (other predicates).
func HasBookingWith(preds ...predicate.Booking) predicate.Passenger {
	return predicate.Passenger(func(s *sql.Selector) {
		step := newBookingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Passenger) predicate.Passenger {
	return predicate.Passenger(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Passenger) predicate.Passenger {
	return predicate.Passenger(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Passenger) predicate.Passenger {
	return predicate.Passenger(sql.NotPredicates(p))
}
