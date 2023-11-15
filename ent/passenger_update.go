// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/passenger"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// PassengerUpdate is the builder for updating Passenger entities.
type PassengerUpdate struct {
	config
	hooks     []Hook
	mutation  *PassengerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PassengerUpdate builder.
func (pu *PassengerUpdate) Where(ps ...predicate.Passenger) *PassengerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PassengerUpdate) SetUpdatedAt(t time.Time) *PassengerUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetFullName sets the "full_name" field.
func (pu *PassengerUpdate) SetFullName(s string) *PassengerUpdate {
	pu.mutation.SetFullName(s)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PassengerUpdate) SetAmount(f float64) *PassengerUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PassengerUpdate) SetNillableAmount(f *float64) *PassengerUpdate {
	if f != nil {
		pu.SetAmount(*f)
	}
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *PassengerUpdate) AddAmount(f float64) *PassengerUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetMaturity sets the "maturity" field.
func (pu *PassengerUpdate) SetMaturity(pa passenger.Maturity) *PassengerUpdate {
	pu.mutation.SetMaturity(pa)
	return pu
}

// SetNillableMaturity sets the "maturity" field if the given value is not nil.
func (pu *PassengerUpdate) SetNillableMaturity(pa *passenger.Maturity) *PassengerUpdate {
	if pa != nil {
		pu.SetMaturity(*pa)
	}
	return pu
}

