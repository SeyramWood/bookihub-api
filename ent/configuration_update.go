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
	"github.com/SeyramWood/bookibus/ent/configuration"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/schema"
)

// ConfigurationUpdate is the builder for updating Configuration entities.
type ConfigurationUpdate struct {
	config
	hooks     []Hook
	mutation  *ConfigurationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ConfigurationUpdate builder.
func (cu *ConfigurationUpdate) Where(ps ...predicate.Configuration) *ConfigurationUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ConfigurationUpdate) SetUpdatedAt(t time.Time) *ConfigurationUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetCharge sets the "charge" field.
func (cu *ConfigurationUpdate) SetCharge(s *schema.Charge) *ConfigurationUpdate {
	cu.mutation.SetCharge(s)
	return cu
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cu *ConfigurationUpdate) Mutation() *ConfigurationMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConfigurationUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfigurationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfigurationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfigurationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConfigurationUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := configuration.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *ConfigurationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ConfigurationUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *ConfigurationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(configuration.Table, configuration.Columns, sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(configuration.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Charge(); ok {
		_spec.SetField(configuration.FieldCharge, field.TypeJSON, value)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConfigurationUpdateOne is the builder for updating a single Configuration entity.
type ConfigurationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ConfigurationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ConfigurationUpdateOne) SetUpdatedAt(t time.Time) *ConfigurationUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetCharge sets the "charge" field.
func (cuo *ConfigurationUpdateOne) SetCharge(s *schema.Charge) *ConfigurationUpdateOne {
	cuo.mutation.SetCharge(s)
	return cuo
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cuo *ConfigurationUpdateOne) Mutation() *ConfigurationMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ConfigurationUpdate builder.
func (cuo *ConfigurationUpdateOne) Where(ps ...predicate.Configuration) *ConfigurationUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConfigurationUpdateOne) Select(field string, fields ...string) *ConfigurationUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Configuration entity.
func (cuo *ConfigurationUpdateOne) Save(ctx context.Context) (*Configuration, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) SaveX(ctx context.Context) *Configuration {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConfigurationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConfigurationUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := configuration.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *ConfigurationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ConfigurationUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *ConfigurationUpdateOne) sqlSave(ctx context.Context) (_node *Configuration, err error) {
	_spec := sqlgraph.NewUpdateSpec(configuration.Table, configuration.Columns, sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Configuration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, configuration.FieldID)
		for _, f := range fields {
			if !configuration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != configuration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(configuration.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Charge(); ok {
		_spec.SetField(configuration.FieldCharge, field.TypeJSON, value)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Configuration{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}