// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/bookibususer"
	"github.com/SeyramWood/ent/notification"
	"github.com/SeyramWood/ent/user"
)

// BookibusUserCreate is the builder for creating a BookibusUser entity.
type BookibusUserCreate struct {
	config
	mutation *BookibusUserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (buc *BookibusUserCreate) SetCreatedAt(t time.Time) *BookibusUserCreate {
	buc.mutation.SetCreatedAt(t)
	return buc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buc *BookibusUserCreate) SetNillableCreatedAt(t *time.Time) *BookibusUserCreate {
	if t != nil {
		buc.SetCreatedAt(*t)
	}
	return buc
}

// SetUpdatedAt sets the "updated_at" field.
func (buc *BookibusUserCreate) SetUpdatedAt(t time.Time) *BookibusUserCreate {
	buc.mutation.SetUpdatedAt(t)
	return buc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buc *BookibusUserCreate) SetNillableUpdatedAt(t *time.Time) *BookibusUserCreate {
	if t != nil {
		buc.SetUpdatedAt(*t)
	}
	return buc
}

// SetLastName sets the "last_name" field.
func (buc *BookibusUserCreate) SetLastName(s string) *BookibusUserCreate {
	buc.mutation.SetLastName(s)
	return buc
}

// SetOtherName sets the "other_name" field.
func (buc *BookibusUserCreate) SetOtherName(s string) *BookibusUserCreate {
	buc.mutation.SetOtherName(s)
	return buc
}

// SetPhone sets the "phone" field.
func (buc *BookibusUserCreate) SetPhone(s string) *BookibusUserCreate {
	buc.mutation.SetPhone(s)
	return buc
}

// SetOtherPhone sets the "other_phone" field.
func (buc *BookibusUserCreate) SetOtherPhone(s string) *BookibusUserCreate {
	buc.mutation.SetOtherPhone(s)
	return buc
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (buc *BookibusUserCreate) SetNillableOtherPhone(s *string) *BookibusUserCreate {
	if s != nil {
		buc.SetOtherPhone(*s)
	}
	return buc
}

// SetRole sets the "role" field.
func (buc *BookibusUserCreate) SetRole(b bookibususer.Role) *BookibusUserCreate {
	buc.mutation.SetRole(b)
	return buc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (buc *BookibusUserCreate) SetNillableRole(b *bookibususer.Role) *BookibusUserCreate {
	if b != nil {
		buc.SetRole(*b)
	}
	return buc
}

// SetProfileID sets the "profile" edge to the User entity by ID.
func (buc *BookibusUserCreate) SetProfileID(id int) *BookibusUserCreate {
	buc.mutation.SetProfileID(id)
	return buc
}

// SetNillableProfileID sets the "profile" edge to the User entity by ID if the given value is not nil.
func (buc *BookibusUserCreate) SetNillableProfileID(id *int) *BookibusUserCreate {
	if id != nil {
		buc = buc.SetProfileID(*id)
	}
	return buc
}

// SetProfile sets the "profile" edge to the User entity.
func (buc *BookibusUserCreate) SetProfile(u *User) *BookibusUserCreate {
	return buc.SetProfileID(u.ID)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (buc *BookibusUserCreate) AddNotificationIDs(ids ...int) *BookibusUserCreate {
	buc.mutation.AddNotificationIDs(ids...)
	return buc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (buc *BookibusUserCreate) AddNotifications(n ...*Notification) *BookibusUserCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return buc.AddNotificationIDs(ids...)
}

// Mutation returns the BookibusUserMutation object of the builder.
func (buc *BookibusUserCreate) Mutation() *BookibusUserMutation {
	return buc.mutation
}

// Save creates the BookibusUser in the database.
func (buc *BookibusUserCreate) Save(ctx context.Context) (*BookibusUser, error) {
	buc.defaults()
	return withHooks(ctx, buc.sqlSave, buc.mutation, buc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (buc *BookibusUserCreate) SaveX(ctx context.Context) *BookibusUser {
	v, err := buc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (buc *BookibusUserCreate) Exec(ctx context.Context) error {
	_, err := buc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buc *BookibusUserCreate) ExecX(ctx context.Context) {
	if err := buc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buc *BookibusUserCreate) defaults() {
	if _, ok := buc.mutation.CreatedAt(); !ok {
		v := bookibususer.DefaultCreatedAt()
		buc.mutation.SetCreatedAt(v)
	}
	if _, ok := buc.mutation.UpdatedAt(); !ok {
		v := bookibususer.DefaultUpdatedAt()
		buc.mutation.SetUpdatedAt(v)
	}
	if _, ok := buc.mutation.Role(); !ok {
		v := bookibususer.DefaultRole
		buc.mutation.SetRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buc *BookibusUserCreate) check() error {
	if _, ok := buc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BookibusUser.created_at"`)}
	}
	if _, ok := buc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BookibusUser.updated_at"`)}
	}
	if _, ok := buc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "BookibusUser.last_name"`)}
	}
	if v, ok := buc.mutation.LastName(); ok {
		if err := bookibususer.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "BookibusUser.last_name": %w`, err)}
		}
	}
	if _, ok := buc.mutation.OtherName(); !ok {
		return &ValidationError{Name: "other_name", err: errors.New(`ent: missing required field "BookibusUser.other_name"`)}
	}
	if v, ok := buc.mutation.OtherName(); ok {
		if err := bookibususer.OtherNameValidator(v); err != nil {
			return &ValidationError{Name: "other_name", err: fmt.Errorf(`ent: validator failed for field "BookibusUser.other_name": %w`, err)}
		}
	}
	if _, ok := buc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "BookibusUser.phone"`)}
	}
	if v, ok := buc.mutation.Phone(); ok {
		if err := bookibususer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "BookibusUser.phone": %w`, err)}
		}
	}
	if _, ok := buc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "BookibusUser.role"`)}
	}
	if v, ok := buc.mutation.Role(); ok {
		if err := bookibususer.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "BookibusUser.role": %w`, err)}
		}
	}
	return nil
}

