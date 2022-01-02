// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/airport"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AirportCreate is the builder for creating a Airport entity.
type AirportCreate struct {
	config
	mutation *AirportMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AirportCreate) SetName(s string) *AirportCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetIataCode sets the "iata_code" field.
func (ac *AirportCreate) SetIataCode(s string) *AirportCreate {
	ac.mutation.SetIataCode(s)
	return ac
}

// SetIcaoCode sets the "icao_code" field.
func (ac *AirportCreate) SetIcaoCode(s string) *AirportCreate {
	ac.mutation.SetIcaoCode(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AirportCreate) SetCreatedAt(t time.Time) *AirportCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AirportCreate) SetNillableCreatedAt(t *time.Time) *AirportCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AirportCreate) SetUpdatedAt(t time.Time) *AirportCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AirportCreate) SetNillableUpdatedAt(t *time.Time) *AirportCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AirportCreate) SetID(u uuid.UUID) *AirportCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AirportMutation object of the builder.
func (ac *AirportCreate) Mutation() *AirportMutation {
	return ac.mutation
}

// Save creates the Airport in the database.
func (ac *AirportCreate) Save(ctx context.Context) (*Airport, error) {
	var (
		err  error
		node *Airport
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AirportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AirportCreate) SaveX(ctx context.Context) *Airport {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AirportCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AirportCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AirportCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := airport.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := airport.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := airport.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AirportCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := airport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.IataCode(); !ok {
		return &ValidationError{Name: "iata_code", err: errors.New(`ent: missing required field "iata_code"`)}
	}
	if v, ok := ac.mutation.IataCode(); ok {
		if err := airport.IataCodeValidator(v); err != nil {
			return &ValidationError{Name: "iata_code", err: fmt.Errorf(`ent: validator failed for field "iata_code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.IcaoCode(); !ok {
		return &ValidationError{Name: "icao_code", err: errors.New(`ent: missing required field "icao_code"`)}
	}
	if v, ok := ac.mutation.IcaoCode(); ok {
		if err := airport.IcaoCodeValidator(v); err != nil {
			return &ValidationError{Name: "icao_code", err: fmt.Errorf(`ent: validator failed for field "icao_code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (ac *AirportCreate) sqlSave(ctx context.Context) (*Airport, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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

func (ac *AirportCreate) createSpec() (*Airport, *sqlgraph.CreateSpec) {
	var (
		_node = &Airport{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: airport.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: airport.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: airport.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.IataCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: airport.FieldIataCode,
		})
		_node.IataCode = value
	}
	if value, ok := ac.mutation.IcaoCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: airport.FieldIcaoCode,
		})
		_node.IcaoCode = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: airport.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: airport.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AirportCreateBulk is the builder for creating many Airport entities in bulk.
type AirportCreateBulk struct {
	config
	builders []*AirportCreate
}

// Save creates the Airport entities in the database.
func (acb *AirportCreateBulk) Save(ctx context.Context) ([]*Airport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Airport, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AirportMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AirportCreateBulk) SaveX(ctx context.Context) []*Airport {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AirportCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AirportCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
