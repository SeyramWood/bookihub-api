// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/companyuser"
	"github.com/SeyramWood/ent/parcel"
	"github.com/SeyramWood/ent/parcelimage"
	"github.com/SeyramWood/ent/trip"
)

// ParcelCreate is the builder for creating a Parcel entity.
type ParcelCreate struct {
	config
	mutation *ParcelMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *ParcelCreate) SetCreatedAt(t time.Time) *ParcelCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ParcelCreate) SetNillableCreatedAt(t *time.Time) *ParcelCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ParcelCreate) SetUpdatedAt(t time.Time) *ParcelCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ParcelCreate) SetNillableUpdatedAt(t *time.Time) *ParcelCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetParcelCode sets the "parcel_code" field.
func (pc *ParcelCreate) SetParcelCode(s string) *ParcelCreate {
	pc.mutation.SetParcelCode(s)
	return pc
}

// SetSenderName sets the "sender_name" field.
func (pc *ParcelCreate) SetSenderName(s string) *ParcelCreate {
	pc.mutation.SetSenderName(s)
	return pc
}

// SetSenderPhone sets the "sender_phone" field.
func (pc *ParcelCreate) SetSenderPhone(s string) *ParcelCreate {
	pc.mutation.SetSenderPhone(s)
	return pc
}

// SetRecipientName sets the "recipient_name" field.
func (pc *ParcelCreate) SetRecipientName(s string) *ParcelCreate {
	pc.mutation.SetRecipientName(s)
	return pc
}

// SetRecipientPhone sets the "recipient_phone" field.
func (pc *ParcelCreate) SetRecipientPhone(s string) *ParcelCreate {
	pc.mutation.SetRecipientPhone(s)
	return pc
}

// SetRecipientLocation sets the "recipient_location" field.
func (pc *ParcelCreate) SetRecipientLocation(s string) *ParcelCreate {
	pc.mutation.SetRecipientLocation(s)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *ParcelCreate) SetAmount(f float64) *ParcelCreate {
	pc.mutation.SetAmount(f)
	return pc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pc *ParcelCreate) SetNillableAmount(f *float64) *ParcelCreate {
	if f != nil {
		pc.SetAmount(*f)
	}
	return pc
}

