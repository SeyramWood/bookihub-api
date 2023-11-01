// Code generated by ent, DO NOT EDIT.

package trip

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldUpdatedAt, v))
}

// DepartureDate applies equality check predicate on the "departure_date" field. It's identical to DepartureDateEQ.
func DepartureDate(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldDepartureDate, v))
}

// ArrivalDate applies equality check predicate on the "arrival_date" field. It's identical to ArrivalDateEQ.
func ArrivalDate(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldArrivalDate, v))
}

// ReturnDate applies equality check predicate on the "return_date" field. It's identical to ReturnDateEQ.
func ReturnDate(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldReturnDate, v))
}

// ExteriorInspected applies equality check predicate on the "exterior_inspected" field. It's identical to ExteriorInspectedEQ.
func ExteriorInspected(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldExteriorInspected, v))
}

// InteriorInspected applies equality check predicate on the "interior_inspected" field. It's identical to InteriorInspectedEQ.
func InteriorInspected(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldInteriorInspected, v))
}

// EngineCompartmentInspected applies equality check predicate on the "engine_compartment_inspected" field. It's identical to EngineCompartmentInspectedEQ.
func EngineCompartmentInspected(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEngineCompartmentInspected, v))
}

// BrakeAndSteeringInspected applies equality check predicate on the "brake_and_steering_inspected" field. It's identical to BrakeAndSteeringInspectedEQ.
func BrakeAndSteeringInspected(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldBrakeAndSteeringInspected, v))
}

// EmergencyEquipmentInspected applies equality check predicate on the "emergency_equipment_inspected" field. It's identical to EmergencyEquipmentInspectedEQ.
func EmergencyEquipmentInspected(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEmergencyEquipmentInspected, v))
}

// FuelAndFluidsInspected applies equality check predicate on the "fuel_and_fluids_inspected" field. It's identical to FuelAndFluidsInspectedEQ.
func FuelAndFluidsInspected(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldFuelAndFluidsInspected, v))
}

// Scheduled applies equality check predicate on the "scheduled" field. It's identical to ScheduledEQ.
func Scheduled(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldScheduled, v))
}

// SeatLeft applies equality check predicate on the "seat_left" field. It's identical to SeatLeftEQ.
func SeatLeft(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldSeatLeft, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldUpdatedAt, v))
}

// DepartureDateEQ applies the EQ predicate on the "departure_date" field.
func DepartureDateEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldDepartureDate, v))
}

// DepartureDateNEQ applies the NEQ predicate on the "departure_date" field.
func DepartureDateNEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldDepartureDate, v))
}

// DepartureDateIn applies the In predicate on the "departure_date" field.
func DepartureDateIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldDepartureDate, vs...))
}

// DepartureDateNotIn applies the NotIn predicate on the "departure_date" field.
func DepartureDateNotIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldDepartureDate, vs...))
}

// DepartureDateGT applies the GT predicate on the "departure_date" field.
func DepartureDateGT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldDepartureDate, v))
}

// DepartureDateGTE applies the GTE predicate on the "departure_date" field.
func DepartureDateGTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldDepartureDate, v))
}

// DepartureDateLT applies the LT predicate on the "departure_date" field.
func DepartureDateLT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldDepartureDate, v))
}

// DepartureDateLTE applies the LTE predicate on the "departure_date" field.
func DepartureDateLTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldDepartureDate, v))
}

// DepartureDateIsNil applies the IsNil predicate on the "departure_date" field.
func DepartureDateIsNil() predicate.Trip {
	return predicate.Trip(sql.FieldIsNull(FieldDepartureDate))
}

// DepartureDateNotNil applies the NotNil predicate on the "departure_date" field.
func DepartureDateNotNil() predicate.Trip {
	return predicate.Trip(sql.FieldNotNull(FieldDepartureDate))
}

// ArrivalDateEQ applies the EQ predicate on the "arrival_date" field.
func ArrivalDateEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldArrivalDate, v))
}

// ArrivalDateNEQ applies the NEQ predicate on the "arrival_date" field.
func ArrivalDateNEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldArrivalDate, v))
}

// ArrivalDateIn applies the In predicate on the "arrival_date" field.
func ArrivalDateIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldArrivalDate, vs...))
}

