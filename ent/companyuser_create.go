// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/incident"
	"github.com/SeyramWood/bookibus/ent/notification"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/trip"
	"github.com/SeyramWood/bookibus/ent/user"
)

// CompanyUserCreate is the builder for creating a CompanyUser entity.
type CompanyUserCreate struct {
	config
	mutation *CompanyUserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cuc *CompanyUserCreate) SetCreatedAt(t time.Time) *CompanyUserCreate {
	cuc.mutation.SetCreatedAt(t)
	return cuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableCreatedAt(t *time.Time) *CompanyUserCreate {
	if t != nil {
		cuc.SetCreatedAt(*t)
	}
	return cuc
}

// SetUpdatedAt sets the "updated_at" field.
func (cuc *CompanyUserCreate) SetUpdatedAt(t time.Time) *CompanyUserCreate {
	cuc.mutation.SetUpdatedAt(t)
	return cuc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableUpdatedAt(t *time.Time) *CompanyUserCreate {
	if t != nil {
		cuc.SetUpdatedAt(*t)
	}
	return cuc
}

// SetLastName sets the "last_name" field.
func (cuc *CompanyUserCreate) SetLastName(s string) *CompanyUserCreate {
	cuc.mutation.SetLastName(s)
	return cuc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableLastName(s *string) *CompanyUserCreate {
	if s != nil {
		cuc.SetLastName(*s)
	}
	return cuc
}

// SetOtherName sets the "other_name" field.
func (cuc *CompanyUserCreate) SetOtherName(s string) *CompanyUserCreate {
	cuc.mutation.SetOtherName(s)
	return cuc
}

// SetNillableOtherName sets the "other_name" field if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableOtherName(s *string) *CompanyUserCreate {
	if s != nil {
		cuc.SetOtherName(*s)
	}
	return cuc
}

// SetPhone sets the "phone" field.
func (cuc *CompanyUserCreate) SetPhone(s string) *CompanyUserCreate {
	cuc.mutation.SetPhone(s)
	return cuc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillablePhone(s *string) *CompanyUserCreate {
	if s != nil {
		cuc.SetPhone(*s)
	}
	return cuc
}

// SetOtherPhone sets the "other_phone" field.
func (cuc *CompanyUserCreate) SetOtherPhone(s string) *CompanyUserCreate {
	cuc.mutation.SetOtherPhone(s)
	return cuc
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableOtherPhone(s *string) *CompanyUserCreate {
	if s != nil {
		cuc.SetOtherPhone(*s)
	}
	return cuc
}

// SetUserRole sets the "user_role" field.
func (cuc *CompanyUserCreate) SetUserRole(cr companyuser.UserRole) *CompanyUserCreate {
	cuc.mutation.SetUserRole(cr)
	return cuc
}

// SetNillableUserRole sets the "user_role" field if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableUserRole(cr *companyuser.UserRole) *CompanyUserCreate {
	if cr != nil {
		cuc.SetUserRole(*cr)
	}
	return cuc
}

// SetProfileID sets the "profile" edge to the User entity by ID.
func (cuc *CompanyUserCreate) SetProfileID(id int) *CompanyUserCreate {
	cuc.mutation.SetProfileID(id)
	return cuc
}

// SetNillableProfileID sets the "profile" edge to the User entity by ID if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableProfileID(id *int) *CompanyUserCreate {
	if id != nil {
		cuc = cuc.SetProfileID(*id)
	}
	return cuc
}

// SetProfile sets the "profile" edge to the User entity.
func (cuc *CompanyUserCreate) SetProfile(u *User) *CompanyUserCreate {
	return cuc.SetProfileID(u.ID)
}

// AddTripIDs adds the "trips" edge to the Trip entity by IDs.
func (cuc *CompanyUserCreate) AddTripIDs(ids ...int) *CompanyUserCreate {
	cuc.mutation.AddTripIDs(ids...)
	return cuc
}

// AddTrips adds the "trips" edges to the Trip entity.
func (cuc *CompanyUserCreate) AddTrips(t ...*Trip) *CompanyUserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuc.AddTripIDs(ids...)
}

// AddIncidentIDs adds the "incidents" edge to the Incident entity by IDs.
func (cuc *CompanyUserCreate) AddIncidentIDs(ids ...int) *CompanyUserCreate {
	cuc.mutation.AddIncidentIDs(ids...)
	return cuc
}

// AddIncidents adds the "incidents" edges to the Incident entity.
func (cuc *CompanyUserCreate) AddIncidents(i ...*Incident) *CompanyUserCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuc.AddIncidentIDs(ids...)
}

// AddParcelIDs adds the "parcels" edge to the Parcel entity by IDs.
func (cuc *CompanyUserCreate) AddParcelIDs(ids ...int) *CompanyUserCreate {
	cuc.mutation.AddParcelIDs(ids...)
	return cuc
}

