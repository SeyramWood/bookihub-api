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
	"github.com/SeyramWood/bookibus/ent/terminal"
	"github.com/SeyramWood/bookibus/ent/trip"
)

// TerminalUpdate is the builder for updating Terminal entities.
type TerminalUpdate struct {
	config
	hooks     []Hook
	mutation  *TerminalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TerminalUpdate builder.
func (tu *TerminalUpdate) Where(ps ...predicate.Terminal) *TerminalUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TerminalUpdate) SetUpdatedAt(t time.Time) *TerminalUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetAddress sets the "address" field.
func (tu *TerminalUpdate) SetAddress(s string) *TerminalUpdate {
	tu.mutation.SetAddress(s)
	return tu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (tu *TerminalUpdate) SetNillableAddress(s *string) *TerminalUpdate {
	if s != nil {
		tu.SetAddress(*s)
	}
	return tu
}

// ClearAddress clears the value of the "address" field.
func (tu *TerminalUpdate) ClearAddress() *TerminalUpdate {
	tu.mutation.ClearAddress()
	return tu
}

// SetLatitude sets the "latitude" field.
func (tu *TerminalUpdate) SetLatitude(f float64) *TerminalUpdate {
	tu.mutation.ResetLatitude()
	tu.mutation.SetLatitude(f)
	return tu
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (tu *TerminalUpdate) SetNillableLatitude(f *float64) *TerminalUpdate {
	if f != nil {
		tu.SetLatitude(*f)
	}
	return tu
}

// AddLatitude adds f to the "latitude" field.
func (tu *TerminalUpdate) AddLatitude(f float64) *TerminalUpdate {
	tu.mutation.AddLatitude(f)
	return tu
}

// ClearLatitude clears the value of the "latitude" field.
func (tu *TerminalUpdate) ClearLatitude() *TerminalUpdate {
	tu.mutation.ClearLatitude()
	return tu
}

// SetLongitude sets the "longitude" field.
func (tu *TerminalUpdate) SetLongitude(f float64) *TerminalUpdate {
	tu.mutation.ResetLongitude()
	tu.mutation.SetLongitude(f)
	return tu
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (tu *TerminalUpdate) SetNillableLongitude(f *float64) *TerminalUpdate {
	if f != nil {
		tu.SetLongitude(*f)
	}
	return tu
}

// AddLongitude adds f to the "longitude" field.
func (tu *TerminalUpdate) AddLongitude(f float64) *TerminalUpdate {
	tu.mutation.AddLongitude(f)
	return tu
}

// ClearLongitude clears the value of the "longitude" field.
func (tu *TerminalUpdate) ClearLongitude() *TerminalUpdate {
	tu.mutation.ClearLongitude()
	return tu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (tu *TerminalUpdate) SetCompanyID(id int) *TerminalUpdate {
	tu.mutation.SetCompanyID(id)
	return tu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (tu *TerminalUpdate) SetNillableCompanyID(id *int) *TerminalUpdate {
	if id != nil {
		tu = tu.SetCompanyID(*id)
	}
	return tu
}

// SetCompany sets the "company" edge to the Company entity.
func (tu *TerminalUpdate) SetCompany(c *Company) *TerminalUpdate {
	return tu.SetCompanyID(c.ID)
}

// AddFromIDs adds the "from" edge to the Trip entity by IDs.
func (tu *TerminalUpdate) AddFromIDs(ids ...int) *TerminalUpdate {
	tu.mutation.AddFromIDs(ids...)
	return tu
}

// AddFrom adds the "from" edges to the Trip entity.
func (tu *TerminalUpdate) AddFrom(t ...*Trip) *TerminalUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddFromIDs(ids...)
}

// AddToIDs adds the "to" edge to the Trip entity by IDs.
func (tu *TerminalUpdate) AddToIDs(ids ...int) *TerminalUpdate {
	tu.mutation.AddToIDs(ids...)
	return tu
}

// AddTo adds the "to" edges to the Trip entity.
func (tu *TerminalUpdate) AddTo(t ...*Trip) *TerminalUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddToIDs(ids...)
}

// Mutation returns the TerminalMutation object of the builder.
func (tu *TerminalUpdate) Mutation() *TerminalMutation {
	return tu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (tu *TerminalUpdate) ClearCompany() *TerminalUpdate {
	tu.mutation.ClearCompany()
	return tu
}

// ClearFrom clears all "from" edges to the Trip entity.
func (tu *TerminalUpdate) ClearFrom() *TerminalUpdate {
	tu.mutation.ClearFrom()
	return tu
}

// RemoveFromIDs removes the "from" edge to Trip entities by IDs.
func (tu *TerminalUpdate) RemoveFromIDs(ids ...int) *TerminalUpdate {
	tu.mutation.RemoveFromIDs(ids...)
	return tu
}

// RemoveFrom removes "from" edges to Trip entities.
func (tu *TerminalUpdate) RemoveFrom(t ...*Trip) *TerminalUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveFromIDs(ids...)
}

// ClearTo clears all "to" edges to the Trip entity.
func (tu *TerminalUpdate) ClearTo() *TerminalUpdate {
	tu.mutation.ClearTo()
	return tu
}

// RemoveToIDs removes the "to" edge to Trip entities by IDs.
func (tu *TerminalUpdate) RemoveToIDs(ids ...int) *TerminalUpdate {
	tu.mutation.RemoveToIDs(ids...)
	return tu
}

// RemoveTo removes "to" edges to Trip entities.
func (tu *TerminalUpdate) RemoveTo(t ...*Trip) *TerminalUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveToIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TerminalUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TerminalUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TerminalUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TerminalUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TerminalUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := terminal.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TerminalUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TerminalUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TerminalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(terminal.Table, terminal.Columns, sqlgraph.NewFieldSpec(terminal.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(terminal.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Address(); ok {
		_spec.SetField(terminal.FieldAddress, field.TypeString, value)
	}
	if tu.mutation.AddressCleared() {
		_spec.ClearField(terminal.FieldAddress, field.TypeString)
	}
	if value, ok := tu.mutation.Latitude(); ok {
		_spec.SetField(terminal.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedLatitude(); ok {
		_spec.AddField(terminal.FieldLatitude, field.TypeFloat64, value)
	}
	if tu.mutation.LatitudeCleared() {
		_spec.ClearField(terminal.FieldLatitude, field.TypeFloat64)
	}
	if value, ok := tu.mutation.Longitude(); ok {
		_spec.SetField(terminal.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedLongitude(); ok {
		_spec.AddField(terminal.FieldLongitude, field.TypeFloat64, value)
	}
	if tu.mutation.LongitudeCleared() {
		_spec.ClearField(terminal.FieldLongitude, field.TypeFloat64)
	}
	if tu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   terminal.CompanyTable,
			Columns: []string{terminal.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   terminal.CompanyTable,
			Columns: []string{terminal.CompanyColumn},
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
	if tu.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.FromTable,
			Columns: []string{terminal.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedFromIDs(); len(nodes) > 0 && !tu.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.FromTable,
			Columns: []string{terminal.FromColumn},
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
	if nodes := tu.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.FromTable,
			Columns: []string{terminal.FromColumn},
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
	if tu.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.ToTable,
			Columns: []string{terminal.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedToIDs(); len(nodes) > 0 && !tu.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.ToTable,
			Columns: []string{terminal.ToColumn},
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
	if nodes := tu.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.ToTable,
			Columns: []string{terminal.ToColumn},
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
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{terminal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TerminalUpdateOne is the builder for updating a single Terminal entity.
type TerminalUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TerminalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TerminalUpdateOne) SetUpdatedAt(t time.Time) *TerminalUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetAddress sets the "address" field.
func (tuo *TerminalUpdateOne) SetAddress(s string) *TerminalUpdateOne {
	tuo.mutation.SetAddress(s)
	return tuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (tuo *TerminalUpdateOne) SetNillableAddress(s *string) *TerminalUpdateOne {
	if s != nil {
		tuo.SetAddress(*s)
	}
	return tuo
}

// ClearAddress clears the value of the "address" field.
func (tuo *TerminalUpdateOne) ClearAddress() *TerminalUpdateOne {
	tuo.mutation.ClearAddress()
	return tuo
}

// SetLatitude sets the "latitude" field.
func (tuo *TerminalUpdateOne) SetLatitude(f float64) *TerminalUpdateOne {
	tuo.mutation.ResetLatitude()
	tuo.mutation.SetLatitude(f)
	return tuo
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (tuo *TerminalUpdateOne) SetNillableLatitude(f *float64) *TerminalUpdateOne {
	if f != nil {
		tuo.SetLatitude(*f)
	}
	return tuo
}

// AddLatitude adds f to the "latitude" field.
func (tuo *TerminalUpdateOne) AddLatitude(f float64) *TerminalUpdateOne {
	tuo.mutation.AddLatitude(f)
	return tuo
}

// ClearLatitude clears the value of the "latitude" field.
func (tuo *TerminalUpdateOne) ClearLatitude() *TerminalUpdateOne {
	tuo.mutation.ClearLatitude()
	return tuo
}

// SetLongitude sets the "longitude" field.
func (tuo *TerminalUpdateOne) SetLongitude(f float64) *TerminalUpdateOne {
	tuo.mutation.ResetLongitude()
	tuo.mutation.SetLongitude(f)
	return tuo
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (tuo *TerminalUpdateOne) SetNillableLongitude(f *float64) *TerminalUpdateOne {
	if f != nil {
		tuo.SetLongitude(*f)
	}
	return tuo
}

// AddLongitude adds f to the "longitude" field.
func (tuo *TerminalUpdateOne) AddLongitude(f float64) *TerminalUpdateOne {
	tuo.mutation.AddLongitude(f)
	return tuo
}

// ClearLongitude clears the value of the "longitude" field.
func (tuo *TerminalUpdateOne) ClearLongitude() *TerminalUpdateOne {
	tuo.mutation.ClearLongitude()
	return tuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (tuo *TerminalUpdateOne) SetCompanyID(id int) *TerminalUpdateOne {
	tuo.mutation.SetCompanyID(id)
	return tuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (tuo *TerminalUpdateOne) SetNillableCompanyID(id *int) *TerminalUpdateOne {
	if id != nil {
		tuo = tuo.SetCompanyID(*id)
	}
	return tuo
}

// SetCompany sets the "company" edge to the Company entity.
func (tuo *TerminalUpdateOne) SetCompany(c *Company) *TerminalUpdateOne {
	return tuo.SetCompanyID(c.ID)
}

// AddFromIDs adds the "from" edge to the Trip entity by IDs.
func (tuo *TerminalUpdateOne) AddFromIDs(ids ...int) *TerminalUpdateOne {
	tuo.mutation.AddFromIDs(ids...)
	return tuo
}

// AddFrom adds the "from" edges to the Trip entity.
func (tuo *TerminalUpdateOne) AddFrom(t ...*Trip) *TerminalUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddFromIDs(ids...)
}

// AddToIDs adds the "to" edge to the Trip entity by IDs.
func (tuo *TerminalUpdateOne) AddToIDs(ids ...int) *TerminalUpdateOne {
	tuo.mutation.AddToIDs(ids...)
	return tuo
}

// AddTo adds the "to" edges to the Trip entity.
func (tuo *TerminalUpdateOne) AddTo(t ...*Trip) *TerminalUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddToIDs(ids...)
}

// Mutation returns the TerminalMutation object of the builder.
func (tuo *TerminalUpdateOne) Mutation() *TerminalMutation {
	return tuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (tuo *TerminalUpdateOne) ClearCompany() *TerminalUpdateOne {
	tuo.mutation.ClearCompany()
	return tuo
}

// ClearFrom clears all "from" edges to the Trip entity.
func (tuo *TerminalUpdateOne) ClearFrom() *TerminalUpdateOne {
	tuo.mutation.ClearFrom()
	return tuo
}

// RemoveFromIDs removes the "from" edge to Trip entities by IDs.
func (tuo *TerminalUpdateOne) RemoveFromIDs(ids ...int) *TerminalUpdateOne {
	tuo.mutation.RemoveFromIDs(ids...)
	return tuo
}

// RemoveFrom removes "from" edges to Trip entities.
func (tuo *TerminalUpdateOne) RemoveFrom(t ...*Trip) *TerminalUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveFromIDs(ids...)
}

// ClearTo clears all "to" edges to the Trip entity.
func (tuo *TerminalUpdateOne) ClearTo() *TerminalUpdateOne {
	tuo.mutation.ClearTo()
	return tuo
}

// RemoveToIDs removes the "to" edge to Trip entities by IDs.
func (tuo *TerminalUpdateOne) RemoveToIDs(ids ...int) *TerminalUpdateOne {
	tuo.mutation.RemoveToIDs(ids...)
	return tuo
}

// RemoveTo removes "to" edges to Trip entities.
func (tuo *TerminalUpdateOne) RemoveTo(t ...*Trip) *TerminalUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveToIDs(ids...)
}

// Where appends a list predicates to the TerminalUpdate builder.
func (tuo *TerminalUpdateOne) Where(ps ...predicate.Terminal) *TerminalUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TerminalUpdateOne) Select(field string, fields ...string) *TerminalUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Terminal entity.
func (tuo *TerminalUpdateOne) Save(ctx context.Context) (*Terminal, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TerminalUpdateOne) SaveX(ctx context.Context) *Terminal {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TerminalUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TerminalUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TerminalUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := terminal.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TerminalUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TerminalUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TerminalUpdateOne) sqlSave(ctx context.Context) (_node *Terminal, err error) {
	_spec := sqlgraph.NewUpdateSpec(terminal.Table, terminal.Columns, sqlgraph.NewFieldSpec(terminal.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Terminal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, terminal.FieldID)
		for _, f := range fields {
			if !terminal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != terminal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(terminal.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Address(); ok {
		_spec.SetField(terminal.FieldAddress, field.TypeString, value)
	}
	if tuo.mutation.AddressCleared() {
		_spec.ClearField(terminal.FieldAddress, field.TypeString)
	}
	if value, ok := tuo.mutation.Latitude(); ok {
		_spec.SetField(terminal.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedLatitude(); ok {
		_spec.AddField(terminal.FieldLatitude, field.TypeFloat64, value)
	}
	if tuo.mutation.LatitudeCleared() {
		_spec.ClearField(terminal.FieldLatitude, field.TypeFloat64)
	}
	if value, ok := tuo.mutation.Longitude(); ok {
		_spec.SetField(terminal.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedLongitude(); ok {
		_spec.AddField(terminal.FieldLongitude, field.TypeFloat64, value)
	}
	if tuo.mutation.LongitudeCleared() {
		_spec.ClearField(terminal.FieldLongitude, field.TypeFloat64)
	}
	if tuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   terminal.CompanyTable,
			Columns: []string{terminal.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   terminal.CompanyTable,
			Columns: []string{terminal.CompanyColumn},
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
	if tuo.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.FromTable,
			Columns: []string{terminal.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedFromIDs(); len(nodes) > 0 && !tuo.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.FromTable,
			Columns: []string{terminal.FromColumn},
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
	if nodes := tuo.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.FromTable,
			Columns: []string{terminal.FromColumn},
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
	if tuo.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.ToTable,
			Columns: []string{terminal.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedToIDs(); len(nodes) > 0 && !tuo.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.ToTable,
			Columns: []string{terminal.ToColumn},
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
	if nodes := tuo.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   terminal.ToTable,
			Columns: []string{terminal.ToColumn},
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
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Terminal{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{terminal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
