// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/incidentimage"
	"github.com/SeyramWood/ent/predicate"
)

// IncidentImageDelete is the builder for deleting a IncidentImage entity.
type IncidentImageDelete struct {
	config
	hooks    []Hook
	mutation *IncidentImageMutation
}

// Where appends a list predicates to the IncidentImageDelete builder.
func (iid *IncidentImageDelete) Where(ps ...predicate.IncidentImage) *IncidentImageDelete {
	iid.mutation.Where(ps...)
	return iid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (iid *IncidentImageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, iid.sqlExec, iid.mutation, iid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (iid *IncidentImageDelete) ExecX(ctx context.Context) int {
	n, err := iid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (iid *IncidentImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(incidentimage.Table, sqlgraph.NewFieldSpec(incidentimage.FieldID, field.TypeInt))
	if ps := iid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, iid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	iid.mutation.done = true
	return affected, err
}

// IncidentImageDeleteOne is the builder for deleting a single IncidentImage entity.
type IncidentImageDeleteOne struct {
	iid *IncidentImageDelete
}

// Where appends a list predicates to the IncidentImageDelete builder.
func (iido *IncidentImageDeleteOne) Where(ps ...predicate.IncidentImage) *IncidentImageDeleteOne {
	iido.iid.mutation.Where(ps...)
	return iido
}

// Exec executes the deletion query.
func (iido *IncidentImageDeleteOne) Exec(ctx context.Context) error {
	n, err := iido.iid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{incidentimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iido *IncidentImageDeleteOne) ExecX(ctx context.Context) {
	if err := iido.Exec(ctx); err != nil {
		panic(err)
	}
}
