// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/address"
	"airbound/internal/ent/airport"
	"airbound/internal/ent/flight"
	"airbound/internal/ent/frontdesk"
	"airbound/internal/ent/itenerary"
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

// AirportQuery is the builder for querying Airport entities.
type AirportQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Airport
	// eager-loading edges.
	withAddress                *AddressQuery
	withFrontDesks             *FrontDeskQuery
	withDepartureFlights       *FlightQuery
	withArrivalFlights         *FlightQuery
	withOriginIteneraries      *IteneraryQuery
	withDestinationIteneraries *IteneraryQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AirportQuery builder.
func (aq *AirportQuery) Where(ps ...predicate.Airport) *AirportQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AirportQuery) Limit(limit int) *AirportQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AirportQuery) Offset(offset int) *AirportQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AirportQuery) Unique(unique bool) *AirportQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AirportQuery) Order(o ...OrderFunc) *AirportQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAddress chains the current query on the "address" edge.
func (aq *AirportQuery) QueryAddress() *AddressQuery {
	query := &AddressQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, airport.AddressTable, airport.AddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFrontDesks chains the current query on the "front_desks" edge.
func (aq *AirportQuery) QueryFrontDesks() *FrontDeskQuery {
	query := &FrontDeskQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(frontdesk.Table, frontdesk.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, airport.FrontDesksTable, airport.FrontDesksColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepartureFlights chains the current query on the "departure_flights" edge.
func (aq *AirportQuery) QueryDepartureFlights() *FlightQuery {
	query := &FlightQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(flight.Table, flight.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, airport.DepartureFlightsTable, airport.DepartureFlightsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArrivalFlights chains the current query on the "arrival_flights" edge.
func (aq *AirportQuery) QueryArrivalFlights() *FlightQuery {
	query := &FlightQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(flight.Table, flight.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, airport.ArrivalFlightsTable, airport.ArrivalFlightsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOriginIteneraries chains the current query on the "origin_iteneraries" edge.
func (aq *AirportQuery) QueryOriginIteneraries() *IteneraryQuery {
	query := &IteneraryQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(itenerary.Table, itenerary.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, airport.OriginItenerariesTable, airport.OriginItenerariesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDestinationIteneraries chains the current query on the "destination_iteneraries" edge.
func (aq *AirportQuery) QueryDestinationIteneraries() *IteneraryQuery {
	query := &IteneraryQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(itenerary.Table, itenerary.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, airport.DestinationItenerariesTable, airport.DestinationItenerariesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Airport entity from the query.
// Returns a *NotFoundError when no Airport was found.
func (aq *AirportQuery) First(ctx context.Context) (*Airport, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{airport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AirportQuery) FirstX(ctx context.Context) *Airport {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Airport ID from the query.
// Returns a *NotFoundError when no Airport ID was found.
func (aq *AirportQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{airport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AirportQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Airport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Airport entity is not found.
// Returns a *NotFoundError when no Airport entities are found.
func (aq *AirportQuery) Only(ctx context.Context) (*Airport, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{airport.Label}
	default:
		return nil, &NotSingularError{airport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AirportQuery) OnlyX(ctx context.Context) *Airport {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Airport ID in the query.
// Returns a *NotSingularError when exactly one Airport ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aq *AirportQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = &NotSingularError{airport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AirportQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Airports.
func (aq *AirportQuery) All(ctx context.Context) ([]*Airport, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AirportQuery) AllX(ctx context.Context) []*Airport {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Airport IDs.
func (aq *AirportQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := aq.Select(airport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AirportQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AirportQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AirportQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AirportQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AirportQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AirportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AirportQuery) Clone() *AirportQuery {
	if aq == nil {
		return nil
	}
	return &AirportQuery{
		config:                     aq.config,
		limit:                      aq.limit,
		offset:                     aq.offset,
		order:                      append([]OrderFunc{}, aq.order...),
		predicates:                 append([]predicate.Airport{}, aq.predicates...),
		withAddress:                aq.withAddress.Clone(),
		withFrontDesks:             aq.withFrontDesks.Clone(),
		withDepartureFlights:       aq.withDepartureFlights.Clone(),
		withArrivalFlights:         aq.withArrivalFlights.Clone(),
		withOriginIteneraries:      aq.withOriginIteneraries.Clone(),
		withDestinationIteneraries: aq.withDestinationIteneraries.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithAddress tells the query-builder to eager-load the nodes that are connected to
// the "address" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithAddress(opts ...func(*AddressQuery)) *AirportQuery {
	query := &AddressQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAddress = query
	return aq
}

// WithFrontDesks tells the query-builder to eager-load the nodes that are connected to
// the "front_desks" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithFrontDesks(opts ...func(*FrontDeskQuery)) *AirportQuery {
	query := &FrontDeskQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withFrontDesks = query
	return aq
}

// WithDepartureFlights tells the query-builder to eager-load the nodes that are connected to
// the "departure_flights" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithDepartureFlights(opts ...func(*FlightQuery)) *AirportQuery {
	query := &FlightQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withDepartureFlights = query
	return aq
}

// WithArrivalFlights tells the query-builder to eager-load the nodes that are connected to
// the "arrival_flights" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithArrivalFlights(opts ...func(*FlightQuery)) *AirportQuery {
	query := &FlightQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withArrivalFlights = query
	return aq
}

// WithOriginIteneraries tells the query-builder to eager-load the nodes that are connected to
// the "origin_iteneraries" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithOriginIteneraries(opts ...func(*IteneraryQuery)) *AirportQuery {
	query := &IteneraryQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withOriginIteneraries = query
	return aq
}

// WithDestinationIteneraries tells the query-builder to eager-load the nodes that are connected to
// the "destination_iteneraries" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithDestinationIteneraries(opts ...func(*IteneraryQuery)) *AirportQuery {
	query := &IteneraryQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withDestinationIteneraries = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Airport.Query().
//		GroupBy(airport.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AirportQuery) GroupBy(field string, fields ...string) *AirportGroupBy {
	group := &AirportGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Airport.Query().
//		Select(airport.FieldName).
//		Scan(ctx, &v)
//
func (aq *AirportQuery) Select(fields ...string) *AirportSelect {
	aq.fields = append(aq.fields, fields...)
	return &AirportSelect{AirportQuery: aq}
}

func (aq *AirportQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !airport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AirportQuery) sqlAll(ctx context.Context) ([]*Airport, error) {
	var (
		nodes       = []*Airport{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [6]bool{
			aq.withAddress != nil,
			aq.withFrontDesks != nil,
			aq.withDepartureFlights != nil,
			aq.withArrivalFlights != nil,
			aq.withOriginIteneraries != nil,
			aq.withDestinationIteneraries != nil,
		}
	)
	if aq.withAddress != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, airport.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Airport{config: aq.config}
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
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withAddress; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Airport)
		for i := range nodes {
			if nodes[i].address_id == nil {
				continue
			}
			fk := *nodes[i].address_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(address.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "address_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Address = n
			}
		}
	}

	if query := aq.withFrontDesks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Airport)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FrontDesks = []*FrontDesk{}
		}
		query.withFKs = true
		query.Where(predicate.FrontDesk(func(s *sql.Selector) {
			s.Where(sql.InValues(airport.FrontDesksColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.airport_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "airport_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "airport_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FrontDesks = append(node.Edges.FrontDesks, n)
		}
	}

	if query := aq.withDepartureFlights; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Airport)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DepartureFlights = []*Flight{}
		}
		query.withFKs = true
		query.Where(predicate.Flight(func(s *sql.Selector) {
			s.Where(sql.InValues(airport.DepartureFlightsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.depature_airport_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "depature_airport_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "depature_airport_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DepartureFlights = append(node.Edges.DepartureFlights, n)
		}
	}

	if query := aq.withArrivalFlights; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Airport)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ArrivalFlights = []*Flight{}
		}
		query.withFKs = true
		query.Where(predicate.Flight(func(s *sql.Selector) {
			s.Where(sql.InValues(airport.ArrivalFlightsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.arrival_airport_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "arrival_airport_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "arrival_airport_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ArrivalFlights = append(node.Edges.ArrivalFlights, n)
		}
	}

	if query := aq.withOriginIteneraries; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Airport)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OriginIteneraries = []*Itenerary{}
		}
		query.withFKs = true
		query.Where(predicate.Itenerary(func(s *sql.Selector) {
			s.Where(sql.InValues(airport.OriginItenerariesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.origin_airport_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "origin_airport_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "origin_airport_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OriginIteneraries = append(node.Edges.OriginIteneraries, n)
		}
	}

	if query := aq.withDestinationIteneraries; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Airport)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DestinationIteneraries = []*Itenerary{}
		}
		query.withFKs = true
		query.Where(predicate.Itenerary(func(s *sql.Selector) {
			s.Where(sql.InValues(airport.DestinationItenerariesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.destination_airport_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "destination_airport_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "destination_airport_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DestinationIteneraries = append(node.Edges.DestinationIteneraries, n)
		}
	}

	return nodes, nil
}

func (aq *AirportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AirportQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AirportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   airport.Table,
			Columns: airport.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: airport.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airport.FieldID)
		for i := range fields {
			if fields[i] != airport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AirportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(airport.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = airport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AirportGroupBy is the group-by builder for Airport entities.
type AirportGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AirportGroupBy) Aggregate(fns ...AggregateFunc) *AirportGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AirportGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *AirportGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AirportGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *AirportGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = agb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (agb *AirportGroupBy) StringX(ctx context.Context) string {
	v, err := agb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AirportGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *AirportGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = agb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (agb *AirportGroupBy) IntX(ctx context.Context) int {
	v, err := agb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AirportGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *AirportGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = agb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (agb *AirportGroupBy) Float64X(ctx context.Context) float64 {
	v, err := agb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AirportGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *AirportGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AirportGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = agb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (agb *AirportGroupBy) BoolX(ctx context.Context) bool {
	v, err := agb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *AirportGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !airport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AirportGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AirportSelect is the builder for selecting fields of Airport entities.
type AirportSelect struct {
	*AirportQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AirportSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AirportQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *AirportSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AirportSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *AirportSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = as.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (as *AirportSelect) StringX(ctx context.Context) string {
	v, err := as.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AirportSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *AirportSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = as.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (as *AirportSelect) IntX(ctx context.Context) int {
	v, err := as.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AirportSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *AirportSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = as.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (as *AirportSelect) Float64X(ctx context.Context) float64 {
	v, err := as.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AirportSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *AirportSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (as *AirportSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = as.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = fmt.Errorf("ent: AirportSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (as *AirportSelect) BoolX(ctx context.Context) bool {
	v, err := as.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *AirportSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}