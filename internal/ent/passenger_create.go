// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/flightreservation"
	"airbound/internal/ent/flightseat"
	"airbound/internal/ent/passenger"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PassengerCreate is the builder for creating a Passenger entity.
type PassengerCreate struct {
	config
	mutation *PassengerMutation
	hooks    []Hook
}

// SetFirstname sets the "firstname" field.
func (pc *PassengerCreate) SetFirstname(s string) *PassengerCreate {
	pc.mutation.SetFirstname(s)
	return pc
}

// SetLastname sets the "lastname" field.
func (pc *PassengerCreate) SetLastname(s string) *PassengerCreate {
	pc.mutation.SetLastname(s)
	return pc
}

// SetAge sets the "age" field.
func (pc *PassengerCreate) SetAge(i int) *PassengerCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetNationality sets the "nationality" field.
func (pc *PassengerCreate) SetNationality(s string) *PassengerCreate {
	pc.mutation.SetNationality(s)
	return pc
}

// SetPassportNumber sets the "passport_number" field.
func (pc *PassengerCreate) SetPassportNumber(s string) *PassengerCreate {
	pc.mutation.SetPassportNumber(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PassengerCreate) SetCreatedAt(t time.Time) *PassengerCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PassengerCreate) SetNillableCreatedAt(t *time.Time) *PassengerCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PassengerCreate) SetUpdatedAt(t time.Time) *PassengerCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PassengerCreate) SetNillableUpdatedAt(t *time.Time) *PassengerCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PassengerCreate) SetID(u uuid.UUID) *PassengerCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PassengerCreate) SetNillableID(u *uuid.UUID) *PassengerCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetFlightReservationID sets the "flight_reservation" edge to the FlightReservation entity by ID.
func (pc *PassengerCreate) SetFlightReservationID(id uuid.UUID) *PassengerCreate {
	pc.mutation.SetFlightReservationID(id)
	return pc
}

// SetNillableFlightReservationID sets the "flight_reservation" edge to the FlightReservation entity by ID if the given value is not nil.
func (pc *PassengerCreate) SetNillableFlightReservationID(id *uuid.UUID) *PassengerCreate {
	if id != nil {
		pc = pc.SetFlightReservationID(*id)
	}
	return pc
}

// SetFlightReservation sets the "flight_reservation" edge to the FlightReservation entity.
func (pc *PassengerCreate) SetFlightReservation(f *FlightReservation) *PassengerCreate {
	return pc.SetFlightReservationID(f.ID)
}

// SetFlightSeatID sets the "flight_seat" edge to the FlightSeat entity by ID.
func (pc *PassengerCreate) SetFlightSeatID(id uuid.UUID) *PassengerCreate {
	pc.mutation.SetFlightSeatID(id)
	return pc
}

// SetNillableFlightSeatID sets the "flight_seat" edge to the FlightSeat entity by ID if the given value is not nil.
func (pc *PassengerCreate) SetNillableFlightSeatID(id *uuid.UUID) *PassengerCreate {
	if id != nil {
		pc = pc.SetFlightSeatID(*id)
	}
	return pc
}

// SetFlightSeat sets the "flight_seat" edge to the FlightSeat entity.
func (pc *PassengerCreate) SetFlightSeat(f *FlightSeat) *PassengerCreate {
	return pc.SetFlightSeatID(f.ID)
}

// Mutation returns the PassengerMutation object of the builder.
func (pc *PassengerCreate) Mutation() *PassengerMutation {
	return pc.mutation
}

// Save creates the Passenger in the database.
func (pc *PassengerCreate) Save(ctx context.Context) (*Passenger, error) {
	var (
		err  error
		node *Passenger
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PassengerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PassengerCreate) SaveX(ctx context.Context) *Passenger {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PassengerCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PassengerCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PassengerCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := passenger.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := passenger.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := passenger.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PassengerCreate) check() error {
	if _, ok := pc.mutation.Firstname(); !ok {
		return &ValidationError{Name: "firstname", err: errors.New(`ent: missing required field "Passenger.firstname"`)}
	}
	if v, ok := pc.mutation.Firstname(); ok {
		if err := passenger.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf(`ent: validator failed for field "Passenger.firstname": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Lastname(); !ok {
		return &ValidationError{Name: "lastname", err: errors.New(`ent: missing required field "Passenger.lastname"`)}
	}
	if v, ok := pc.mutation.Lastname(); ok {
		if err := passenger.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf(`ent: validator failed for field "Passenger.lastname": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Passenger.age"`)}
	}
	if v, ok := pc.mutation.Age(); ok {
		if err := passenger.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "Passenger.age": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Nationality(); !ok {
		return &ValidationError{Name: "nationality", err: errors.New(`ent: missing required field "Passenger.nationality"`)}
	}
	if v, ok := pc.mutation.Nationality(); ok {
		if err := passenger.NationalityValidator(v); err != nil {
			return &ValidationError{Name: "nationality", err: fmt.Errorf(`ent: validator failed for field "Passenger.nationality": %w`, err)}
		}
	}
	if _, ok := pc.mutation.PassportNumber(); !ok {
		return &ValidationError{Name: "passport_number", err: errors.New(`ent: missing required field "Passenger.passport_number"`)}
	}
	if v, ok := pc.mutation.PassportNumber(); ok {
		if err := passenger.PassportNumberValidator(v); err != nil {
			return &ValidationError{Name: "passport_number", err: fmt.Errorf(`ent: validator failed for field "Passenger.passport_number": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Passenger.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Passenger.updated_at"`)}
	}
	return nil
}

func (pc *PassengerCreate) sqlSave(ctx context.Context) (*Passenger, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pc *PassengerCreate) createSpec() (*Passenger, *sqlgraph.CreateSpec) {
	var (
		_node = &Passenger{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: passenger.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: passenger.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Firstname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passenger.FieldFirstname,
		})
		_node.Firstname = value
	}
	if value, ok := pc.mutation.Lastname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passenger.FieldLastname,
		})
		_node.Lastname = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: passenger.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := pc.mutation.Nationality(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passenger.FieldNationality,
		})
		_node.Nationality = value
	}
	if value, ok := pc.mutation.PassportNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passenger.FieldPassportNumber,
		})
		_node.PassportNumber = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: passenger.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: passenger.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.FlightReservationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passenger.FlightReservationTable,
			Columns: []string{passenger.FlightReservationColumn},
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
		_node.flight_reservation_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FlightSeatIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   passenger.FlightSeatTable,
			Columns: []string{passenger.FlightSeatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flightseat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PassengerCreateBulk is the builder for creating many Passenger entities in bulk.
type PassengerCreateBulk struct {
	config
	builders []*PassengerCreate
}

// Save creates the Passenger entities in the database.
func (pcb *PassengerCreateBulk) Save(ctx context.Context) ([]*Passenger, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Passenger, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PassengerMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PassengerCreateBulk) SaveX(ctx context.Context) []*Passenger {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PassengerCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PassengerCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
