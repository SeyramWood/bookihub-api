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
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/routestop"
	"github.com/SeyramWood/bookibus/ent/trip"
)

// RouteStopUpdate is the builder for updating RouteStop entities.
type RouteStopUpdate struct {
	config
	hooks     []Hook
	mutation  *RouteStopMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RouteStopUpdate builder.
func (rsu *RouteStopUpdate) Where(ps ...predicate.RouteStop) *RouteStopUpdate {
	rsu.mutation.Where(ps...)
	return rsu
}

// SetUpdatedAt sets the "updated_at" field.
func (rsu *RouteStopUpdate) SetUpdatedAt(t time.Time) *RouteStopUpdate {
	rsu.mutation.SetUpdatedAt(t)
	return rsu
}

// SetAddress sets the "address" field.
func (rsu *RouteStopUpdate) SetAddress(s string) *RouteStopUpdate {
	rsu.mutation.SetAddress(s)
	return rsu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (rsu *RouteStopUpdate) SetNillableAddress(s *string) *RouteStopUpdate {
	if s != nil {
		rsu.SetAddress(*s)
	}
	return rsu
}

// ClearAddress clears the value of the "address" field.
func (rsu *RouteStopUpdate) ClearAddress() *RouteStopUpdate {
	rsu.mutation.ClearAddress()
	return rsu
}

// SetLatitude sets the "latitude" field.
func (rsu *RouteStopUpdate) SetLatitude(f float64) *RouteStopUpdate {
	rsu.mutation.ResetLatitude()
	rsu.mutation.SetLatitude(f)
	return rsu
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (rsu *RouteStopUpdate) SetNillableLatitude(f *float64) *RouteStopUpdate {
	if f != nil {
		rsu.SetLatitude(*f)
	}
	return rsu
}

// AddLatitude adds f to the "latitude" field.
func (rsu *RouteStopUpdate) AddLatitude(f float64) *RouteStopUpdate {
	rsu.mutation.AddLatitude(f)
	return rsu
}

// ClearLatitude clears the value of the "latitude" field.
func (rsu *RouteStopUpdate) ClearLatitude() *RouteStopUpdate {
	rsu.mutation.ClearLatitude()
	return rsu
}

// SetLongitude sets the "longitude" field.
func (rsu *RouteStopUpdate) SetLongitude(f float64) *RouteStopUpdate {
	rsu.mutation.ResetLongitude()
	rsu.mutation.SetLongitude(f)
	return rsu
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (rsu *RouteStopUpdate) SetNillableLongitude(f *float64) *RouteStopUpdate {
	if f != nil {
		rsu.SetLongitude(*f)
	}
	return rsu
}

// AddLongitude adds f to the "longitude" field.
func (rsu *RouteStopUpdate) AddLongitude(f float64) *RouteStopUpdate {
	rsu.mutation.AddLongitude(f)
	return rsu
}

// ClearLongitude clears the value of the "longitude" field.
func (rsu *RouteStopUpdate) ClearLongitude() *RouteStopUpdate {
	rsu.mutation.ClearLongitude()
	return rsu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (rsu *RouteStopUpdate) SetCompanyID(id int) *RouteStopUpdate {
	rsu.mutation.SetCompanyID(id)
	return rsu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (rsu *RouteStopUpdate) SetNillableCompanyID(id *int) *RouteStopUpdate {
	if id != nil {
		rsu = rsu.SetCompanyID(*id)
	}
	return rsu
}

// SetCompany sets the "company" edge to the Company entity.
func (rsu *RouteStopUpdate) SetCompany(c *Company) *RouteStopUpdate {
	return rsu.SetCompanyID(c.ID)
}

// AddTripIDs adds the "trip" edge to the Trip entity by IDs.
func (rsu *RouteStopUpdate) AddTripIDs(ids ...int) *RouteStopUpdate {
	rsu.mutation.AddTripIDs(ids...)
	return rsu
}

// AddTrip adds the "trip" edges to the Trip entity.
func (rsu *RouteStopUpdate) AddTrip(t ...*Trip) *RouteStopUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rsu.AddTripIDs(ids...)
}

// Mutation returns the RouteStopMutation object of the builder.
func (rsu *RouteStopUpdate) Mutation() *RouteStopMutation {
	return rsu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (rsu *RouteStopUpdate) ClearCompany() *RouteStopUpdate {
	rsu.mutation.ClearCompany()
	return rsu
}

// ClearTrip clears all "trip" edges to the Trip entity.
func (rsu *RouteStopUpdate) ClearTrip() *RouteStopUpdate {
	rsu.mutation.ClearTrip()
	return rsu
}

// RemoveTripIDs removes the "trip" edge to Trip entities by IDs.
func (rsu *RouteStopUpdate) RemoveTripIDs(ids ...int) *RouteStopUpdate {
	rsu.mutation.RemoveTripIDs(ids...)
	return rsu
}

// RemoveTrip removes "trip" edges to Trip entities.
func (rsu *RouteStopUpdate) RemoveTrip(t ...*Trip) *RouteStopUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rsu.RemoveTripIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rsu *RouteStopUpdate) Save(ctx context.Context) (int, error) {
	rsu.defaults()
	return withHooks(ctx, rsu.sqlSave, rsu.mutation, rsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rsu *RouteStopUpdate) SaveX(ctx context.Context) int {
	affected, err := rsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rsu *RouteStopUpdate) Exec(ctx context.Context) error {
	_, err := rsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsu *RouteStopUpdate) ExecX(ctx context.Context) {
	if err := rsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rsu *RouteStopUpdate) defaults() {
	if _, ok := rsu.mutation.UpdatedAt(); !ok {
		v := routestop.UpdateDefaultUpdatedAt()
		rsu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rsu *RouteStopUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RouteStopUpdate {
	rsu.modifiers = append(rsu.modifiers, modifiers...)
	return rsu
}

func (rsu *RouteStopUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(routestop.Table, routestop.Columns, sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt))
	if ps := rsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsu.mutation.UpdatedAt(); ok {
		_spec.SetField(routestop.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rsu.mutation.Address(); ok {
		_spec.SetField(routestop.FieldAddress, field.TypeString, value)
	}
	if rsu.mutation.AddressCleared() {
		_spec.ClearField(routestop.FieldAddress, field.TypeString)
	}
	if value, ok := rsu.mutation.Latitude(); ok {
		_spec.SetField(routestop.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := rsu.mutation.AddedLatitude(); ok {
		_spec.AddField(routestop.FieldLatitude, field.TypeFloat64, value)
	}
	if rsu.mutation.LatitudeCleared() {
		_spec.ClearField(routestop.FieldLatitude, field.TypeFloat64)
	}
	if value, ok := rsu.mutation.Longitude(); ok {
		_spec.SetField(routestop.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := rsu.mutation.AddedLongitude(); ok {
		_spec.AddField(routestop.FieldLongitude, field.TypeFloat64, value)
	}
	if rsu.mutation.LongitudeCleared() {
		_spec.ClearField(routestop.FieldLongitude, field.TypeFloat64)
	}
	if rsu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routestop.CompanyTable,
			Columns: []string{routestop.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routestop.CompanyTable,
			Columns: []string{routestop.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rsu.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   routestop.TripTable,
			Columns: routestop.TripPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsu.mutation.RemovedTripIDs(); len(nodes) > 0 && !rsu.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   routestop.TripTable,
			Columns: routestop.TripPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsu.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   routestop.TripTable,
			Columns: routestop.TripPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(rsu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routestop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rsu.mutation.done = true
	return n, nil
}

// RouteStopUpdateOne is the builder for updating a single RouteStop entity.
type RouteStopUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RouteStopMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (rsuo *RouteStopUpdateOne) SetUpdatedAt(t time.Time) *RouteStopUpdateOne {
	rsuo.mutation.SetUpdatedAt(t)
	return rsuo
}

// SetAddress sets the "address" field.
func (rsuo *RouteStopUpdateOne) SetAddress(s string) *RouteStopUpdateOne {
	rsuo.mutation.SetAddress(s)
	return rsuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (rsuo *RouteStopUpdateOne) SetNillableAddress(s *string) *RouteStopUpdateOne {
	if s != nil {
		rsuo.SetAddress(*s)
	}
	return rsuo
}

// ClearAddress clears the value of the "address" field.
func (rsuo *RouteStopUpdateOne) ClearAddress() *RouteStopUpdateOne {
	rsuo.mutation.ClearAddress()
	return rsuo
}

// SetLatitude sets the "latitude" field.
func (rsuo *RouteStopUpdateOne) SetLatitude(f float64) *RouteStopUpdateOne {
	rsuo.mutation.ResetLatitude()
	rsuo.mutation.SetLatitude(f)
	return rsuo
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (rsuo *RouteStopUpdateOne) SetNillableLatitude(f *float64) *RouteStopUpdateOne {
	if f != nil {
		rsuo.SetLatitude(*f)
	}
	return rsuo
}

// AddLatitude adds f to the "latitude" field.
func (rsuo *RouteStopUpdateOne) AddLatitude(f float64) *RouteStopUpdateOne {
	rsuo.mutation.AddLatitude(f)
	return rsuo
}

// ClearLatitude clears the value of the "latitude" field.
func (rsuo *RouteStopUpdateOne) ClearLatitude() *RouteStopUpdateOne {
	rsuo.mutation.ClearLatitude()
	return rsuo
}

// SetLongitude sets the "longitude" field.
func (rsuo *RouteStopUpdateOne) SetLongitude(f float64) *RouteStopUpdateOne {
	rsuo.mutation.ResetLongitude()
	rsuo.mutation.SetLongitude(f)
	return rsuo
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (rsuo *RouteStopUpdateOne) SetNillableLongitude(f *float64) *RouteStopUpdateOne {
	if f != nil {
		rsuo.SetLongitude(*f)
	}
	return rsuo
}

// AddLongitude adds f to the "longitude" field.
func (rsuo *RouteStopUpdateOne) AddLongitude(f float64) *RouteStopUpdateOne {
	rsuo.mutation.AddLongitude(f)
	return rsuo
}

// ClearLongitude clears the value of the "longitude" field.
func (rsuo *RouteStopUpdateOne) ClearLongitude() *RouteStopUpdateOne {
	rsuo.mutation.ClearLongitude()
	return rsuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (rsuo *RouteStopUpdateOne) SetCompanyID(id int) *RouteStopUpdateOne {
	rsuo.mutation.SetCompanyID(id)
	return rsuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (rsuo *RouteStopUpdateOne) SetNillableCompanyID(id *int) *RouteStopUpdateOne {
	if id != nil {
		rsuo = rsuo.SetCompanyID(*id)
	}
	return rsuo
}

// SetCompany sets the "company" edge to the Company entity.
func (rsuo *RouteStopUpdateOne) SetCompany(c *Company) *RouteStopUpdateOne {
	return rsuo.SetCompanyID(c.ID)
}

// AddTripIDs adds the "trip" edge to the Trip entity by IDs.
func (rsuo *RouteStopUpdateOne) AddTripIDs(ids ...int) *RouteStopUpdateOne {
	rsuo.mutation.AddTripIDs(ids...)
	return rsuo
}

// AddTrip adds the "trip" edges to the Trip entity.
func (rsuo *RouteStopUpdateOne) AddTrip(t ...*Trip) *RouteStopUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rsuo.AddTripIDs(ids...)
}

// Mutation returns the RouteStopMutation object of the builder.
func (rsuo *RouteStopUpdateOne) Mutation() *RouteStopMutation {
	return rsuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (rsuo *RouteStopUpdateOne) ClearCompany() *RouteStopUpdateOne {
	rsuo.mutation.ClearCompany()
	return rsuo
}

// ClearTrip clears all "trip" edges to the Trip entity.
func (rsuo *RouteStopUpdateOne) ClearTrip() *RouteStopUpdateOne {
	rsuo.mutation.ClearTrip()
	return rsuo
}

// RemoveTripIDs removes the "trip" edge to Trip entities by IDs.
func (rsuo *RouteStopUpdateOne) RemoveTripIDs(ids ...int) *RouteStopUpdateOne {
	rsuo.mutation.RemoveTripIDs(ids...)
	return rsuo
}

// RemoveTrip removes "trip" edges to Trip entities.
func (rsuo *RouteStopUpdateOne) RemoveTrip(t ...*Trip) *RouteStopUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rsuo.RemoveTripIDs(ids...)
}

// Where appends a list predicates to the RouteStopUpdate builder.
func (rsuo *RouteStopUpdateOne) Where(ps ...predicate.RouteStop) *RouteStopUpdateOne {
	rsuo.mutation.Where(ps...)
	return rsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rsuo *RouteStopUpdateOne) Select(field string, fields ...string) *RouteStopUpdateOne {
	rsuo.fields = append([]string{field}, fields...)
	return rsuo
}

// Save executes the query and returns the updated RouteStop entity.
func (rsuo *RouteStopUpdateOne) Save(ctx context.Context) (*RouteStop, error) {
	rsuo.defaults()
	return withHooks(ctx, rsuo.sqlSave, rsuo.mutation, rsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rsuo *RouteStopUpdateOne) SaveX(ctx context.Context) *RouteStop {
	node, err := rsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rsuo *RouteStopUpdateOne) Exec(ctx context.Context) error {
	_, err := rsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsuo *RouteStopUpdateOne) ExecX(ctx context.Context) {
	if err := rsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rsuo *RouteStopUpdateOne) defaults() {
	if _, ok := rsuo.mutation.UpdatedAt(); !ok {
		v := routestop.UpdateDefaultUpdatedAt()
		rsuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rsuo *RouteStopUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RouteStopUpdateOne {
	rsuo.modifiers = append(rsuo.modifiers, modifiers...)
	return rsuo
}

func (rsuo *RouteStopUpdateOne) sqlSave(ctx context.Context) (_node *RouteStop, err error) {
	_spec := sqlgraph.NewUpdateSpec(routestop.Table, routestop.Columns, sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt))
	id, ok := rsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RouteStop.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routestop.FieldID)
		for _, f := range fields {
			if !routestop.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != routestop.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(routestop.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rsuo.mutation.Address(); ok {
		_spec.SetField(routestop.FieldAddress, field.TypeString, value)
	}
	if rsuo.mutation.AddressCleared() {
		_spec.ClearField(routestop.FieldAddress, field.TypeString)
	}
	if value, ok := rsuo.mutation.Latitude(); ok {
		_spec.SetField(routestop.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := rsuo.mutation.AddedLatitude(); ok {
		_spec.AddField(routestop.FieldLatitude, field.TypeFloat64, value)
	}
	if rsuo.mutation.LatitudeCleared() {
		_spec.ClearField(routestop.FieldLatitude, field.TypeFloat64)
	}
	if value, ok := rsuo.mutation.Longitude(); ok {
		_spec.SetField(routestop.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := rsuo.mutation.AddedLongitude(); ok {
		_spec.AddField(routestop.FieldLongitude, field.TypeFloat64, value)
	}
	if rsuo.mutation.LongitudeCleared() {
		_spec.ClearField(routestop.FieldLongitude, field.TypeFloat64)
	}
	if rsuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routestop.CompanyTable,
			Columns: []string{routestop.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routestop.CompanyTable,
			Columns: []string{routestop.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rsuo.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   routestop.TripTable,
			Columns: routestop.TripPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsuo.mutation.RemovedTripIDs(); len(nodes) > 0 && !rsuo.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   routestop.TripTable,
			Columns: routestop.TripPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsuo.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   routestop.TripTable,
			Columns: routestop.TripPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(rsuo.modifiers...)
	_node = &RouteStop{config: rsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routestop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rsuo.mutation.done = true
	return _node, nil
}
