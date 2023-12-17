// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/incident"
	"github.com/SeyramWood/bookibus/ent/notification"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/route"
	"github.com/SeyramWood/bookibus/ent/schema"
	"github.com/SeyramWood/bookibus/ent/terminal"
	"github.com/SeyramWood/bookibus/ent/transaction"
	"github.com/SeyramWood/bookibus/ent/trip"
	"github.com/SeyramWood/bookibus/ent/vehicle"
)

// CompanyCreate is the builder for creating a Company entity.
type CompanyCreate struct {
	config
	mutation *CompanyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CompanyCreate) SetCreatedAt(t time.Time) *CompanyCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableCreatedAt(t *time.Time) *CompanyCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CompanyCreate) SetUpdatedAt(t time.Time) *CompanyCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableUpdatedAt(t *time.Time) *CompanyCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CompanyCreate) SetName(s string) *CompanyCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetPhone sets the "phone" field.
func (cc *CompanyCreate) SetPhone(s string) *CompanyCreate {
	cc.mutation.SetPhone(s)
	return cc
}

// SetEmail sets the "email" field.
func (cc *CompanyCreate) SetEmail(s string) *CompanyCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetCertificate sets the "certificate" field.
func (cc *CompanyCreate) SetCertificate(s string) *CompanyCreate {
	cc.mutation.SetCertificate(s)
	return cc
}

// SetNillableCertificate sets the "certificate" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableCertificate(s *string) *CompanyCreate {
	if s != nil {
		cc.SetCertificate(*s)
	}
	return cc
}

// SetBankAccount sets the "bank_account" field.
func (cc *CompanyCreate) SetBankAccount(sa *schema.BankAccount) *CompanyCreate {
	cc.mutation.SetBankAccount(sa)
	return cc
}

// SetContactPerson sets the "contact_person" field.
func (cc *CompanyCreate) SetContactPerson(sp *schema.ContactPerson) *CompanyCreate {
	cc.mutation.SetContactPerson(sp)
	return cc
}

// SetLogo sets the "logo" field.
func (cc *CompanyCreate) SetLogo(s string) *CompanyCreate {
	cc.mutation.SetLogo(s)
	return cc
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableLogo(s *string) *CompanyCreate {
	if s != nil {
		cc.SetLogo(*s)
	}
	return cc
}

// SetOnboardingStatus sets the "onboarding_status" field.
func (cc *CompanyCreate) SetOnboardingStatus(cs company.OnboardingStatus) *CompanyCreate {
	cc.mutation.SetOnboardingStatus(cs)
	return cc
}

// SetNillableOnboardingStatus sets the "onboarding_status" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableOnboardingStatus(cs *company.OnboardingStatus) *CompanyCreate {
	if cs != nil {
		cc.SetOnboardingStatus(*cs)
	}
	return cc
}

// SetOnboardingStage sets the "onboarding_stage" field.
func (cc *CompanyCreate) SetOnboardingStage(i int8) *CompanyCreate {
	cc.mutation.SetOnboardingStage(i)
	return cc
}

// SetNillableOnboardingStage sets the "onboarding_stage" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableOnboardingStage(i *int8) *CompanyCreate {
	if i != nil {
		cc.SetOnboardingStage(*i)
	}
	return cc
}

// AddProfileIDs adds the "profile" edge to the CompanyUser entity by IDs.
func (cc *CompanyCreate) AddProfileIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddProfileIDs(ids...)
	return cc
}

// AddProfile adds the "profile" edges to the CompanyUser entity.
func (cc *CompanyCreate) AddProfile(c ...*CompanyUser) *CompanyCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddProfileIDs(ids...)
}

// AddTerminalIDs adds the "terminals" edge to the Terminal entity by IDs.
func (cc *CompanyCreate) AddTerminalIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddTerminalIDs(ids...)
	return cc
}

// AddTerminals adds the "terminals" edges to the Terminal entity.
func (cc *CompanyCreate) AddTerminals(t ...*Terminal) *CompanyCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTerminalIDs(ids...)
}

// AddVehicleIDs adds the "vehicles" edge to the Vehicle entity by IDs.
func (cc *CompanyCreate) AddVehicleIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddVehicleIDs(ids...)
	return cc
}

// AddVehicles adds the "vehicles" edges to the Vehicle entity.
func (cc *CompanyCreate) AddVehicles(v ...*Vehicle) *CompanyCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cc.AddVehicleIDs(ids...)
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (cc *CompanyCreate) AddRouteIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddRouteIDs(ids...)
	return cc
}

// AddRoutes adds the "routes" edges to the Route entity.
func (cc *CompanyCreate) AddRoutes(r ...*Route) *CompanyCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRouteIDs(ids...)
}

// AddTripIDs adds the "trips" edge to the Trip entity by IDs.
func (cc *CompanyCreate) AddTripIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddTripIDs(ids...)
	return cc
}

// AddTrips adds the "trips" edges to the Trip entity.
func (cc *CompanyCreate) AddTrips(t ...*Trip) *CompanyCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTripIDs(ids...)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (cc *CompanyCreate) AddBookingIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddBookingIDs(ids...)
	return cc
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (cc *CompanyCreate) AddBookings(b ...*Booking) *CompanyCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cc.AddBookingIDs(ids...)
}

// AddIncidentIDs adds the "incidents" edge to the Incident entity by IDs.
func (cc *CompanyCreate) AddIncidentIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddIncidentIDs(ids...)
	return cc
}

