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
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/parcelimage"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// ParcelImageUpdate is the builder for updating ParcelImage entities.
type ParcelImageUpdate struct {
	config
	hooks     []Hook
	mutation  *ParcelImageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ParcelImageUpdate builder.
func (piu *ParcelImageUpdate) Where(ps ...predicate.ParcelImage) *ParcelImageUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetUpdatedAt sets the "updated_at" field.
func (piu *ParcelImageUpdate) SetUpdatedAt(t time.Time) *ParcelImageUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetImage sets the "image" field.
func (piu *ParcelImageUpdate) SetImage(s string) *ParcelImageUpdate {
	piu.mutation.SetImage(s)
	return piu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (piu *ParcelImageUpdate) SetNillableImage(s *string) *ParcelImageUpdate {
	if s != nil {
		piu.SetImage(*s)
	}
	return piu
}

// ClearImage clears the value of the "image" field.
func (piu *ParcelImageUpdate) ClearImage() *ParcelImageUpdate {
	piu.mutation.ClearImage()
	return piu
}

// SetKind sets the "kind" field.
func (piu *ParcelImageUpdate) SetKind(pa parcelimage.Kind) *ParcelImageUpdate {
	piu.mutation.SetKind(pa)
	return piu
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (piu *ParcelImageUpdate) SetNillableKind(pa *parcelimage.Kind) *ParcelImageUpdate {
	if pa != nil {
		piu.SetKind(*pa)
	}
	return piu
}

// SetParcelID sets the "parcel" edge to the Parcel entity by ID.
func (piu *ParcelImageUpdate) SetParcelID(id int) *ParcelImageUpdate {
	piu.mutation.SetParcelID(id)
	return piu
}

// SetNillableParcelID sets the "parcel" edge to the Parcel entity by ID if the given value is not nil.
func (piu *ParcelImageUpdate) SetNillableParcelID(id *int) *ParcelImageUpdate {
	if id != nil {
		piu = piu.SetParcelID(*id)
	}
	return piu
}

// SetParcel sets the "parcel" edge to the Parcel entity.
func (piu *ParcelImageUpdate) SetParcel(p *Parcel) *ParcelImageUpdate {
	return piu.SetParcelID(p.ID)
}

// Mutation returns the ParcelImageMutation object of the builder.
func (piu *ParcelImageUpdate) Mutation() *ParcelImageMutation {
	return piu.mutation
}

// ClearParcel clears the "parcel" edge to the Parcel entity.
func (piu *ParcelImageUpdate) ClearParcel() *ParcelImageUpdate {
	piu.mutation.ClearParcel()
	return piu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *ParcelImageUpdate) Save(ctx context.Context) (int, error) {
	piu.defaults()
	return withHooks(ctx, piu.sqlSave, piu.mutation, piu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piu *ParcelImageUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *ParcelImageUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *ParcelImageUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piu *ParcelImageUpdate) defaults() {
	if _, ok := piu.mutation.UpdatedAt(); !ok {
		v := parcelimage.UpdateDefaultUpdatedAt()
		piu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piu *ParcelImageUpdate) check() error {
	if v, ok := piu.mutation.Kind(); ok {
		if err := parcelimage.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "ParcelImage.kind": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (piu *ParcelImageUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ParcelImageUpdate {
	piu.modifiers = append(piu.modifiers, modifiers...)
	return piu
}

func (piu *ParcelImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := piu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(parcelimage.Table, parcelimage.Columns, sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt))
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.SetField(parcelimage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := piu.mutation.Image(); ok {
		_spec.SetField(parcelimage.FieldImage, field.TypeString, value)
	}
	if piu.mutation.ImageCleared() {
		_spec.ClearField(parcelimage.FieldImage, field.TypeString)
	}
	if value, ok := piu.mutation.Kind(); ok {
		_spec.SetField(parcelimage.FieldKind, field.TypeEnum, value)
	}
	if piu.mutation.ParcelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcelimage.ParcelTable,
			Columns: []string{parcelimage.ParcelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.ParcelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcelimage.ParcelTable,
			Columns: []string{parcelimage.ParcelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(piu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{parcelimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piu.mutation.done = true
	return n, nil
}

// ParcelImageUpdateOne is the builder for updating a single ParcelImage entity.
type ParcelImageUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ParcelImageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (piuo *ParcelImageUpdateOne) SetUpdatedAt(t time.Time) *ParcelImageUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetImage sets the "image" field.
func (piuo *ParcelImageUpdateOne) SetImage(s string) *ParcelImageUpdateOne {
	piuo.mutation.SetImage(s)
	return piuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (piuo *ParcelImageUpdateOne) SetNillableImage(s *string) *ParcelImageUpdateOne {
	if s != nil {
		piuo.SetImage(*s)
	}
	return piuo
}

// ClearImage clears the value of the "image" field.
func (piuo *ParcelImageUpdateOne) ClearImage() *ParcelImageUpdateOne {
	piuo.mutation.ClearImage()
	return piuo
}

// SetKind sets the "kind" field.
func (piuo *ParcelImageUpdateOne) SetKind(pa parcelimage.Kind) *ParcelImageUpdateOne {
	piuo.mutation.SetKind(pa)
	return piuo
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (piuo *ParcelImageUpdateOne) SetNillableKind(pa *parcelimage.Kind) *ParcelImageUpdateOne {
	if pa != nil {
		piuo.SetKind(*pa)
	}
	return piuo
}

// SetParcelID sets the "parcel" edge to the Parcel entity by ID.
func (piuo *ParcelImageUpdateOne) SetParcelID(id int) *ParcelImageUpdateOne {
	piuo.mutation.SetParcelID(id)
	return piuo
}

// SetNillableParcelID sets the "parcel" edge to the Parcel entity by ID if the given value is not nil.
func (piuo *ParcelImageUpdateOne) SetNillableParcelID(id *int) *ParcelImageUpdateOne {
	if id != nil {
		piuo = piuo.SetParcelID(*id)
	}
	return piuo
}

// SetParcel sets the "parcel" edge to the Parcel entity.
func (piuo *ParcelImageUpdateOne) SetParcel(p *Parcel) *ParcelImageUpdateOne {
	return piuo.SetParcelID(p.ID)
}

// Mutation returns the ParcelImageMutation object of the builder.
func (piuo *ParcelImageUpdateOne) Mutation() *ParcelImageMutation {
	return piuo.mutation
}

// ClearParcel clears the "parcel" edge to the Parcel entity.
func (piuo *ParcelImageUpdateOne) ClearParcel() *ParcelImageUpdateOne {
	piuo.mutation.ClearParcel()
	return piuo
}

// Where appends a list predicates to the ParcelImageUpdate builder.
func (piuo *ParcelImageUpdateOne) Where(ps ...predicate.ParcelImage) *ParcelImageUpdateOne {
	piuo.mutation.Where(ps...)
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *ParcelImageUpdateOne) Select(field string, fields ...string) *ParcelImageUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated ParcelImage entity.
func (piuo *ParcelImageUpdateOne) Save(ctx context.Context) (*ParcelImage, error) {
	piuo.defaults()
	return withHooks(ctx, piuo.sqlSave, piuo.mutation, piuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *ParcelImageUpdateOne) SaveX(ctx context.Context) *ParcelImage {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *ParcelImageUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *ParcelImageUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piuo *ParcelImageUpdateOne) defaults() {
	if _, ok := piuo.mutation.UpdatedAt(); !ok {
		v := parcelimage.UpdateDefaultUpdatedAt()
		piuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piuo *ParcelImageUpdateOne) check() error {
	if v, ok := piuo.mutation.Kind(); ok {
		if err := parcelimage.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "ParcelImage.kind": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (piuo *ParcelImageUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ParcelImageUpdateOne {
	piuo.modifiers = append(piuo.modifiers, modifiers...)
	return piuo
}

func (piuo *ParcelImageUpdateOne) sqlSave(ctx context.Context) (_node *ParcelImage, err error) {
	if err := piuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(parcelimage.Table, parcelimage.Columns, sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt))
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ParcelImage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, parcelimage.FieldID)
		for _, f := range fields {
			if !parcelimage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != parcelimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.SetField(parcelimage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := piuo.mutation.Image(); ok {
		_spec.SetField(parcelimage.FieldImage, field.TypeString, value)
	}
	if piuo.mutation.ImageCleared() {
		_spec.ClearField(parcelimage.FieldImage, field.TypeString)
	}
	if value, ok := piuo.mutation.Kind(); ok {
		_spec.SetField(parcelimage.FieldKind, field.TypeEnum, value)
	}
	if piuo.mutation.ParcelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcelimage.ParcelTable,
			Columns: []string{parcelimage.ParcelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.ParcelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcelimage.ParcelTable,
			Columns: []string{parcelimage.ParcelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(piuo.modifiers...)
	_node = &ParcelImage{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{parcelimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piuo.mutation.done = true
	return _node, nil
}
