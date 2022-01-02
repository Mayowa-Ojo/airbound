// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/aircraft"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AircraftCreate is the builder for creating a Aircraft entity.
type AircraftCreate struct {
	config
	mutation *AircraftMutation
	hooks    []Hook
}

// SetTailNumber sets the "tail_number" field.
func (ac *AircraftCreate) SetTailNumber(s string) *AircraftCreate {
	ac.mutation.SetTailNumber(s)
	return ac
}

// SetManufacturer sets the "manufacturer" field.
func (ac *AircraftCreate) SetManufacturer(s string) *AircraftCreate {
	ac.mutation.SetManufacturer(s)
	return ac
}

// SetModel sets the "model" field.
func (ac *AircraftCreate) SetModel(s string) *AircraftCreate {
	ac.mutation.SetModel(s)
	return ac
}

// SetCapacity sets the "capacity" field.
func (ac *AircraftCreate) SetCapacity(i int) *AircraftCreate {
	ac.mutation.SetCapacity(i)
	return ac
}

// SetRange sets the "range" field.
func (ac *AircraftCreate) SetRange(i int) *AircraftCreate {
	ac.mutation.SetRange(i)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AircraftCreate) SetCreatedAt(t time.Time) *AircraftCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableCreatedAt(t *time.Time) *AircraftCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AircraftCreate) SetUpdatedAt(t time.Time) *AircraftCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableUpdatedAt(t *time.Time) *AircraftCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AircraftCreate) SetID(u uuid.UUID) *AircraftCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AircraftMutation object of the builder.
func (ac *AircraftCreate) Mutation() *AircraftMutation {
	return ac.mutation
}

// Save creates the Aircraft in the database.
func (ac *AircraftCreate) Save(ctx context.Context) (*Aircraft, error) {
	var (
		err  error
		node *Aircraft
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AircraftMutation)
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
func (ac *AircraftCreate) SaveX(ctx context.Context) *Aircraft {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AircraftCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AircraftCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AircraftCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := aircraft.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := aircraft.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := aircraft.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AircraftCreate) check() error {
	if _, ok := ac.mutation.TailNumber(); !ok {
		return &ValidationError{Name: "tail_number", err: errors.New(`ent: missing required field "tail_number"`)}
	}
	if v, ok := ac.mutation.TailNumber(); ok {
		if err := aircraft.TailNumberValidator(v); err != nil {
			return &ValidationError{Name: "tail_number", err: fmt.Errorf(`ent: validator failed for field "tail_number": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Manufacturer(); !ok {
		return &ValidationError{Name: "manufacturer", err: errors.New(`ent: missing required field "manufacturer"`)}
	}
	if v, ok := ac.mutation.Manufacturer(); ok {
		if err := aircraft.ManufacturerValidator(v); err != nil {
			return &ValidationError{Name: "manufacturer", err: fmt.Errorf(`ent: validator failed for field "manufacturer": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "model"`)}
	}
	if v, ok := ac.mutation.Model(); ok {
		if err := aircraft.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "model": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Capacity(); !ok {
		return &ValidationError{Name: "capacity", err: errors.New(`ent: missing required field "capacity"`)}
	}
	if v, ok := ac.mutation.Capacity(); ok {
		if err := aircraft.CapacityValidator(v); err != nil {
			return &ValidationError{Name: "capacity", err: fmt.Errorf(`ent: validator failed for field "capacity": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Range(); !ok {
		return &ValidationError{Name: "range", err: errors.New(`ent: missing required field "range"`)}
	}
	if v, ok := ac.mutation.Range(); ok {
		if err := aircraft.RangeValidator(v); err != nil {
			return &ValidationError{Name: "range", err: fmt.Errorf(`ent: validator failed for field "range": %w`, err)}
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

func (ac *AircraftCreate) sqlSave(ctx context.Context) (*Aircraft, error) {
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

func (ac *AircraftCreate) createSpec() (*Aircraft, *sqlgraph.CreateSpec) {
	var (
		_node = &Aircraft{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: aircraft.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: aircraft.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.TailNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldTailNumber,
		})
		_node.TailNumber = value
	}
	if value, ok := ac.mutation.Manufacturer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldManufacturer,
		})
		_node.Manufacturer = value
	}
	if value, ok := ac.mutation.Model(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldModel,
		})
		_node.Model = value
	}
	if value, ok := ac.mutation.Capacity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: aircraft.FieldCapacity,
		})
		_node.Capacity = value
	}
	if value, ok := ac.mutation.Range(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: aircraft.FieldRange,
		})
		_node.Range = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: aircraft.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: aircraft.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AircraftCreateBulk is the builder for creating many Aircraft entities in bulk.
type AircraftCreateBulk struct {
	config
	builders []*AircraftCreate
}

// Save creates the Aircraft entities in the database.
func (acb *AircraftCreateBulk) Save(ctx context.Context) ([]*Aircraft, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Aircraft, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AircraftMutation)
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
func (acb *AircraftCreateBulk) SaveX(ctx context.Context) []*Aircraft {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AircraftCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AircraftCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
