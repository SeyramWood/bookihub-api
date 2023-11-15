// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/incident"
	"github.com/SeyramWood/bookibus/ent/notification"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/trip"
	"github.com/SeyramWood/bookibus/ent/user"
)

// CompanyUserQuery is the builder for querying CompanyUser entities.
type CompanyUserQuery struct {
	config
	ctx               *QueryContext
	order             []companyuser.OrderOption
	inters            []Interceptor
	predicates        []predicate.CompanyUser
	withProfile       *UserQuery
	withTrips         *TripQuery
	withIncidents     *IncidentQuery
	withParcels       *ParcelQuery
	withNotifications *NotificationQuery
	withCompany       *CompanyQuery
	withFKs           bool
	modifiers         []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompanyUserQuery builder.
func (cuq *CompanyUserQuery) Where(ps ...predicate.CompanyUser) *CompanyUserQuery {
	cuq.predicates = append(cuq.predicates, ps...)
	return cuq
}

// Limit the number of records to be returned by this query.
func (cuq *CompanyUserQuery) Limit(limit int) *CompanyUserQuery {
	cuq.ctx.Limit = &limit
	return cuq
}

// Offset to start from.
func (cuq *CompanyUserQuery) Offset(offset int) *CompanyUserQuery {
	cuq.ctx.Offset = &offset
	return cuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cuq *CompanyUserQuery) Unique(unique bool) *CompanyUserQuery {
	cuq.ctx.Unique = &unique
	return cuq
}

// Order specifies how the records should be ordered.
func (cuq *CompanyUserQuery) Order(o ...companyuser.OrderOption) *CompanyUserQuery {
	cuq.order = append(cuq.order, o...)
	return cuq
}

// QueryProfile chains the current query on the "profile" edge.
func (cuq *CompanyUserQuery) QueryProfile() *UserQuery {
	query := (&UserClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companyuser.Table, companyuser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, companyuser.ProfileTable, companyuser.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTrips chains the current query on the "trips" edge.
func (cuq *CompanyUserQuery) QueryTrips() *TripQuery {
	query := (&TripClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companyuser.Table, companyuser.FieldID, selector),
			sqlgraph.To(trip.Table, trip.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, companyuser.TripsTable, companyuser.TripsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIncidents chains the current query on the "incidents" edge.
func (cuq *CompanyUserQuery) QueryIncidents() *IncidentQuery {
	query := (&IncidentClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companyuser.Table, companyuser.FieldID, selector),
			sqlgraph.To(incident.Table, incident.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, companyuser.IncidentsTable, companyuser.IncidentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParcels chains the current query on the "parcels" edge.
func (cuq *CompanyUserQuery) QueryParcels() *ParcelQuery {
	query := (&ParcelClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companyuser.Table, companyuser.FieldID, selector),
			sqlgraph.To(parcel.Table, parcel.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, companyuser.ParcelsTable, companyuser.ParcelsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifications chains the current query on the "notifications" edge.
func (cuq *CompanyUserQuery) QueryNotifications() *NotificationQuery {
	query := (&NotificationClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companyuser.Table, companyuser.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, companyuser.NotificationsTable, companyuser.NotificationsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompany chains the current query on the "company" edge.
func (cuq *CompanyUserQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companyuser.Table, companyuser.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, companyuser.CompanyTable, companyuser.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CompanyUser entity from the query.
// Returns a *NotFoundError when no CompanyUser was found.
func (cuq *CompanyUserQuery) First(ctx context.Context) (*CompanyUser, error) {
	nodes, err := cuq.Limit(1).All(setContextOp(ctx, cuq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{companyuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cuq *CompanyUserQuery) FirstX(ctx context.Context) *CompanyUser {
	node, err := cuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CompanyUser ID from the query.
// Returns a *NotFoundError when no CompanyUser ID was found.
func (cuq *CompanyUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cuq.Limit(1).IDs(setContextOp(ctx, cuq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{companyuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cuq *CompanyUserQuery) FirstIDX(ctx context.Context) int {
	id, err := cuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CompanyUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CompanyUser entity is found.
// Returns a *NotFoundError when no CompanyUser entities are found.
func (cuq *CompanyUserQuery) Only(ctx context.Context) (*CompanyUser, error) {
	nodes, err := cuq.Limit(2).All(setContextOp(ctx, cuq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{companyuser.Label}
	default:
		return nil, &NotSingularError{companyuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cuq *CompanyUserQuery) OnlyX(ctx context.Context) *CompanyUser {
	node, err := cuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CompanyUser ID in the query.
// Returns a *NotSingularError when more than one CompanyUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (cuq *CompanyUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cuq.Limit(2).IDs(setContextOp(ctx, cuq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{companyuser.Label}
	default:
		err = &NotSingularError{companyuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cuq *CompanyUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := cuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CompanyUsers.
func (cuq *CompanyUserQuery) All(ctx context.Context) ([]*CompanyUser, error) {
	ctx = setContextOp(ctx, cuq.ctx, "All")
	if err := cuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CompanyUser, *CompanyUserQuery]()
	return withInterceptors[[]*CompanyUser](ctx, cuq, qr, cuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cuq *CompanyUserQuery) AllX(ctx context.Context) []*CompanyUser {
	nodes, err := cuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CompanyUser IDs.
func (cuq *CompanyUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cuq.ctx.Unique == nil && cuq.path != nil {
		cuq.Unique(true)
	}
	ctx = setContextOp(ctx, cuq.ctx, "IDs")
	if err = cuq.Select(companyuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cuq *CompanyUserQuery) IDsX(ctx context.Context) []int {
	ids, err := cuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cuq *CompanyUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cuq.ctx, "Count")
	if err := cuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cuq, querierCount[*CompanyUserQuery](), cuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cuq *CompanyUserQuery) CountX(ctx context.Context) int {
	count, err := cuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cuq *CompanyUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cuq.ctx, "Exist")
	switch _, err := cuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cuq *CompanyUserQuery) ExistX(ctx context.Context) bool {
	exist, err := cuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompanyUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cuq *CompanyUserQuery) Clone() *CompanyUserQuery {
	if cuq == nil {
		return nil
	}
	return &CompanyUserQuery{
		config:            cuq.config,
		ctx:               cuq.ctx.Clone(),
		order:             append([]companyuser.OrderOption{}, cuq.order...),
		inters:            append([]Interceptor{}, cuq.inters...),
		predicates:        append([]predicate.CompanyUser{}, cuq.predicates...),
		withProfile:       cuq.withProfile.Clone(),
		withTrips:         cuq.withTrips.Clone(),
		withIncidents:     cuq.withIncidents.Clone(),
		withParcels:       cuq.withParcels.Clone(),
		withNotifications: cuq.withNotifications.Clone(),
		withCompany:       cuq.withCompany.Clone(),
		// clone intermediate query.
		sql:  cuq.sql.Clone(),
		path: cuq.path,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *CompanyUserQuery) WithProfile(opts ...func(*UserQuery)) *CompanyUserQuery {
	query := (&UserClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withProfile = query
	return cuq
}

// WithTrips tells the query-builder to eager-load the nodes that are connected to
// the "trips" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *CompanyUserQuery) WithTrips(opts ...func(*TripQuery)) *CompanyUserQuery {
	query := (&TripClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withTrips = query
	return cuq
}

// WithIncidents tells the query-builder to eager-load the nodes that are connected to
// the "incidents" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *CompanyUserQuery) WithIncidents(opts ...func(*IncidentQuery)) *CompanyUserQuery {
	query := (&IncidentClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withIncidents = query
	return cuq
}

// WithParcels tells the query-builder to eager-load the nodes that are connected to
// the "parcels" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *CompanyUserQuery) WithParcels(opts ...func(*ParcelQuery)) *CompanyUserQuery {
	query := (&ParcelClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withParcels = query
	return cuq
}

// WithNotifications tells the query-builder to eager-load the nodes that are connected to
// the "notifications" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *CompanyUserQuery) WithNotifications(opts ...func(*NotificationQuery)) *CompanyUserQuery {
	query := (&NotificationClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withNotifications = query
	return cuq
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *CompanyUserQuery) WithCompany(opts ...func(*CompanyQuery)) *CompanyUserQuery {
	query := (&CompanyClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withCompany = query
	return cuq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CompanyUser.Query().
//		GroupBy(companyuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cuq *CompanyUserQuery) GroupBy(field string, fields ...string) *CompanyUserGroupBy {
	cuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CompanyUserGroupBy{build: cuq}
	grbuild.flds = &cuq.ctx.Fields
	grbuild.label = companyuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.CompanyUser.Query().
//		Select(companyuser.FieldCreatedAt).
//		Scan(ctx, &v)
func (cuq *CompanyUserQuery) Select(fields ...string) *CompanyUserSelect {
	cuq.ctx.Fields = append(cuq.ctx.Fields, fields...)
	sbuild := &CompanyUserSelect{CompanyUserQuery: cuq}
	sbuild.label = companyuser.Label
	sbuild.flds, sbuild.scan = &cuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompanyUserSelect configured with the given aggregations.
func (cuq *CompanyUserQuery) Aggregate(fns ...AggregateFunc) *CompanyUserSelect {
	return cuq.Select().Aggregate(fns...)
}

func (cuq *CompanyUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cuq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cuq); err != nil {
				return err
			}
		}
	}
	for _, f := range cuq.ctx.Fields {
		if !companyuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cuq.path != nil {
		prev, err := cuq.path(ctx)
		if err != nil {
			return err
		}
		cuq.sql = prev
	}
	return nil
}

func (cuq *CompanyUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CompanyUser, error) {
	var (
		nodes       = []*CompanyUser{}
		withFKs     = cuq.withFKs
		_spec       = cuq.querySpec()
		loadedTypes = [6]bool{
			cuq.withProfile != nil,
			cuq.withTrips != nil,
			cuq.withIncidents != nil,
			cuq.withParcels != nil,
			cuq.withNotifications != nil,
			cuq.withCompany != nil,
		}
	)
	if cuq.withCompany != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, companyuser.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CompanyUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CompanyUser{config: cuq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cuq.modifiers) > 0 {
		_spec.Modifiers = cuq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cuq.withProfile; query != nil {
		if err := cuq.loadProfile(ctx, query, nodes, nil,
			func(n *CompanyUser, e *User) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	if query := cuq.withTrips; query != nil {
		if err := cuq.loadTrips(ctx, query, nodes,
			func(n *CompanyUser) { n.Edges.Trips = []*Trip{} },
			func(n *CompanyUser, e *Trip) { n.Edges.Trips = append(n.Edges.Trips, e) }); err != nil {
			return nil, err
		}
	}
	if query := cuq.withIncidents; query != nil {
		if err := cuq.loadIncidents(ctx, query, nodes,
			func(n *CompanyUser) { n.Edges.Incidents = []*Incident{} },
			func(n *CompanyUser, e *Incident) { n.Edges.Incidents = append(n.Edges.Incidents, e) }); err != nil {
			return nil, err
		}
	}
	if query := cuq.withParcels; query != nil {
		if err := cuq.loadParcels(ctx, query, nodes,
			func(n *CompanyUser) { n.Edges.Parcels = []*Parcel{} },
			func(n *CompanyUser, e *Parcel) { n.Edges.Parcels = append(n.Edges.Parcels, e) }); err != nil {
			return nil, err
		}
	}
	if query := cuq.withNotifications; query != nil {
		if err := cuq.loadNotifications(ctx, query, nodes,
			func(n *CompanyUser) { n.Edges.Notifications = []*Notification{} },
			func(n *CompanyUser, e *Notification) { n.Edges.Notifications = append(n.Edges.Notifications, e) }); err != nil {
			return nil, err
		}
	}
	if query := cuq.withCompany; query != nil {
		if err := cuq.loadCompany(ctx, query, nodes, nil,
			func(n *CompanyUser, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cuq *CompanyUserQuery) loadProfile(ctx context.Context, query *UserQuery, nodes []*CompanyUser, init func(*CompanyUser), assign func(*CompanyUser, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*CompanyUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(companyuser.ProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_user_profile
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_user_profile" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_user_profile" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cuq *CompanyUserQuery) loadTrips(ctx context.Context, query *TripQuery, nodes []*CompanyUser, init func(*CompanyUser), assign func(*CompanyUser, *Trip)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*CompanyUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Trip(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(companyuser.TripsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_user_trips
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_user_trips" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_user_trips" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cuq *CompanyUserQuery) loadIncidents(ctx context.Context, query *IncidentQuery, nodes []*CompanyUser, init func(*CompanyUser), assign func(*CompanyUser, *Incident)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*CompanyUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Incident(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(companyuser.IncidentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_user_incidents
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_user_incidents" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_user_incidents" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cuq *CompanyUserQuery) loadParcels(ctx context.Context, query *ParcelQuery, nodes []*CompanyUser, init func(*CompanyUser), assign func(*CompanyUser, *Parcel)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*CompanyUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Parcel(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(companyuser.ParcelsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_user_parcels
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_user_parcels" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_user_parcels" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cuq *CompanyUserQuery) loadNotifications(ctx context.Context, query *NotificationQuery, nodes []*CompanyUser, init func(*CompanyUser), assign func(*CompanyUser, *Notification)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*CompanyUser)
	nids := make(map[int]map[*CompanyUser]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(companyuser.NotificationsTable)
		s.Join(joinT).On(s.C(notification.FieldID), joinT.C(companyuser.NotificationsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(companyuser.NotificationsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(companyuser.NotificationsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*CompanyUser]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Notification](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "notifications" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cuq *CompanyUserQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*CompanyUser, init func(*CompanyUser), assign func(*CompanyUser, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CompanyUser)
	for i := range nodes {
		if nodes[i].company_profile == nil {
			continue
		}
		fk := *nodes[i].company_profile
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_profile" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cuq *CompanyUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cuq.querySpec()
	if len(cuq.modifiers) > 0 {
		_spec.Modifiers = cuq.modifiers
	}
	_spec.Node.Columns = cuq.ctx.Fields
	if len(cuq.ctx.Fields) > 0 {
		_spec.Unique = cuq.ctx.Unique != nil && *cuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cuq.driver, _spec)
}

func (cuq *CompanyUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(companyuser.Table, companyuser.Columns, sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt))
	_spec.From = cuq.sql
	if unique := cuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cuq.path != nil {
		_spec.Unique = true
	}
	if fields := cuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companyuser.FieldID)
		for i := range fields {
			if fields[i] != companyuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cuq *CompanyUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cuq.driver.Dialect())
	t1 := builder.Table(companyuser.Table)
	columns := cuq.ctx.Fields
	if len(columns) == 0 {
		columns = companyuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cuq.sql != nil {
		selector = cuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cuq.ctx.Unique != nil && *cuq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cuq.modifiers {
		m(selector)
	}
	for _, p := range cuq.predicates {
		p(selector)
	}
	for _, p := range cuq.order {
		p(selector)
	}
	if offset := cuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cuq *CompanyUserQuery) Modify(modifiers ...func(s *sql.Selector)) *CompanyUserSelect {
	cuq.modifiers = append(cuq.modifiers, modifiers...)
	return cuq.Select()
}

// CompanyUserGroupBy is the group-by builder for CompanyUser entities.
type CompanyUserGroupBy struct {
	selector
	build *CompanyUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cugb *CompanyUserGroupBy) Aggregate(fns ...AggregateFunc) *CompanyUserGroupBy {
	cugb.fns = append(cugb.fns, fns...)
	return cugb
}

// Scan applies the selector query and scans the result into the given value.
func (cugb *CompanyUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cugb.build.ctx, "GroupBy")
	if err := cugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyUserQuery, *CompanyUserGroupBy](ctx, cugb.build, cugb, cugb.build.inters, v)
}

func (cugb *CompanyUserGroupBy) sqlScan(ctx context.Context, root *CompanyUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cugb.fns))
	for _, fn := range cugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cugb.flds)+len(cugb.fns))
		for _, f := range *cugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompanyUserSelect is the builder for selecting fields of CompanyUser entities.
type CompanyUserSelect struct {
	*CompanyUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cus *CompanyUserSelect) Aggregate(fns ...AggregateFunc) *CompanyUserSelect {
	cus.fns = append(cus.fns, fns...)
	return cus
}

// Scan applies the selector query and scans the result into the given value.
func (cus *CompanyUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cus.ctx, "Select")
	if err := cus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyUserQuery, *CompanyUserSelect](ctx, cus.CompanyUserQuery, cus, cus.inters, v)
}

func (cus *CompanyUserSelect) sqlScan(ctx context.Context, root *CompanyUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cus.fns))
	for _, fn := range cus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cus *CompanyUserSelect) Modify(modifiers ...func(s *sql.Selector)) *CompanyUserSelect {
	cus.modifiers = append(cus.modifiers, modifiers...)
	return cus
}
