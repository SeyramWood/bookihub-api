// Code generated by ent, DO NOT EDIT.

package customerluggage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldUpdatedAt, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldQuantity, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldAmount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLTE(FieldUpdatedAt, v))
}

// BaggageEQ applies the EQ predicate on the "baggage" field.
func BaggageEQ(v Baggage) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldBaggage, v))
}

// BaggageNEQ applies the NEQ predicate on the "baggage" field.
func BaggageNEQ(v Baggage) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNEQ(FieldBaggage, v))
}

// BaggageIn applies the In predicate on the "baggage" field.
func BaggageIn(vs ...Baggage) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldIn(FieldBaggage, vs...))
}

// BaggageNotIn applies the NotIn predicate on the "baggage" field.
func BaggageNotIn(vs ...Baggage) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNotIn(FieldBaggage, vs...))
}

// BaggageIsNil applies the IsNil predicate on the "baggage" field.
func BaggageIsNil() predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldIsNull(FieldBaggage))
}

// BaggageNotNil applies the NotNil predicate on the "baggage" field.
func BaggageNotNil() predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNotNull(FieldBaggage))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLTE(FieldQuantity, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.FieldLTE(FieldAmount, v))
}

// HasBooking applies the HasEdge predicate on the "booking" edge.
func HasBooking() predicate.CustomerLuggage {
	return predicate.CustomerLuggage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BookingTable, BookingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingWith applies the HasEdge predicate on the "booking" edge with a given conditions (other predicates).
func HasBookingWith(preds ...predicate.Booking) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(func(s *sql.Selector) {
		step := newBookingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CustomerLuggage) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CustomerLuggage) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CustomerLuggage) predicate.CustomerLuggage {
	return predicate.CustomerLuggage(sql.NotPredicates(p))
}
