// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/aircraft"
	"airbound/internal/ent/flight"
	"airbound/internal/ent/flightinstance"
	"airbound/internal/ent/flightreservation"
	"airbound/internal/ent/flightseat"
	"airbound/internal/ent/predicate"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightInstanceQuery is the builder for querying FlightInstance entities.
type FlightInstanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlightInstance
	// eager-loading edges.
	withFlight             *FlightQuery
	withAircraft           *AircraftQuery
	withFlightReservations *FlightReservationQuery
	withFlightSeats        *FlightSeatQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlightInstanceQuery builder.
func (fiq *FlightInstanceQuery) Where(ps ...predicate.FlightInstance) *FlightInstanceQuery {
	fiq.predicates = append(fiq.predicates, ps...)
	return fiq
}

// Limit adds a limit step to the query.
func (fiq *FlightInstanceQuery) Limit(limit int) *FlightInstanceQuery {
	fiq.limit = &limit
	return fiq
}

// Offset adds an offset step to the query.
func (fiq *FlightInstanceQuery) Offset(offset int) *FlightInstanceQuery {
	fiq.offset = &offset
	return fiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fiq *FlightInstanceQuery) Unique(unique bool) *FlightInstanceQuery {
	fiq.unique = &unique
	return fiq
}

// Order adds an order step to the query.
func (fiq *FlightInstanceQuery) Order(o ...OrderFunc) *FlightInstanceQuery {
	fiq.order = append(fiq.order, o...)
	return fiq
}