// ArrivalDateNotIn applies the NotIn predicate on the "arrival_date" field.
func ArrivalDateNotIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldArrivalDate, vs...))
}

// ArrivalDateGT applies the GT predicate on the "arrival_date" field.
func ArrivalDateGT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldArrivalDate, v))
}

// ArrivalDateGTE applies the GTE predicate on the "arrival_date" field.
func ArrivalDateGTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldArrivalDate, v))
}

// ArrivalDateLT applies the LT predicate on the "arrival_date" field.
func ArrivalDateLT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldArrivalDate, v))
}

// ArrivalDateLTE applies the LTE predicate on the "arrival_date" field.
func ArrivalDateLTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldArrivalDate, v))
}

// ArrivalDateIsNil applies the IsNil predicate on the "arrival_date" field.
func ArrivalDateIsNil() predicate.Trip {
	return predicate.Trip(sql.FieldIsNull(FieldArrivalDate))
}

// ArrivalDateNotNil applies the NotNil predicate on the "arrival_date" field.
func ArrivalDateNotNil() predicate.Trip {
	return predicate.Trip(sql.FieldNotNull(FieldArrivalDate))
}

// ReturnDateEQ applies the EQ predicate on the "return_date" field.
func ReturnDateEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldReturnDate, v))
}

// ReturnDateNEQ applies the NEQ predicate on the "return_date" field.
func ReturnDateNEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldReturnDate, v))
}

// ReturnDateIn applies the In predicate on the "return_date" field.
func ReturnDateIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldReturnDate, vs...))
}

// ReturnDateNotIn applies the NotIn predicate on the "return_date" field.
func ReturnDateNotIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldReturnDate, vs...))
}

// ReturnDateGT applies the GT predicate on the "return_date" field.
func ReturnDateGT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldReturnDate, v))
}

// ReturnDateGTE applies the GTE predicate on the "return_date" field.
func ReturnDateGTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldReturnDate, v))
}

// ReturnDateLT applies the LT predicate on the "return_date" field.
func ReturnDateLT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldReturnDate, v))
}

// ReturnDateLTE applies the LTE predicate on the "return_date" field.
func ReturnDateLTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldReturnDate, v))
}

// ReturnDateIsNil applies the IsNil predicate on the "return_date" field.
func ReturnDateIsNil() predicate.Trip {
	return predicate.Trip(sql.FieldIsNull(FieldReturnDate))
}

// ReturnDateNotNil applies the NotNil predicate on the "return_date" field.
func ReturnDateNotNil() predicate.Trip {
	return predicate.Trip(sql.FieldNotNull(FieldReturnDate))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldType, vs...))
}

// ExteriorInspectedEQ applies the EQ predicate on the "exterior_inspected" field.
func ExteriorInspectedEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldExteriorInspected, v))
}

// ExteriorInspectedNEQ applies the NEQ predicate on the "exterior_inspected" field.
func ExteriorInspectedNEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldExteriorInspected, v))
}

// InteriorInspectedEQ applies the EQ predicate on the "interior_inspected" field.
func InteriorInspectedEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldInteriorInspected, v))
}

// InteriorInspectedNEQ applies the NEQ predicate on the "interior_inspected" field.
func InteriorInspectedNEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldInteriorInspected, v))
}

// EngineCompartmentInspectedEQ applies the EQ predicate on the "engine_compartment_inspected" field.
func EngineCompartmentInspectedEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEngineCompartmentInspected, v))
}

// EngineCompartmentInspectedNEQ applies the NEQ predicate on the "engine_compartment_inspected" field.
func EngineCompartmentInspectedNEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldEngineCompartmentInspected, v))
}

// BrakeAndSteeringInspectedEQ applies the EQ predicate on the "brake_and_steering_inspected" field.
func BrakeAndSteeringInspectedEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldBrakeAndSteeringInspected, v))
}

// BrakeAndSteeringInspectedNEQ applies the NEQ predicate on the "brake_and_steering_inspected" field.
func BrakeAndSteeringInspectedNEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldBrakeAndSteeringInspected, v))
}