// SetGender sets the "gender" field.
func (pu *PassengerUpdate) SetGender(pa passenger.Gender) *PassengerUpdate {
	pu.mutation.SetGender(pa)
	return pu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (pu *PassengerUpdate) SetNillableGender(pa *passenger.Gender) *PassengerUpdate {
	if pa != nil {
		pu.SetGender(*pa)
	}
	return pu
}

// ClearGender clears the value of the "gender" field.
func (pu *PassengerUpdate) ClearGender() *PassengerUpdate {
	pu.mutation.ClearGender()
	return pu
}

// SetBookingID sets the "booking" edge to the Booking entity by ID.
func (pu *PassengerUpdate) SetBookingID(id int) *PassengerUpdate {
	pu.mutation.SetBookingID(id)
	return pu
}

// SetBooking sets the "booking" edge to the Booking entity.
func (pu *PassengerUpdate) SetBooking(b *Booking) *PassengerUpdate {
	return pu.SetBookingID(b.ID)
}

// Mutation returns the PassengerMutation object of the builder.
func (pu *PassengerUpdate) Mutation() *PassengerMutation {
	return pu.mutation
}

// ClearBooking clears the "booking" edge to the Booking entity.
func (pu *PassengerUpdate) ClearBooking() *PassengerUpdate {
	pu.mutation.ClearBooking()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PassengerUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PassengerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PassengerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PassengerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PassengerUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := passenger.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PassengerUpdate) check() error {
	if v, ok := pu.mutation.FullName(); ok {
		if err := passenger.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Passenger.full_name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Maturity(); ok {
		if err := passenger.MaturityValidator(v); err != nil {
			return &ValidationError{Name: "maturity", err: fmt.Errorf(`ent: validator failed for field "Passenger.maturity": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Gender(); ok {
		if err := passenger.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Passenger.gender": %w`, err)}
		}
	}
	if _, ok := pu.mutation.BookingID(); pu.mutation.BookingCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Passenger.booking"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PassengerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PassengerUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PassengerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(passenger.Table, passenger.Columns, sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(passenger.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.FullName(); ok {
		_spec.SetField(passenger.FieldFullName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(passenger.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(passenger.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Maturity(); ok {
		_spec.SetField(passenger.FieldMaturity, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Gender(); ok {
		_spec.SetField(passenger.FieldGender, field.TypeEnum, value)
	}
	if pu.mutation.GenderCleared() {
		_spec.ClearField(passenger.FieldGender, field.TypeEnum)
	}
	if pu.mutation.BookingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passenger.BookingTable,
			Columns: []string{passenger.BookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BookingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passenger.BookingTable,
			Columns: []string{passenger.BookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passenger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PassengerUpdateOne is the builder for updating a single Passenger entity.
type PassengerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PassengerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PassengerUpdateOne) SetUpdatedAt(t time.Time) *PassengerUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetFullName sets the "full_name" field.
func (puo *PassengerUpdateOne) SetFullName(s string) *PassengerUpdateOne {
	puo.mutation.SetFullName(s)
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PassengerUpdateOne) SetAmount(f float64) *PassengerUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PassengerUpdateOne) SetNillableAmount(f *float64) *PassengerUpdateOne {
	if f != nil {
		puo.SetAmount(*f)
	}
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *PassengerUpdateOne) AddAmount(f float64) *PassengerUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetMaturity sets the "maturity" field.
func (puo *PassengerUpdateOne) SetMaturity(pa passenger.Maturity) *PassengerUpdateOne {
	puo.mutation.SetMaturity(pa)
	return puo
}

// SetNillableMaturity sets the "maturity" field if the given value is not nil.
func (puo *PassengerUpdateOne) SetNillableMaturity(pa *passenger.Maturity) *PassengerUpdateOne {
	if pa != nil {
		puo.SetMaturity(*pa)
	}
	return puo
}

// SetGender sets the "gender" field.
func (puo *PassengerUpdateOne) SetGender(pa passenger.Gender) *PassengerUpdateOne {
	puo.mutation.SetGender(pa)
	return puo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (puo *PassengerUpdateOne) SetNillableGender(pa *passenger.Gender) *PassengerUpdateOne {
	if pa != nil {
		puo.SetGender(*pa)
	}
	return puo
}

// ClearGender clears the value of the "gender" field.
func (puo *PassengerUpdateOne) ClearGender() *PassengerUpdateOne {
	puo.mutation.ClearGender()
	return puo
}

// SetBookingID sets the "booking" edge to the Booking entity by ID.
func (puo *PassengerUpdateOne) SetBookingID(id int) *PassengerUpdateOne {
	puo.mutation.SetBookingID(id)
	return puo
}

// SetBooking sets the "booking" edge to the Booking entity.
func (puo *PassengerUpdateOne) SetBooking(b *Booking) *PassengerUpdateOne {
	return puo.SetBookingID(b.ID)
}

// Mutation returns the PassengerMutation object of the builder.
func (puo *PassengerUpdateOne) Mutation() *PassengerMutation {
	return puo.mutation
}

// ClearBooking clears the "booking" edge to the Booking entity.
func (puo *PassengerUpdateOne) ClearBooking() *PassengerUpdateOne {
	puo.mutation.ClearBooking()
	return puo
}

// Where appends a list predicates to the PassengerUpdate builder.
func (puo *PassengerUpdateOne) Where(ps ...predicate.Passenger) *PassengerUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PassengerUpdateOne) Select(field string, fields ...string) *PassengerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Passenger entity.
func (puo *PassengerUpdateOne) Save(ctx context.Context) (*Passenger, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PassengerUpdateOne) SaveX(ctx context.Context) *Passenger {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PassengerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PassengerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PassengerUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := passenger.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PassengerUpdateOne) check() error {
	if v, ok := puo.mutation.FullName(); ok {
		if err := passenger.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Passenger.full_name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Maturity(); ok {
		if err := passenger.MaturityValidator(v); err != nil {
			return &ValidationError{Name: "maturity", err: fmt.Errorf(`ent: validator failed for field "Passenger.maturity": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Gender(); ok {
		if err := passenger.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Passenger.gender": %w`, err)}
		}
	}
	if _, ok := puo.mutation.BookingID(); puo.mutation.BookingCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Passenger.booking"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PassengerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PassengerUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PassengerUpdateOne) sqlSave(ctx context.Context) (_node *Passenger, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(passenger.Table, passenger.Columns, sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Passenger.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passenger.FieldID)
		for _, f := range fields {
			if !passenger.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passenger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(passenger.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.FullName(); ok {
		_spec.SetField(passenger.FieldFullName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(passenger.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(passenger.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Maturity(); ok {
		_spec.SetField(passenger.FieldMaturity, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Gender(); ok {
		_spec.SetField(passenger.FieldGender, field.TypeEnum, value)
	}
	if puo.mutation.GenderCleared() {
		_spec.ClearField(passenger.FieldGender, field.TypeEnum)
	}
	if puo.mutation.BookingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passenger.BookingTable,
			Columns: []string{passenger.BookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BookingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passenger.BookingTable,
			Columns: []string{passenger.BookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Passenger{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passenger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