// QueryFlight chains the current query on the "flight" edge.
func (fiq *FlightInstanceQuery) QueryFlight() *FlightQuery {
	query := &FlightQuery{config: fiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flightinstance.Table, flightinstance.FieldID, selector),
			sqlgraph.To(flight.Table, flight.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flightinstance.FlightTable, flightinstance.FlightColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAircraft chains the current query on the "aircraft" edge.
func (fiq *FlightInstanceQuery) QueryAircraft() *AircraftQuery {
	query := &AircraftQuery{config: fiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flightinstance.Table, flightinstance.FieldID, selector),
			sqlgraph.To(aircraft.Table, aircraft.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, flightinstance.AircraftTable, flightinstance.AircraftColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlightReservations chains the current query on the "flight_reservations" edge.
func (fiq *FlightInstanceQuery) QueryFlightReservations() *FlightReservationQuery {
	query := &FlightReservationQuery{config: fiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flightinstance.Table, flightinstance.FieldID, selector),
			sqlgraph.To(flightreservation.Table, flightreservation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flightinstance.FlightReservationsTable, flightinstance.FlightReservationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlightSeats chains the current query on the "flight_seats" edge.
func (fiq *FlightInstanceQuery) QueryFlightSeats() *FlightSeatQuery {
	query := &FlightSeatQuery{config: fiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flightinstance.Table, flightinstance.FieldID, selector),
			sqlgraph.To(flightseat.Table, flightseat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flightinstance.FlightSeatsTable, flightinstance.FlightSeatsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlightInstance entity from the query.
// Returns a *NotFoundError when no FlightInstance was found.
func (fiq *FlightInstanceQuery) First(ctx context.Context) (*FlightInstance, error) {
	nodes, err := fiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flightinstance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fiq *FlightInstanceQuery) FirstX(ctx context.Context) *FlightInstance {
	node, err := fiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlightInstance ID from the query.
// Returns a *NotFoundError when no FlightInstance ID was found.
func (fiq *FlightInstanceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flightinstance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fiq *FlightInstanceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlightInstance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one FlightInstance entity is not found.
// Returns a *NotFoundError when no FlightInstance entities are found.
func (fiq *FlightInstanceQuery) Only(ctx context.Context) (*FlightInstance, error) {
	nodes, err := fiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flightinstance.Label}
	default:
		return nil, &NotSingularError{flightinstance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fiq *FlightInstanceQuery) OnlyX(ctx context.Context) *FlightInstance {
	node, err := fiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlightInstance ID in the query.
// Returns a *NotSingularError when exactly one FlightInstance ID is not found.
// Returns a *NotFoundError when no entities are found.
func (fiq *FlightInstanceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = &NotSingularError{flightinstance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fiq *FlightInstanceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlightInstances.
func (fiq *FlightInstanceQuery) All(ctx context.Context) ([]*FlightInstance, error) {
	if err := fiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fiq *FlightInstanceQuery) AllX(ctx context.Context) []*FlightInstance {
	nodes, err := fiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlightInstance IDs.
func (fiq *FlightInstanceQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := fiq.Select(flightinstance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fiq *FlightInstanceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fiq *FlightInstanceQuery) Count(ctx context.Context) (int, error) {
	if err := fiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fiq *FlightInstanceQuery) CountX(ctx context.Context) int {
	count, err := fiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fiq *FlightInstanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := fiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fiq *FlightInstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := fiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlightInstanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fiq *FlightInstanceQuery) Clone() *FlightInstanceQuery {
	if fiq == nil {
		return nil
	}
	return &FlightInstanceQuery{
		config:                 fiq.config,
		limit:                  fiq.limit,
		offset:                 fiq.offset,
		order:                  append([]OrderFunc{}, fiq.order...),
		predicates:             append([]predicate.FlightInstance{}, fiq.predicates...),
		withFlight:             fiq.withFlight.Clone(),
		withAircraft:           fiq.withAircraft.Clone(),
		withFlightReservations: fiq.withFlightReservations.Clone(),
		withFlightSeats:        fiq.withFlightSeats.Clone(),
		// clone intermediate query.
		sql:  fiq.sql.Clone(),
		path: fiq.path,
	}
}

// WithFlight tells the query-builder to eager-load the nodes that are connected to
// the "flight" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FlightInstanceQuery) WithFlight(opts ...func(*FlightQuery)) *FlightInstanceQuery {
	query := &FlightQuery{config: fiq.config}
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFlight = query
	return fiq
}

// WithAircraft tells the query-builder to eager-load the nodes that are connected to
// the "aircraft" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FlightInstanceQuery) WithAircraft(opts ...func(*AircraftQuery)) *FlightInstanceQuery {
	query := &AircraftQuery{config: fiq.config}
	for _, opt := range opts {
		opt(query)
	}
	fiq.withAircraft = query
	return fiq
}

// WithFlightReservations tells the query-builder to eager-load the nodes that are connected to
// the "flight_reservations" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FlightInstanceQuery) WithFlightReservations(opts ...func(*FlightReservationQuery)) *FlightInstanceQuery {
	query := &FlightReservationQuery{config: fiq.config}
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFlightReservations = query
	return fiq
}

// WithFlightSeats tells the query-builder to eager-load the nodes that are connected to
// the "flight_seats" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FlightInstanceQuery) WithFlightSeats(opts ...func(*FlightSeatQuery)) *FlightInstanceQuery {
	query := &FlightSeatQuery{config: fiq.config}
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFlightSeats = query
	return fiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DepartureGate int `json:"departure_gate,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlightInstance.Query().
//		GroupBy(flightinstance.FieldDepartureGate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fiq *FlightInstanceQuery) GroupBy(field string, fields ...string) *FlightInstanceGroupBy {
	group := &FlightInstanceGroupBy{config: fiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fiq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DepartureGate int `json:"departure_gate,omitempty"`
//	}
//
//	client.FlightInstance.Query().
//		Select(flightinstance.FieldDepartureGate).
//		Scan(ctx, &v)
//
func (fiq *FlightInstanceQuery) Select(fields ...string) *FlightInstanceSelect {
	fiq.fields = append(fiq.fields, fields...)
	return &FlightInstanceSelect{FlightInstanceQuery: fiq}
}

func (fiq *FlightInstanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fiq.fields {
		if !flightinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fiq.path != nil {
		prev, err := fiq.path(ctx)
		if err != nil {
			return err
		}
		fiq.sql = prev
	}
	return nil
}

func (fiq *FlightInstanceQuery) sqlAll(ctx context.Context) ([]*FlightInstance, error) {
	var (
		nodes       = []*FlightInstance{}
		withFKs     = fiq.withFKs
		_spec       = fiq.querySpec()
		loadedTypes = [4]bool{
			fiq.withFlight != nil,
			fiq.withAircraft != nil,
			fiq.withFlightReservations != nil,
			fiq.withFlightSeats != nil,
		}
	)
	if fiq.withFlight != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flightinstance.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FlightInstance{config: fiq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, fiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fiq.withFlight; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*FlightInstance)
		for i := range nodes {
			if nodes[i].flight_id == nil {
				continue
			}
			fk := *nodes[i].flight_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(flight.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flight_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Flight = n
			}
		}
	}

	if query := fiq.withAircraft; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*FlightInstance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Aircraft(func(s *sql.Selector) {
			s.Where(sql.InValues(flightinstance.AircraftColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.flight_instance_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "flight_instance_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flight_instance_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Aircraft = n
		}
	}

	if query := fiq.withFlightReservations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*FlightInstance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlightReservations = []*FlightReservation{}
		}
		query.withFKs = true
		query.Where(predicate.FlightReservation(func(s *sql.Selector) {
			s.Where(sql.InValues(flightinstance.FlightReservationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.flight_instance_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "flight_instance_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flight_instance_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlightReservations = append(node.Edges.FlightReservations, n)
		}
	}

	if query := fiq.withFlightSeats; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*FlightInstance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlightSeats = []*FlightSeat{}
		}
		query.withFKs = true
		query.Where(predicate.FlightSeat(func(s *sql.Selector) {
			s.Where(sql.InValues(flightinstance.FlightSeatsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.flight_instance_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "flight_instance_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flight_instance_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlightSeats = append(node.Edges.FlightSeats, n)
		}
	}

	return nodes, nil
}

func (fiq *FlightInstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fiq.querySpec()
	return sqlgraph.CountNodes(ctx, fiq.driver, _spec)
}

func (fiq *FlightInstanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fiq *FlightInstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flightinstance.Table,
			Columns: flightinstance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: flightinstance.FieldID,
			},
		},
		From:   fiq.sql,
		Unique: true,
	}
	if unique := fiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flightinstance.FieldID)
		for i := range fields {
			if fields[i] != flightinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fiq *FlightInstanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fiq.driver.Dialect())
	t1 := builder.Table(flightinstance.Table)
	columns := fiq.fields
	if len(columns) == 0 {
		columns = flightinstance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fiq.sql != nil {
		selector = fiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range fiq.predicates {
		p(selector)
	}
	for _, p := range fiq.order {
		p(selector)
	}
	if offset := fiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlightInstanceGroupBy is the group-by builder for FlightInstance entities.
type FlightInstanceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (figb *FlightInstanceGroupBy) Aggregate(fns ...AggregateFunc) *FlightInstanceGroupBy {
	figb.fns = append(figb.fns, fns...)
	return figb
}

// Scan applies the group-by query and scans the result into the given value.
func (figb *FlightInstanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := figb.path(ctx)
	if err != nil {
		return err
	}
	figb.sql = query
	return figb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := figb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(figb.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := figb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := figb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = figb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) StringX(ctx context.Context) string {
	v, err := figb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(figb.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := figb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := figb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = figb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) IntX(ctx context.Context) int {
	v, err := figb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(figb.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := figb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := figb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = figb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := figb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(figb.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := figb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := figb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (figb *FlightInstanceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = figb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (figb *FlightInstanceGroupBy) BoolX(ctx context.Context) bool {
	v, err := figb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (figb *FlightInstanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range figb.fields {
		if !flightinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := figb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := figb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (figb *FlightInstanceGroupBy) sqlQuery() *sql.Selector {
	selector := figb.sql.Select()
	aggregation := make([]string, 0, len(figb.fns))
	for _, fn := range figb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(figb.fields)+len(figb.fns))
		for _, f := range figb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(figb.fields...)...)
}

// FlightInstanceSelect is the builder for selecting fields of FlightInstance entities.
type FlightInstanceSelect struct {
	*FlightInstanceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fis *FlightInstanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fis.prepareQuery(ctx); err != nil {
		return err
	}
	fis.sql = fis.FlightInstanceQuery.sqlQuery(ctx)
	return fis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fis *FlightInstanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fis.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fis *FlightInstanceSelect) StringsX(ctx context.Context) []string {
	v, err := fis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fis *FlightInstanceSelect) StringX(ctx context.Context) string {
	v, err := fis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fis.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fis *FlightInstanceSelect) IntsX(ctx context.Context) []int {
	v, err := fis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fis *FlightInstanceSelect) IntX(ctx context.Context) int {
	v, err := fis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fis.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fis *FlightInstanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fis *FlightInstanceSelect) Float64X(ctx context.Context) float64 {
	v, err := fis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fis.fields) > 1 {
		return nil, errors.New("ent: FlightInstanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fis *FlightInstanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := fis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fis *FlightInstanceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightinstance.Label}
	default:
		err = fmt.Errorf("ent: FlightInstanceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fis *FlightInstanceSelect) BoolX(ctx context.Context) bool {
	v, err := fis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fis *FlightInstanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fis.sql.Query()
	if err := fis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}