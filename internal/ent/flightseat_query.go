// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/flightinstance"
	"airbound/internal/ent/flightseat"
	"airbound/internal/ent/passenger"
	"airbound/internal/ent/predicate"
	"airbound/internal/ent/seat"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightSeatQuery is the builder for querying FlightSeat entities.
type FlightSeatQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlightSeat
	// eager-loading edges.
	withFlightInstance *FlightInstanceQuery
	withSeat           *SeatQuery
	withPassenger      *PassengerQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlightSeatQuery builder.
func (fsq *FlightSeatQuery) Where(ps ...predicate.FlightSeat) *FlightSeatQuery {
	fsq.predicates = append(fsq.predicates, ps...)
	return fsq
}

// Limit adds a limit step to the query.
func (fsq *FlightSeatQuery) Limit(limit int) *FlightSeatQuery {
	fsq.limit = &limit
	return fsq
}

// Offset adds an offset step to the query.
func (fsq *FlightSeatQuery) Offset(offset int) *FlightSeatQuery {
	fsq.offset = &offset
	return fsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fsq *FlightSeatQuery) Unique(unique bool) *FlightSeatQuery {
	fsq.unique = &unique
	return fsq
}

// Order adds an order step to the query.
func (fsq *FlightSeatQuery) Order(o ...OrderFunc) *FlightSeatQuery {
	fsq.order = append(fsq.order, o...)
	return fsq
}

