// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/airport"
	"airbound/internal/ent/customer"
	"airbound/internal/ent/flightreservation"
	"airbound/internal/ent/itenerary"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// IteneraryCreate is the builder for creating a Itenerary entity.
type IteneraryCreate struct {
	config
	mutation *IteneraryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *IteneraryCreate) SetCreatedAt(t time.Time) *IteneraryCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IteneraryCreate) SetNillableCreatedAt(t *time.Time) *IteneraryCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *IteneraryCreate) SetUpdatedAt(t time.Time) *IteneraryCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *IteneraryCreate) SetNillableUpdatedAt(t *time.Time) *IteneraryCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IteneraryCreate) SetID(u uuid.UUID) *IteneraryCreate {
	ic.mutation.SetID(u)
	return ic
}

// AddFlightReservationIDs adds the "flight_reservations" edge to the FlightReservation entity by IDs.
func (ic *IteneraryCreate) AddFlightReservationIDs(ids ...uuid.UUID) *IteneraryCreate {
	ic.mutation.AddFlightReservationIDs(ids...)
	return ic
}

// AddFlightReservations adds the "flight_reservations" edges to the FlightReservation entity.
func (ic *IteneraryCreate) AddFlightReservations(f ...*FlightReservation) *IteneraryCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ic.AddFlightReservationIDs(ids...)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (ic *IteneraryCreate) SetCustomerID(id uuid.UUID) *IteneraryCreate {
	ic.mutation.SetCustomerID(id)
	return ic
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (ic *IteneraryCreate) SetNillableCustomerID(id *uuid.UUID) *IteneraryCreate {
	if id != nil {
		ic = ic.SetCustomerID(*id)
	}
	return ic
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ic *IteneraryCreate) SetCustomer(c *Customer) *IteneraryCreate {
	return ic.SetCustomerID(c.ID)
}

// SetOriginAirportID sets the "origin_airport" edge to the Airport entity by ID.
func (ic *IteneraryCreate) SetOriginAirportID(id uuid.UUID) *IteneraryCreate {
	ic.mutation.SetOriginAirportID(id)
	return ic
}

// SetNillableOriginAirportID sets the "origin_airport" edge to the Airport entity by ID if the given value is not nil.
func (ic *IteneraryCreate) SetNillableOriginAirportID(id *uuid.UUID) *IteneraryCreate {
	if id != nil {
		ic = ic.SetOriginAirportID(*id)
	}
	return ic
}

// SetOriginAirport sets the "origin_airport" edge to the Airport entity.
func (ic *IteneraryCreate) SetOriginAirport(a *Airport) *IteneraryCreate {
	return ic.SetOriginAirportID(a.ID)
}

// SetDestinationAirportID sets the "destination_airport" edge to the Airport entity by ID.
func (ic *IteneraryCreate) SetDestinationAirportID(id uuid.UUID) *IteneraryCreate {
	ic.mutation.SetDestinationAirportID(id)
	return ic
}

// SetNillableDestinationAirportID sets the "destination_airport" edge to the Airport entity by ID if the given value is not nil.
func (ic *IteneraryCreate) SetNillableDestinationAirportID(id *uuid.UUID) *IteneraryCreate {
	if id != nil {
		ic = ic.SetDestinationAirportID(*id)
	}
	return ic
}

// SetDestinationAirport sets the "destination_airport" edge to the Airport entity.
func (ic *IteneraryCreate) SetDestinationAirport(a *Airport) *IteneraryCreate {
	return ic.SetDestinationAirportID(a.ID)
}

// Mutation returns the IteneraryMutation object of the builder.
func (ic *IteneraryCreate) Mutation() *IteneraryMutation {
	return ic.mutation
}

// Save creates the Itenerary in the database.
func (ic *IteneraryCreate) Save(ctx context.Context) (*Itenerary, error) {
	var (
		err  error
		node *Itenerary
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IteneraryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IteneraryCreate) SaveX(ctx context.Context) *Itenerary {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IteneraryCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IteneraryCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IteneraryCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := itenerary.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := itenerary.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := itenerary.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IteneraryCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (ic *IteneraryCreate) sqlSave(ctx context.Context) (*Itenerary, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (ic *IteneraryCreate) createSpec() (*Itenerary, *sqlgraph.CreateSpec) {
	var (
		_node = &Itenerary{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: itenerary.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itenerary.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: itenerary.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: itenerary.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ic.mutation.FlightReservationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   itenerary.FlightReservationsTable,
			Columns: []string{itenerary.FlightReservationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flightreservation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itenerary.CustomerTable,
			Columns: []string{itenerary.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.OriginAirportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itenerary.OriginAirportTable,
			Columns: []string{itenerary.OriginAirportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: airport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.origin_airport_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.DestinationAirportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itenerary.DestinationAirportTable,
			Columns: []string{itenerary.DestinationAirportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: airport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.destination_airport_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IteneraryCreateBulk is the builder for creating many Itenerary entities in bulk.
type IteneraryCreateBulk struct {
	config
	builders []*IteneraryCreate
}

// Save creates the Itenerary entities in the database.
func (icb *IteneraryCreateBulk) Save(ctx context.Context) ([]*Itenerary, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Itenerary, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IteneraryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IteneraryCreateBulk) SaveX(ctx context.Context) []*Itenerary {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IteneraryCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IteneraryCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}