func (buc *BookibusUserCreate) sqlSave(ctx context.Context) (*BookibusUser, error) {
	if err := buc.check(); err != nil {
		return nil, err
	}
	_node, _spec := buc.createSpec()
	if err := sqlgraph.CreateNode(ctx, buc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	buc.mutation.id = &_node.ID
	buc.mutation.done = true
	return _node, nil
}

func (buc *BookibusUserCreate) createSpec() (*BookibusUser, *sqlgraph.CreateSpec) {
	var (
		_node = &BookibusUser{config: buc.config}
		_spec = sqlgraph.NewCreateSpec(bookibususer.Table, sqlgraph.NewFieldSpec(bookibususer.FieldID, field.TypeInt))
	)
	if value, ok := buc.mutation.CreatedAt(); ok {
		_spec.SetField(bookibususer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := buc.mutation.UpdatedAt(); ok {
		_spec.SetField(bookibususer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := buc.mutation.LastName(); ok {
		_spec.SetField(bookibususer.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := buc.mutation.OtherName(); ok {
		_spec.SetField(bookibususer.FieldOtherName, field.TypeString, value)
		_node.OtherName = value
	}
	if value, ok := buc.mutation.Phone(); ok {
		_spec.SetField(bookibususer.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := buc.mutation.OtherPhone(); ok {
		_spec.SetField(bookibususer.FieldOtherPhone, field.TypeString, value)
		_node.OtherPhone = value
	}
	if value, ok := buc.mutation.Role(); ok {
		_spec.SetField(bookibususer.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := buc.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   bookibususer.ProfileTable,
			Columns: []string{bookibususer.ProfileColumn},
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
	if nodes := buc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookibususer.NotificationsTable,
			Columns: bookibususer.NotificationsPrimaryKey,
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
	return _node, _spec
}

// BookibusUserCreateBulk is the builder for creating many BookibusUser entities in bulk.
type BookibusUserCreateBulk struct {
	config
	err      error
	builders []*BookibusUserCreate
}

// Save creates the BookibusUser entities in the database.
func (bucb *BookibusUserCreateBulk) Save(ctx context.Context) ([]*BookibusUser, error) {
	if bucb.err != nil {
		return nil, bucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bucb.builders))
	nodes := make([]*BookibusUser, len(bucb.builders))
	mutators := make([]Mutator, len(bucb.builders))
	for i := range bucb.builders {
		func(i int, root context.Context) {
			builder := bucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookibusUserMutation)
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
					_, err = mutators[i+1].Mutate(root, bucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bucb *BookibusUserCreateBulk) SaveX(ctx context.Context) []*BookibusUser {
	v, err := bucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bucb *BookibusUserCreateBulk) Exec(ctx context.Context) error {
	_, err := bucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bucb *BookibusUserCreateBulk) ExecX(ctx context.Context) {
	if err := bucb.Exec(ctx); err != nil {
		panic(err)
	}
}