// QueryFlightInstance chains the current query on the "flight_instance" edge.
func (fsq *FlightSeatQuery) QueryFlightInstance() *FlightInstanceQuery {
	query := &FlightInstanceQuery{config: fsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flightseat.Table, flightseat.FieldID, selector),
			sqlgraph.To(flightinstance.Table, flightinstance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flightseat.FlightInstanceTable, flightseat.FlightInstanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(fsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySeat chains the current query on the "seat" edge.
func (fsq *FlightSeatQuery) QuerySeat() *SeatQuery {
	query := &SeatQuery{config: fsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flightseat.Table, flightseat.FieldID, selector),
			sqlgraph.To(seat.Table, seat.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, flightseat.SeatTable, flightseat.SeatColumn),
		)
		fromU = sqlgraph.SetNeighbors(fsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPassenger chains the current query on the "passenger" edge.
func (fsq *FlightSeatQuery) QueryPassenger() *PassengerQuery {
	query := &PassengerQuery{config: fsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flightseat.Table, flightseat.FieldID, selector),
			sqlgraph.To(passenger.Table, passenger.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, flightseat.PassengerTable, flightseat.PassengerColumn),
		)
		fromU = sqlgraph.SetNeighbors(fsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlightSeat entity from the query.
// Returns a *NotFoundError when no FlightSeat was found.
func (fsq *FlightSeatQuery) First(ctx context.Context) (*FlightSeat, error) {
	nodes, err := fsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flightseat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fsq *FlightSeatQuery) FirstX(ctx context.Context) *FlightSeat {
	node, err := fsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlightSeat ID from the query.
// Returns a *NotFoundError when no FlightSeat ID was found.
func (fsq *FlightSeatQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flightseat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fsq *FlightSeatQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlightSeat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one FlightSeat entity is not found.
// Returns a *NotFoundError when no FlightSeat entities are found.
func (fsq *FlightSeatQuery) Only(ctx context.Context) (*FlightSeat, error) {
	nodes, err := fsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flightseat.Label}
	default:
		return nil, &NotSingularError{flightseat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fsq *FlightSeatQuery) OnlyX(ctx context.Context) *FlightSeat {
	node, err := fsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlightSeat ID in the query.
// Returns a *NotSingularError when exactly one FlightSeat ID is not found.
// Returns a *NotFoundError when no entities are found.
func (fsq *FlightSeatQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = &NotSingularError{flightseat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fsq *FlightSeatQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlightSeats.
func (fsq *FlightSeatQuery) All(ctx context.Context) ([]*FlightSeat, error) {
	if err := fsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fsq *FlightSeatQuery) AllX(ctx context.Context) []*FlightSeat {
	nodes, err := fsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlightSeat IDs.
func (fsq *FlightSeatQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := fsq.Select(flightseat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fsq *FlightSeatQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fsq *FlightSeatQuery) Count(ctx context.Context) (int, error) {
	if err := fsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fsq *FlightSeatQuery) CountX(ctx context.Context) int {
	count, err := fsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fsq *FlightSeatQuery) Exist(ctx context.Context) (bool, error) {
	if err := fsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fsq *FlightSeatQuery) ExistX(ctx context.Context) bool {
	exist, err := fsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlightSeatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fsq *FlightSeatQuery) Clone() *FlightSeatQuery {
	if fsq == nil {
		return nil
	}
	return &FlightSeatQuery{
		config:             fsq.config,
		limit:              fsq.limit,
		offset:             fsq.offset,
		order:              append([]OrderFunc{}, fsq.order...),
		predicates:         append([]predicate.FlightSeat{}, fsq.predicates...),
		withFlightInstance: fsq.withFlightInstance.Clone(),
		withSeat:           fsq.withSeat.Clone(),
		withPassenger:      fsq.withPassenger.Clone(),
		// clone intermediate query.
		sql:  fsq.sql.Clone(),
		path: fsq.path,
	}
}

// WithFlightInstance tells the query-builder to eager-load the nodes that are connected to
// the "flight_instance" edge. The optional arguments are used to configure the query builder of the edge.
func (fsq *FlightSeatQuery) WithFlightInstance(opts ...func(*FlightInstanceQuery)) *FlightSeatQuery {
	query := &FlightInstanceQuery{config: fsq.config}
	for _, opt := range opts {
		opt(query)
	}
	fsq.withFlightInstance = query
	return fsq
}

// WithSeat tells the query-builder to eager-load the nodes that are connected to
// the "seat" edge. The optional arguments are used to configure the query builder of the edge.
func (fsq *FlightSeatQuery) WithSeat(opts ...func(*SeatQuery)) *FlightSeatQuery {
	query := &SeatQuery{config: fsq.config}
	for _, opt := range opts {
		opt(query)
	}
	fsq.withSeat = query
	return fsq
}

// WithPassenger tells the query-builder to eager-load the nodes that are connected to
// the "passenger" edge. The optional arguments are used to configure the query builder of the edge.
func (fsq *FlightSeatQuery) WithPassenger(opts ...func(*PassengerQuery)) *FlightSeatQuery {
	query := &PassengerQuery{config: fsq.config}
	for _, opt := range opts {
		opt(query)
	}
	fsq.withPassenger = query
	return fsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Fare float64 `json:"fare,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlightSeat.Query().
//		GroupBy(flightseat.FieldFare).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fsq *FlightSeatQuery) GroupBy(field string, fields ...string) *FlightSeatGroupBy {
	group := &FlightSeatGroupBy{config: fsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Fare float64 `json:"fare,omitempty"`
//	}
//
//	client.FlightSeat.Query().
//		Select(flightseat.FieldFare).
//		Scan(ctx, &v)
//
func (fsq *FlightSeatQuery) Select(fields ...string) *FlightSeatSelect {
	fsq.fields = append(fsq.fields, fields...)
	return &FlightSeatSelect{FlightSeatQuery: fsq}
}

func (fsq *FlightSeatQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fsq.fields {
		if !flightseat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fsq.path != nil {
		prev, err := fsq.path(ctx)
		if err != nil {
			return err
		}
		fsq.sql = prev
	}
	return nil
}

func (fsq *FlightSeatQuery) sqlAll(ctx context.Context) ([]*FlightSeat, error) {
	var (
		nodes       = []*FlightSeat{}
		withFKs     = fsq.withFKs
		_spec       = fsq.querySpec()
		loadedTypes = [3]bool{
			fsq.withFlightInstance != nil,
			fsq.withSeat != nil,
			fsq.withPassenger != nil,
		}
	)
	if fsq.withFlightInstance != nil || fsq.withSeat != nil || fsq.withPassenger != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flightseat.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FlightSeat{config: fsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, fsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fsq.withFlightInstance; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*FlightSeat)
		for i := range nodes {
			if nodes[i].flight_instance_id == nil {
				continue
			}
			fk := *nodes[i].flight_instance_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(flightinstance.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flight_instance_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.FlightInstance = n
			}
		}
	}

	if query := fsq.withSeat; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*FlightSeat)
		for i := range nodes {
			if nodes[i].seat_id == nil {
				continue
			}
			fk := *nodes[i].seat_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(seat.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "seat_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Seat = n
			}
		}
	}

	if query := fsq.withPassenger; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*FlightSeat)
		for i := range nodes {
			if nodes[i].passenger_id == nil {
				continue
			}
			fk := *nodes[i].passenger_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(passenger.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "passenger_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Passenger = n
			}
		}
	}

	return nodes, nil
}

func (fsq *FlightSeatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fsq.querySpec()
	_spec.Node.Columns = fsq.fields
	if len(fsq.fields) > 0 {
		_spec.Unique = fsq.unique != nil && *fsq.unique
	}
	return sqlgraph.CountNodes(ctx, fsq.driver, _spec)
}

func (fsq *FlightSeatQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fsq *FlightSeatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flightseat.Table,
			Columns: flightseat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: flightseat.FieldID,
			},
		},
		From:   fsq.sql,
		Unique: true,
	}
	if unique := fsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flightseat.FieldID)
		for i := range fields {
			if fields[i] != flightseat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fsq *FlightSeatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fsq.driver.Dialect())
	t1 := builder.Table(flightseat.Table)
	columns := fsq.fields
	if len(columns) == 0 {
		columns = flightseat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fsq.sql != nil {
		selector = fsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fsq.unique != nil && *fsq.unique {
		selector.Distinct()
	}
	for _, p := range fsq.predicates {
		p(selector)
	}
	for _, p := range fsq.order {
		p(selector)
	}
	if offset := fsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlightSeatGroupBy is the group-by builder for FlightSeat entities.
type FlightSeatGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fsgb *FlightSeatGroupBy) Aggregate(fns ...AggregateFunc) *FlightSeatGroupBy {
	fsgb.fns = append(fsgb.fns, fns...)
	return fsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fsgb *FlightSeatGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fsgb.path(ctx)
	if err != nil {
		return err
	}
	fsgb.sql = query
	return fsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fsgb.fields) > 1 {
		return nil, errors.New("ent: FlightSeatGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) StringsX(ctx context.Context) []string {
	v, err := fsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) StringX(ctx context.Context) string {
	v, err := fsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fsgb.fields) > 1 {
		return nil, errors.New("ent: FlightSeatGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) IntsX(ctx context.Context) []int {
	v, err := fsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) IntX(ctx context.Context) int {
	v, err := fsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fsgb.fields) > 1 {
		return nil, errors.New("ent: FlightSeatGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fsgb.fields) > 1 {
		return nil, errors.New("ent: FlightSeatGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fsgb *FlightSeatGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fsgb *FlightSeatGroupBy) BoolX(ctx context.Context) bool {
	v, err := fsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fsgb *FlightSeatGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fsgb.fields {
		if !flightseat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fsgb *FlightSeatGroupBy) sqlQuery() *sql.Selector {
	selector := fsgb.sql.Select()
	aggregation := make([]string, 0, len(fsgb.fns))
	for _, fn := range fsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fsgb.fields)+len(fsgb.fns))
		for _, f := range fsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fsgb.fields...)...)
}

// FlightSeatSelect is the builder for selecting fields of FlightSeat entities.
type FlightSeatSelect struct {
	*FlightSeatQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fss *FlightSeatSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fss.prepareQuery(ctx); err != nil {
		return err
	}
	fss.sql = fss.FlightSeatQuery.sqlQuery(ctx)
	return fss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fss *FlightSeatSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fss.fields) > 1 {
		return nil, errors.New("ent: FlightSeatSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fss *FlightSeatSelect) StringsX(ctx context.Context) []string {
	v, err := fss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fss *FlightSeatSelect) StringX(ctx context.Context) string {
	v, err := fss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fss.fields) > 1 {
		return nil, errors.New("ent: FlightSeatSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fss *FlightSeatSelect) IntsX(ctx context.Context) []int {
	v, err := fss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fss *FlightSeatSelect) IntX(ctx context.Context) int {
	v, err := fss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fss.fields) > 1 {
		return nil, errors.New("ent: FlightSeatSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fss *FlightSeatSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fss *FlightSeatSelect) Float64X(ctx context.Context) float64 {
	v, err := fss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fss.fields) > 1 {
		return nil, errors.New("ent: FlightSeatSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fss *FlightSeatSelect) BoolsX(ctx context.Context) []bool {
	v, err := fss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fss *FlightSeatSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{flightseat.Label}
	default:
		err = fmt.Errorf("ent: FlightSeatSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fss *FlightSeatSelect) BoolX(ctx context.Context) bool {
	v, err := fss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fss *FlightSeatSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fss.sql.Query()
	if err := fss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