// EmergencyEquipmentInspectedEQ applies the EQ predicate on the "emergency_equipment_inspected" field.
func EmergencyEquipmentInspectedEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEmergencyEquipmentInspected, v))
}

// EmergencyEquipmentInspectedNEQ applies the NEQ predicate on the "emergency_equipment_inspected" field.
func EmergencyEquipmentInspectedNEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldEmergencyEquipmentInspected, v))
}

// FuelAndFluidsInspectedEQ applies the EQ predicate on the "fuel_and_fluids_inspected" field.
func FuelAndFluidsInspectedEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldFuelAndFluidsInspected, v))
}

// FuelAndFluidsInspectedNEQ applies the NEQ predicate on the "fuel_and_fluids_inspected" field.
func FuelAndFluidsInspectedNEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldFuelAndFluidsInspected, v))
}

// ScheduledEQ applies the EQ predicate on the "scheduled" field.
func ScheduledEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldScheduled, v))
}

// ScheduledNEQ applies the NEQ predicate on the "scheduled" field.
func ScheduledNEQ(v bool) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldScheduled, v))
}

// SeatLeftEQ applies the EQ predicate on the "seat_left" field.
func SeatLeftEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldSeatLeft, v))
}

// SeatLeftNEQ applies the NEQ predicate on the "seat_left" field.
func SeatLeftNEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldSeatLeft, v))
}

// SeatLeftIn applies the In predicate on the "seat_left" field.
func SeatLeftIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldSeatLeft, vs...))
}

// SeatLeftNotIn applies the NotIn predicate on the "seat_left" field.
func SeatLeftNotIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldSeatLeft, vs...))
}

// SeatLeftGT applies the GT predicate on the "seat_left" field.
func SeatLeftGT(v int) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldSeatLeft, v))
}

// SeatLeftGTE applies the GTE predicate on the "seat_left" field.
func SeatLeftGTE(v int) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldSeatLeft, v))
}

// SeatLeftLT applies the LT predicate on the "seat_left" field.
func SeatLeftLT(v int) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldSeatLeft, v))
}

// SeatLeftLTE applies the LTE predicate on the "seat_left" field.
func SeatLeftLTE(v int) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldSeatLeft, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Trip {
	return predicate.Trip(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Trip {
	return predicate.Trip(sql.FieldNotNull(FieldStatus))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDriver applies the HasEdge predicate on the "driver" edge.
func HasDriver() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DriverTable, DriverColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDriverWith applies the HasEdge predicate on the "driver" edge with a given conditions (other predicates).
func HasDriverWith(preds ...predicate.CompanyUser) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newDriverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVehicle applies the HasEdge predicate on the "vehicle" edge.
func HasVehicle() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VehicleTable, VehicleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVehicleWith applies the HasEdge predicate on the "vehicle" edge with a given conditions (other predicates).
func HasVehicleWith(preds ...predicate.Vehicle) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newVehicleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoute applies the HasEdge predicate on the "route" edge.
func HasRoute() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RouteTable, RouteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRouteWith applies the HasEdge predicate on the "route" edge with a given conditions (other predicates).
func HasRouteWith(preds ...predicate.Route) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newRouteStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBookings applies the HasEdge predicate on the "bookings" edge.
func HasBookings() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingsWith applies the HasEdge predicate on the "bookings" edge with a given conditions (other predicates).
func HasBookingsWith(preds ...predicate.Booking) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newBookingsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIncidents applies the HasEdge predicate on the "incidents" edge.
func HasIncidents() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IncidentsTable, IncidentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIncidentsWith applies the HasEdge predicate on the "incidents" edge with a given conditions (other predicates).
func HasIncidentsWith(preds ...predicate.Incident) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newIncidentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParcels applies the HasEdge predicate on the "parcels" edge.
func HasParcels() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParcelsTable, ParcelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParcelsWith applies the HasEdge predicate on the "parcels" edge with a given conditions (other predicates).
func HasParcelsWith(preds ...predicate.Parcel) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newParcelsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Trip) predicate.Trip {
	return predicate.Trip(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Trip) predicate.Trip {
	return predicate.Trip(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Trip) predicate.Trip {
	return predicate.Trip(sql.NotPredicates(p))
}