// SetPaidAt sets the "paid_at" field.
func (pc *ParcelCreate) SetPaidAt(t time.Time) *ParcelCreate {
	pc.mutation.SetPaidAt(t)
	return pc
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (pc *ParcelCreate) SetNillablePaidAt(t *time.Time) *ParcelCreate {
	if t != nil {
		pc.SetPaidAt(*t)
	}
	return pc
}

// SetTansType sets the "tans_type" field.
func (pc *ParcelCreate) SetTansType(pt parcel.TansType) *ParcelCreate {
	pc.mutation.SetTansType(pt)
	return pc
}

// SetNillableTansType sets the "tans_type" field if the given value is not nil.
func (pc *ParcelCreate) SetNillableTansType(pt *parcel.TansType) *ParcelCreate {
	if pt != nil {
		pc.SetTansType(*pt)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *ParcelCreate) SetStatus(pa parcel.Status) *ParcelCreate {
	pc.mutation.SetStatus(pa)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *ParcelCreate) SetNillableStatus(pa *parcel.Status) *ParcelCreate {
	if pa != nil {
		pc.SetStatus(*pa)
	}
	return pc
}

// AddImageIDs adds the "images" edge to the ParcelImage entity by IDs.
func (pc *ParcelCreate) AddImageIDs(ids ...int) *ParcelCreate {
	pc.mutation.AddImageIDs(ids...)
	return pc
}

// AddImages adds the "images" edges to the ParcelImage entity.
func (pc *ParcelCreate) AddImages(p ...*ParcelImage) *ParcelCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddImageIDs(ids...)
}

// SetTripID sets the "trip" edge to the Trip entity by ID.
func (pc *ParcelCreate) SetTripID(id int) *ParcelCreate {
	pc.mutation.SetTripID(id)
	return pc
}

// SetNillableTripID sets the "trip" edge to the Trip entity by ID if the given value is not nil.
func (pc *ParcelCreate) SetNillableTripID(id *int) *ParcelCreate {
	if id != nil {
		pc = pc.SetTripID(*id)
	}
	return pc
}

// SetTrip sets the "trip" edge to the Trip entity.
func (pc *ParcelCreate) SetTrip(t *Trip) *ParcelCreate {
	return pc.SetTripID(t.ID)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pc *ParcelCreate) SetCompanyID(id int) *ParcelCreate {
	pc.mutation.SetCompanyID(id)
	return pc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (pc *ParcelCreate) SetNillableCompanyID(id *int) *ParcelCreate {
	if id != nil {
		pc = pc.SetCompanyID(*id)
	}
	return pc
}

// SetCompany sets the "company" edge to the Company entity.
func (pc *ParcelCreate) SetCompany(c *Company) *ParcelCreate {
	return pc.SetCompanyID(c.ID)
}

// SetDriverID sets the "driver" edge to the CompanyUser entity by ID.
func (pc *ParcelCreate) SetDriverID(id int) *ParcelCreate {
	pc.mutation.SetDriverID(id)
	return pc
}

// SetNillableDriverID sets the "driver" edge to the CompanyUser entity by ID if the given value is not nil.
func (pc *ParcelCreate) SetNillableDriverID(id *int) *ParcelCreate {
	if id != nil {
		pc = pc.SetDriverID(*id)
	}
	return pc
}

// SetDriver sets the "driver" edge to the CompanyUser entity.
func (pc *ParcelCreate) SetDriver(c *CompanyUser) *ParcelCreate {
	return pc.SetDriverID(c.ID)
}

// Mutation returns the ParcelMutation object of the builder.
func (pc *ParcelCreate) Mutation() *ParcelMutation {
	return pc.mutation
}

// Save creates the Parcel in the database.
func (pc *ParcelCreate) Save(ctx context.Context) (*Parcel, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ParcelCreate) SaveX(ctx context.Context) *Parcel {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ParcelCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ParcelCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ParcelCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := parcel.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := parcel.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Amount(); !ok {
		v := parcel.DefaultAmount
		pc.mutation.SetAmount(v)
	}
	if _, ok := pc.mutation.TansType(); !ok {
		v := parcel.DefaultTansType
		pc.mutation.SetTansType(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := parcel.DefaultStatus
		pc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ParcelCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Parcel.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Parcel.updated_at"`)}
	}
	if _, ok := pc.mutation.ParcelCode(); !ok {
		return &ValidationError{Name: "parcel_code", err: errors.New(`ent: missing required field "Parcel.parcel_code"`)}
	}
	if v, ok := pc.mutation.ParcelCode(); ok {
		if err := parcel.ParcelCodeValidator(v); err != nil {
			return &ValidationError{Name: "parcel_code", err: fmt.Errorf(`ent: validator failed for field "Parcel.parcel_code": %w`, err)}
		}
	}
	if _, ok := pc.mutation.SenderName(); !ok {
		return &ValidationError{Name: "sender_name", err: errors.New(`ent: missing required field "Parcel.sender_name"`)}
	}
	if v, ok := pc.mutation.SenderName(); ok {
		if err := parcel.SenderNameValidator(v); err != nil {
			return &ValidationError{Name: "sender_name", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.SenderPhone(); !ok {
		return &ValidationError{Name: "sender_phone", err: errors.New(`ent: missing required field "Parcel.sender_phone"`)}
	}
	if v, ok := pc.mutation.SenderPhone(); ok {
		if err := parcel.SenderPhoneValidator(v); err != nil {
			return &ValidationError{Name: "sender_phone", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_phone": %w`, err)}
		}
	}
	if _, ok := pc.mutation.RecipientName(); !ok {
		return &ValidationError{Name: "recipient_name", err: errors.New(`ent: missing required field "Parcel.recipient_name"`)}
	}
	if v, ok := pc.mutation.RecipientName(); ok {
		if err := parcel.RecipientNameValidator(v); err != nil {
			return &ValidationError{Name: "recipient_name", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.RecipientPhone(); !ok {
		return &ValidationError{Name: "recipient_phone", err: errors.New(`ent: missing required field "Parcel.recipient_phone"`)}
	}
	if v, ok := pc.mutation.RecipientPhone(); ok {
		if err := parcel.RecipientPhoneValidator(v); err != nil {
			return &ValidationError{Name: "recipient_phone", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_phone": %w`, err)}
		}
	}
	if _, ok := pc.mutation.RecipientLocation(); !ok {
		return &ValidationError{Name: "recipient_location", err: errors.New(`ent: missing required field "Parcel.recipient_location"`)}
	}
	if v, ok := pc.mutation.RecipientLocation(); ok {
		if err := parcel.RecipientLocationValidator(v); err != nil {
			return &ValidationError{Name: "recipient_location", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_location": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Parcel.amount"`)}
	}
	if _, ok := pc.mutation.TansType(); !ok {
		return &ValidationError{Name: "tans_type", err: errors.New(`ent: missing required field "Parcel.tans_type"`)}
	}
	if v, ok := pc.mutation.TansType(); ok {
		if err := parcel.TansTypeValidator(v); err != nil {
			return &ValidationError{Name: "tans_type", err: fmt.Errorf(`ent: validator failed for field "Parcel.tans_type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Parcel.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := parcel.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Parcel.status": %w`, err)}
		}
	}
	return nil
}

func (pc *ParcelCreate) sqlSave(ctx context.Context) (*Parcel, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ParcelCreate) createSpec() (*Parcel, *sqlgraph.CreateSpec) {
	var (
		_node = &Parcel{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(parcel.Table, sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(parcel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(parcel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.ParcelCode(); ok {
		_spec.SetField(parcel.FieldParcelCode, field.TypeString, value)
		_node.ParcelCode = value
	}
	if value, ok := pc.mutation.SenderName(); ok {
		_spec.SetField(parcel.FieldSenderName, field.TypeString, value)
		_node.SenderName = value
	}
	if value, ok := pc.mutation.SenderPhone(); ok {
		_spec.SetField(parcel.FieldSenderPhone, field.TypeString, value)
		_node.SenderPhone = value
	}
	if value, ok := pc.mutation.RecipientName(); ok {
		_spec.SetField(parcel.FieldRecipientName, field.TypeString, value)
		_node.RecipientName = value
	}
	if value, ok := pc.mutation.RecipientPhone(); ok {
		_spec.SetField(parcel.FieldRecipientPhone, field.TypeString, value)
		_node.RecipientPhone = value
	}
	if value, ok := pc.mutation.RecipientLocation(); ok {
		_spec.SetField(parcel.FieldRecipientLocation, field.TypeString, value)
		_node.RecipientLocation = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(parcel.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.PaidAt(); ok {
		_spec.SetField(parcel.FieldPaidAt, field.TypeTime, value)
		_node.PaidAt = value
	}
	if value, ok := pc.mutation.TansType(); ok {
		_spec.SetField(parcel.FieldTansType, field.TypeEnum, value)
		_node.TansType = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(parcel.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := pc.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   parcel.ImagesTable,
			Columns: []string{parcel.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.TripTable,
			Columns: []string{parcel.TripColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.trip_parcels = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.CompanyTable,
			Columns: []string{parcel.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_parcels = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DriverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.DriverTable,
			Columns: []string{parcel.DriverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_user_parcels = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ParcelCreateBulk is the builder for creating many Parcel entities in bulk.
type ParcelCreateBulk struct {
	config
	err      error
	builders []*ParcelCreate
}

// Save creates the Parcel entities in the database.
func (pcb *ParcelCreateBulk) Save(ctx context.Context) ([]*Parcel, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Parcel, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ParcelMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ParcelCreateBulk) SaveX(ctx context.Context) []*Parcel {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ParcelCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ParcelCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