// AddParcels adds the "parcels" edges to the Parcel entity.
func (cuc *CompanyUserCreate) AddParcels(p ...*Parcel) *CompanyUserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuc.AddParcelIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (cuc *CompanyUserCreate) AddNotificationIDs(ids ...int) *CompanyUserCreate {
	cuc.mutation.AddNotificationIDs(ids...)
	return cuc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (cuc *CompanyUserCreate) AddNotifications(n ...*Notification) *CompanyUserCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cuc.AddNotificationIDs(ids...)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (cuc *CompanyUserCreate) SetCompanyID(id int) *CompanyUserCreate {
	cuc.mutation.SetCompanyID(id)
	return cuc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (cuc *CompanyUserCreate) SetNillableCompanyID(id *int) *CompanyUserCreate {
	if id != nil {
		cuc = cuc.SetCompanyID(*id)
	}
	return cuc
}

// SetCompany sets the "company" edge to the Company entity.
func (cuc *CompanyUserCreate) SetCompany(c *Company) *CompanyUserCreate {
	return cuc.SetCompanyID(c.ID)
}

// Mutation returns the CompanyUserMutation object of the builder.
func (cuc *CompanyUserCreate) Mutation() *CompanyUserMutation {
	return cuc.mutation
}

// Save creates the CompanyUser in the database.
func (cuc *CompanyUserCreate) Save(ctx context.Context) (*CompanyUser, error) {
	cuc.defaults()
	return withHooks(ctx, cuc.sqlSave, cuc.mutation, cuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cuc *CompanyUserCreate) SaveX(ctx context.Context) *CompanyUser {
	v, err := cuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cuc *CompanyUserCreate) Exec(ctx context.Context) error {
	_, err := cuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuc *CompanyUserCreate) ExecX(ctx context.Context) {
	if err := cuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuc *CompanyUserCreate) defaults() {
	if _, ok := cuc.mutation.CreatedAt(); !ok {
		v := companyuser.DefaultCreatedAt()
		cuc.mutation.SetCreatedAt(v)
	}
	if _, ok := cuc.mutation.UpdatedAt(); !ok {
		v := companyuser.DefaultUpdatedAt()
		cuc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cuc.mutation.UserRole(); !ok {
		v := companyuser.DefaultUserRole
		cuc.mutation.SetUserRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuc *CompanyUserCreate) check() error {
	if _, ok := cuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CompanyUser.created_at"`)}
	}
	if _, ok := cuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CompanyUser.updated_at"`)}
	}
	if _, ok := cuc.mutation.UserRole(); !ok {
		return &ValidationError{Name: "user_role", err: errors.New(`ent: missing required field "CompanyUser.user_role"`)}
	}
	if v, ok := cuc.mutation.UserRole(); ok {
		if err := companyuser.UserRoleValidator(v); err != nil {
			return &ValidationError{Name: "user_role", err: fmt.Errorf(`ent: validator failed for field "CompanyUser.user_role": %w`, err)}
		}
	}
	return nil
}

func (cuc *CompanyUserCreate) sqlSave(ctx context.Context) (*CompanyUser, error) {
	if err := cuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cuc.mutation.id = &_node.ID
	cuc.mutation.done = true
	return _node, nil
}

func (cuc *CompanyUserCreate) createSpec() (*CompanyUser, *sqlgraph.CreateSpec) {
	var (
		_node = &CompanyUser{config: cuc.config}
		_spec = sqlgraph.NewCreateSpec(companyuser.Table, sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt))
	)
	if value, ok := cuc.mutation.CreatedAt(); ok {
		_spec.SetField(companyuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cuc.mutation.UpdatedAt(); ok {
		_spec.SetField(companyuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cuc.mutation.LastName(); ok {
		_spec.SetField(companyuser.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := cuc.mutation.OtherName(); ok {
		_spec.SetField(companyuser.FieldOtherName, field.TypeString, value)
		_node.OtherName = value
	}
	if value, ok := cuc.mutation.Phone(); ok {
		_spec.SetField(companyuser.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := cuc.mutation.OtherPhone(); ok {
		_spec.SetField(companyuser.FieldOtherPhone, field.TypeString, value)
		_node.OtherPhone = value
	}
	if value, ok := cuc.mutation.UserRole(); ok {
		_spec.SetField(companyuser.FieldUserRole, field.TypeEnum, value)
		_node.UserRole = value
	}
	if nodes := cuc.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   companyuser.ProfileTable,
			Columns: []string{companyuser.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cuc.mutation.TripsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.TripsTable,
			Columns: []string{companyuser.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cuc.mutation.IncidentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.IncidentsTable,
			Columns: []string{companyuser.IncidentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(incident.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cuc.mutation.ParcelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.ParcelsTable,
			Columns: []string{companyuser.ParcelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cuc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   companyuser.NotificationsTable,
			Columns: companyuser.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cuc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companyuser.CompanyTable,
			Columns: []string{companyuser.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompanyUserCreateBulk is the builder for creating many CompanyUser entities in bulk.
type CompanyUserCreateBulk struct {
	config
	err      error
	builders []*CompanyUserCreate
}

// Save creates the CompanyUser entities in the database.
func (cucb *CompanyUserCreateBulk) Save(ctx context.Context) ([]*CompanyUser, error) {
	if cucb.err != nil {
		return nil, cucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cucb.builders))
	nodes := make([]*CompanyUser, len(cucb.builders))
	mutators := make([]Mutator, len(cucb.builders))
	for i := range cucb.builders {
		func(i int, root context.Context) {
			builder := cucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyUserMutation)
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
					_, err = mutators[i+1].Mutate(root, cucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cucb *CompanyUserCreateBulk) SaveX(ctx context.Context) []*CompanyUser {
	v, err := cucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cucb *CompanyUserCreateBulk) Exec(ctx context.Context) error {
	_, err := cucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cucb *CompanyUserCreateBulk) ExecX(ctx context.Context) {
	if err := cucb.Exec(ctx); err != nil {
		panic(err)
	}
}
