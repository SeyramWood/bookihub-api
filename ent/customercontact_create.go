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
	"github.com/SeyramWood/ent/customercontact"
)

// CustomerContactCreate is the builder for creating a CustomerContact entity.
type CustomerContactCreate struct {
	config
	mutation *CustomerContactMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ccc *CustomerContactCreate) SetCreatedAt(t time.Time) *CustomerContactCreate {
	ccc.mutation.SetCreatedAt(t)
	return ccc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccc *CustomerContactCreate) SetNillableCreatedAt(t *time.Time) *CustomerContactCreate {
	if t != nil {
		ccc.SetCreatedAt(*t)
	}
	return ccc
}

// SetUpdatedAt sets the "updated_at" field.
func (ccc *CustomerContactCreate) SetUpdatedAt(t time.Time) *CustomerContactCreate {
	ccc.mutation.SetUpdatedAt(t)
	return ccc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccc *CustomerContactCreate) SetNillableUpdatedAt(t *time.Time) *CustomerContactCreate {
	if t != nil {
		ccc.SetUpdatedAt(*t)
	}
	return ccc
}

// SetFullName sets the "full_name" field.
func (ccc *CustomerContactCreate) SetFullName(s string) *CustomerContactCreate {
	ccc.mutation.SetFullName(s)
	return ccc
}

// SetEmail sets the "email" field.
func (ccc *CustomerContactCreate) SetEmail(s string) *CustomerContactCreate {
	ccc.mutation.SetEmail(s)
	return ccc
}

// SetPhone sets the "phone" field.
func (ccc *CustomerContactCreate) SetPhone(s string) *CustomerContactCreate {
	ccc.mutation.SetPhone(s)
	return ccc
}

// SetBookingID sets the "booking" edge to the Booking entity by ID.
func (ccc *CustomerContactCreate) SetBookingID(id int) *CustomerContactCreate {
	ccc.mutation.SetBookingID(id)
	return ccc
}

// SetNillableBookingID sets the "booking" edge to the Booking entity by ID if the given value is not nil.
func (ccc *CustomerContactCreate) SetNillableBookingID(id *int) *CustomerContactCreate {
	if id != nil {
		ccc = ccc.SetBookingID(*id)
	}
	return ccc
}

// SetBooking sets the "booking" edge to the Booking entity.
func (ccc *CustomerContactCreate) SetBooking(b *Booking) *CustomerContactCreate {
	return ccc.SetBookingID(b.ID)
}

// Mutation returns the CustomerContactMutation object of the builder.
func (ccc *CustomerContactCreate) Mutation() *CustomerContactMutation {
	return ccc.mutation
}

// Save creates the CustomerContact in the database.
func (ccc *CustomerContactCreate) Save(ctx context.Context) (*CustomerContact, error) {
	ccc.defaults()
	return withHooks(ctx, ccc.sqlSave, ccc.mutation, ccc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *CustomerContactCreate) SaveX(ctx context.Context) *CustomerContact {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *CustomerContactCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *CustomerContactCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccc *CustomerContactCreate) defaults() {
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		v := customercontact.DefaultCreatedAt()
		ccc.mutation.SetCreatedAt(v)
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		v := customercontact.DefaultUpdatedAt()
		ccc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccc *CustomerContactCreate) check() error {
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CustomerContact.created_at"`)}
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CustomerContact.updated_at"`)}
	}
	if _, ok := ccc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`ent: missing required field "CustomerContact.full_name"`)}
	}
	if v, ok := ccc.mutation.FullName(); ok {
		if err := customercontact.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "CustomerContact.full_name": %w`, err)}
		}
	}
	if _, ok := ccc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "CustomerContact.email"`)}
	}
	if v, ok := ccc.mutation.Email(); ok {
		if err := customercontact.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "CustomerContact.email": %w`, err)}
		}
	}
	if _, ok := ccc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "CustomerContact.phone"`)}
	}
	if v, ok := ccc.mutation.Phone(); ok {
		if err := customercontact.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "CustomerContact.phone": %w`, err)}
		}
	}
	return nil
}

func (ccc *CustomerContactCreate) sqlSave(ctx context.Context) (*CustomerContact, error) {
	if err := ccc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ccc.mutation.id = &_node.ID
	ccc.mutation.done = true
	return _node, nil
}

func (ccc *CustomerContactCreate) createSpec() (*CustomerContact, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomerContact{config: ccc.config}
		_spec = sqlgraph.NewCreateSpec(customercontact.Table, sqlgraph.NewFieldSpec(customercontact.FieldID, field.TypeInt))
	)
	if value, ok := ccc.mutation.CreatedAt(); ok {
		_spec.SetField(customercontact.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ccc.mutation.UpdatedAt(); ok {
		_spec.SetField(customercontact.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ccc.mutation.FullName(); ok {
		_spec.SetField(customercontact.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := ccc.mutation.Email(); ok {
		_spec.SetField(customercontact.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ccc.mutation.Phone(); ok {
		_spec.SetField(customercontact.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if nodes := ccc.mutation.BookingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   customercontact.BookingTable,
			Columns: []string{customercontact.BookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.booking_contact = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerContactCreateBulk is the builder for creating many CustomerContact entities in bulk.
type CustomerContactCreateBulk struct {
	config
	err      error
	builders []*CustomerContactCreate
}

// Save creates the CustomerContact entities in the database.
func (cccb *CustomerContactCreateBulk) Save(ctx context.Context) ([]*CustomerContact, error) {
	if cccb.err != nil {
		return nil, cccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*CustomerContact, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerContactMutation)
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
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *CustomerContactCreateBulk) SaveX(ctx context.Context) []*CustomerContact {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *CustomerContactCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *CustomerContactCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}