// AddIncidents adds the "incidents" edges to the Incident entity.
func (cc *CompanyCreate) AddIncidents(i ...*Incident) *CompanyCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cc.AddIncidentIDs(ids...)
}

// AddParcelIDs adds the "parcels" edge to the Parcel entity by IDs.
func (cc *CompanyCreate) AddParcelIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddParcelIDs(ids...)
	return cc
}

// AddParcels adds the "parcels" edges to the Parcel entity.
func (cc *CompanyCreate) AddParcels(p ...*Parcel) *CompanyCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddParcelIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (cc *CompanyCreate) AddTransactionIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddTransactionIDs(ids...)
	return cc
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (cc *CompanyCreate) AddTransactions(t ...*Transaction) *CompanyCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTransactionIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (cc *CompanyCreate) AddNotificationIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddNotificationIDs(ids...)
	return cc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (cc *CompanyCreate) AddNotifications(n ...*Notification) *CompanyCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cc.AddNotificationIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cc *CompanyCreate) Mutation() *CompanyMutation {
	return cc.mutation
}

// Save creates the Company in the database.
func (cc *CompanyCreate) Save(ctx context.Context) (*Company, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompanyCreate) SaveX(ctx context.Context) *Company {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompanyCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompanyCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CompanyCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := company.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := company.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.OnboardingStatus(); !ok {
		v := company.DefaultOnboardingStatus
		cc.mutation.SetOnboardingStatus(v)
	}
	if _, ok := cc.mutation.OnboardingStage(); !ok {
		v := company.DefaultOnboardingStage
		cc.mutation.SetOnboardingStage(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompanyCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Company.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Company.updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Company.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := company.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Company.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Company.phone"`)}
	}
	if v, ok := cc.mutation.Phone(); ok {
		if err := company.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Company.phone": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Company.email"`)}
	}
	if v, ok := cc.mutation.Email(); ok {
		if err := company.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Company.email": %w`, err)}
		}
	}
	if _, ok := cc.mutation.OnboardingStatus(); !ok {
		return &ValidationError{Name: "onboarding_status", err: errors.New(`ent: missing required field "Company.onboarding_status"`)}
	}
	if v, ok := cc.mutation.OnboardingStatus(); ok {
		if err := company.OnboardingStatusValidator(v); err != nil {
			return &ValidationError{Name: "onboarding_status", err: fmt.Errorf(`ent: validator failed for field "Company.onboarding_status": %w`, err)}
		}
	}
	if _, ok := cc.mutation.OnboardingStage(); !ok {
		return &ValidationError{Name: "onboarding_stage", err: errors.New(`ent: missing required field "Company.onboarding_stage"`)}
	}
	return nil
}

func (cc *CompanyCreate) sqlSave(ctx context.Context) (*Company, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CompanyCreate) createSpec() (*Company, *sqlgraph.CreateSpec) {
	var (
		_node = &Company{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(company.Table, sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(company.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(company.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(company.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Phone(); ok {
		_spec.SetField(company.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.SetField(company.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cc.mutation.Certificate(); ok {
		_spec.SetField(company.FieldCertificate, field.TypeString, value)
		_node.Certificate = value
	}
	if value, ok := cc.mutation.BankAccount(); ok {
		_spec.SetField(company.FieldBankAccount, field.TypeJSON, value)
		_node.BankAccount = value
	}
	if value, ok := cc.mutation.ContactPerson(); ok {
		_spec.SetField(company.FieldContactPerson, field.TypeJSON, value)
		_node.ContactPerson = value
	}
	if value, ok := cc.mutation.Logo(); ok {
		_spec.SetField(company.FieldLogo, field.TypeString, value)
		_node.Logo = value
	}
	if value, ok := cc.mutation.OnboardingStatus(); ok {
		_spec.SetField(company.FieldOnboardingStatus, field.TypeEnum, value)
		_node.OnboardingStatus = value
	}
	if value, ok := cc.mutation.OnboardingStage(); ok {
		_spec.SetField(company.FieldOnboardingStage, field.TypeInt8, value)
		_node.OnboardingStage = value
	}
	if nodes := cc.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.ProfileTable,
			Columns: []string{company.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TerminalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.TerminalsTable,
			Columns: []string{company.TerminalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(terminal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.VehiclesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.VehiclesTable,
			Columns: []string{company.VehiclesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.RoutesTable,
			Columns: []string{company.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TripsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.TripsTable,
			Columns: []string{company.TripsColumn},
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
	if nodes := cc.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.BookingsTable,
			Columns: []string{company.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.IncidentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.IncidentsTable,
			Columns: []string{company.IncidentsColumn},
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
	if nodes := cc.mutation.ParcelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.ParcelsTable,
			Columns: []string{company.ParcelsColumn},
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
	if nodes := cc.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.TransactionsTable,
			Columns: []string{company.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.NotificationsTable,
			Columns: []string{company.NotificationsColumn},
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

// CompanyCreateBulk is the builder for creating many Company entities in bulk.
type CompanyCreateBulk struct {
	config
	err      error
	builders []*CompanyCreate
}

// Save creates the Company entities in the database.
func (ccb *CompanyCreateBulk) Save(ctx context.Context) ([]*Company, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Company, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompanyCreateBulk) SaveX(ctx context.Context) []*Company {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CompanyCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CompanyCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
