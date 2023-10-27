// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/booking"
	"github.com/SeyramWood/ent/customerluggage"
)

// CustomerLuggageCreate is the builder for creating a CustomerLuggage entity.
type CustomerLuggageCreate struct {
	config
	mutation *CustomerLuggageMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (clc *CustomerLuggageCreate) SetCreatedAt(t time.Time) *CustomerLuggageCreate {
	clc.mutation.SetCreatedAt(t)
	return clc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (clc *CustomerLuggageCreate) SetNillableCreatedAt(t *time.Time) *CustomerLuggageCreate {
	if t != nil {
		clc.SetCreatedAt(*t)
	}
	return clc
}

// SetUpdatedAt sets the "updated_at" field.
func (clc *CustomerLuggageCreate) SetUpdatedAt(t time.Time) *CustomerLuggageCreate {
	clc.mutation.SetUpdatedAt(t)
	return clc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (clc *CustomerLuggageCreate) SetNillableUpdatedAt(t *time.Time) *CustomerLuggageCreate {
	if t != nil {
		clc.SetUpdatedAt(*t)
	}
	return clc
}

// SetBaggage sets the "baggage" field.
func (clc *CustomerLuggageCreate) SetBaggage(c customerluggage.Baggage) *CustomerLuggageCreate {
	clc.mutation.SetBaggage(c)
	return clc
}

// SetNillableBaggage sets the "baggage" field if the given value is not nil.
func (clc *CustomerLuggageCreate) SetNillableBaggage(c *customerluggage.Baggage) *CustomerLuggageCreate {
	if c != nil {
		clc.SetBaggage(*c)
	}
	return clc
}

// SetQuantity sets the "quantity" field.
func (clc *CustomerLuggageCreate) SetQuantity(i int) *CustomerLuggageCreate {
	clc.mutation.SetQuantity(i)
	return clc
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (clc *CustomerLuggageCreate) SetNillableQuantity(i *int) *CustomerLuggageCreate {
	if i != nil {
		clc.SetQuantity(*i)
	}
	return clc
}

// SetAmount sets the "amount" field.
func (clc *CustomerLuggageCreate) SetAmount(f float64) *CustomerLuggageCreate {
	clc.mutation.SetAmount(f)
	return clc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (clc *CustomerLuggageCreate) SetNillableAmount(f *float64) *CustomerLuggageCreate {
	if f != nil {
		clc.SetAmount(*f)
	}
	return clc
}

// SetBookingID sets the "booking" edge to the Booking entity by ID.
func (clc *CustomerLuggageCreate) SetBookingID(id int) *CustomerLuggageCreate {
	clc.mutation.SetBookingID(id)
	return clc
}

// SetNillableBookingID sets the "booking" edge to the Booking entity by ID if the given value is not nil.
func (clc *CustomerLuggageCreate) SetNillableBookingID(id *int) *CustomerLuggageCreate {
	if id != nil {
		clc = clc.SetBookingID(*id)
	}
	return clc
}

// SetBooking sets the "booking" edge to the Booking entity.
func (clc *CustomerLuggageCreate) SetBooking(b *Booking) *CustomerLuggageCreate {
	return clc.SetBookingID(b.ID)
}

// Mutation returns the CustomerLuggageMutation object of the builder.
func (clc *CustomerLuggageCreate) Mutation() *CustomerLuggageMutation {
	return clc.mutation
}

// Save creates the CustomerLuggage in the database.
func (clc *CustomerLuggageCreate) Save(ctx context.Context) (*CustomerLuggage, error) {
	clc.defaults()
	return withHooks(ctx, clc.sqlSave, clc.mutation, clc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (clc *CustomerLuggageCreate) SaveX(ctx context.Context) *CustomerLuggage {
	v, err := clc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clc *CustomerLuggageCreate) Exec(ctx context.Context) error {
	_, err := clc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clc *CustomerLuggageCreate) ExecX(ctx context.Context) {
	if err := clc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (clc *CustomerLuggageCreate) defaults() {
	if _, ok := clc.mutation.CreatedAt(); !ok {
		v := customerluggage.DefaultCreatedAt()
		clc.mutation.SetCreatedAt(v)
	}
	if _, ok := clc.mutation.UpdatedAt(); !ok {
		v := customerluggage.DefaultUpdatedAt()
		clc.mutation.SetUpdatedAt(v)
	}
	if _, ok := clc.mutation.Quantity(); !ok {
		v := customerluggage.DefaultQuantity
		clc.mutation.SetQuantity(v)
	}
	if _, ok := clc.mutation.Amount(); !ok {
		v := customerluggage.DefaultAmount
		clc.mutation.SetAmount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (clc *CustomerLuggageCreate) check() error {
	if _, ok := clc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CustomerLuggage.created_at"`)}
	}
	if _, ok := clc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CustomerLuggage.updated_at"`)}
	}
	if v, ok := clc.mutation.Baggage(); ok {
		if err := customerluggage.BaggageValidator(v); err != nil {
			return &ValidationError{Name: "baggage", err: fmt.Errorf(`ent: validator failed for field "CustomerLuggage.baggage": %w`, err)}
		}
	}
	if _, ok := clc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "CustomerLuggage.quantity"`)}
	}
	if _, ok := clc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "CustomerLuggage.amount"`)}
	}
	return nil
}

func (clc *CustomerLuggageCreate) sqlSave(ctx context.Context) (*CustomerLuggage, error) {
	if err := clc.check(); err != nil {
		return nil, err
	}
	_node, _spec := clc.createSpec()
	if err := sqlgraph.CreateNode(ctx, clc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	clc.mutation.id = &_node.ID
	clc.mutation.done = true
	return _node, nil
}

func (clc *CustomerLuggageCreate) createSpec() (*CustomerLuggage, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomerLuggage{config: clc.config}
		_spec = sqlgraph.NewCreateSpec(customerluggage.Table, sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt))
	)
	if value, ok := clc.mutation.CreatedAt(); ok {
		_spec.SetField(customerluggage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := clc.mutation.UpdatedAt(); ok {
		_spec.SetField(customerluggage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := clc.mutation.Baggage(); ok {
		_spec.SetField(customerluggage.FieldBaggage, field.TypeEnum, value)
		_node.Baggage = value
	}
	if value, ok := clc.mutation.Quantity(); ok {
		_spec.SetField(customerluggage.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := clc.mutation.Amount(); ok {
		_spec.SetField(customerluggage.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if nodes := clc.mutation.BookingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customerluggage.BookingTable,
			Columns: []string{customerluggage.BookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.booking_luggages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerLuggageCreateBulk is the builder for creating many CustomerLuggage entities in bulk.
type CustomerLuggageCreateBulk struct {
	config
	err      error
	builders []*CustomerLuggageCreate
}

// Save creates the CustomerLuggage entities in the database.
func (clcb *CustomerLuggageCreateBulk) Save(ctx context.Context) ([]*CustomerLuggage, error) {
	if clcb.err != nil {
		return nil, clcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(clcb.builders))
	nodes := make([]*CustomerLuggage, len(clcb.builders))
	mutators := make([]Mutator, len(clcb.builders))
	for i := range clcb.builders {
		func(i int, root context.Context) {
			builder := clcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerLuggageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, clcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, clcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, clcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (clcb *CustomerLuggageCreateBulk) SaveX(ctx context.Context) []*CustomerLuggage {
	v, err := clcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clcb *CustomerLuggageCreateBulk) Exec(ctx context.Context) error {
	_, err := clcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clcb *CustomerLuggageCreateBulk) ExecX(ctx context.Context) {
	if err := clcb.Exec(ctx); err != nil {
		panic(err)
	}
}
