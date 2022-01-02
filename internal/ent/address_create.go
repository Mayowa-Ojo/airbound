// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/address"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AddressCreate is the builder for creating a Address entity.
type AddressCreate struct {
	config
	mutation *AddressMutation
	hooks    []Hook
}

// SetStreet sets the "street" field.
func (ac *AddressCreate) SetStreet(s string) *AddressCreate {
	ac.mutation.SetStreet(s)
	return ac
}

// SetCity sets the "city" field.
func (ac *AddressCreate) SetCity(s string) *AddressCreate {
	ac.mutation.SetCity(s)
	return ac
}

// SetState sets the "state" field.
func (ac *AddressCreate) SetState(s string) *AddressCreate {
	ac.mutation.SetState(s)
	return ac
}

// SetZipcode sets the "zipcode" field.
func (ac *AddressCreate) SetZipcode(s string) *AddressCreate {
	ac.mutation.SetZipcode(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AddressCreate) SetCreatedAt(t time.Time) *AddressCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableCreatedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AddressCreate) SetUpdatedAt(t time.Time) *AddressCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableUpdatedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AddressCreate) SetID(u uuid.UUID) *AddressCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AddressMutation object of the builder.
func (ac *AddressCreate) Mutation() *AddressMutation {
	return ac.mutation
}

// Save creates the Address in the database.
func (ac *AddressCreate) Save(ctx context.Context) (*Address, error) {
	var (
		err  error
		node *Address
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
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
func (ac *AddressCreate) SaveX(ctx context.Context) *Address {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AddressCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AddressCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AddressCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := address.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := address.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := address.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddressCreate) check() error {
	if _, ok := ac.mutation.Street(); !ok {
		return &ValidationError{Name: "street", err: errors.New(`ent: missing required field "street"`)}
	}
	if v, ok := ac.mutation.Street(); ok {
		if err := address.StreetValidator(v); err != nil {
			return &ValidationError{Name: "street", err: fmt.Errorf(`ent: validator failed for field "street": %w`, err)}
		}
	}
	if _, ok := ac.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "city"`)}
	}
	if v, ok := ac.mutation.City(); ok {
		if err := address.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "city": %w`, err)}
		}
	}
	if _, ok := ac.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := ac.mutation.State(); ok {
		if err := address.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Zipcode(); !ok {
		return &ValidationError{Name: "zipcode", err: errors.New(`ent: missing required field "zipcode"`)}
	}
	if v, ok := ac.mutation.Zipcode(); ok {
		if err := address.ZipcodeValidator(v); err != nil {
			return &ValidationError{Name: "zipcode", err: fmt.Errorf(`ent: validator failed for field "zipcode": %w`, err)}
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

func (ac *AddressCreate) sqlSave(ctx context.Context) (*Address, error) {
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

func (ac *AddressCreate) createSpec() (*Address, *sqlgraph.CreateSpec) {
	var (
		_node = &Address{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: address.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: address.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Street(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldStreet,
		})
		_node.Street = value
	}
	if value, ok := ac.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldCity,
		})
		_node.City = value
	}
	if value, ok := ac.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldState,
		})
		_node.State = value
	}
	if value, ok := ac.mutation.Zipcode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldZipcode,
		})
		_node.Zipcode = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AddressCreateBulk is the builder for creating many Address entities in bulk.
type AddressCreateBulk struct {
	config
	builders []*AddressCreate
}

// Save creates the Address entities in the database.
func (acb *AddressCreateBulk) Save(ctx context.Context) ([]*Address, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Address, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressMutation)
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
func (acb *AddressCreateBulk) SaveX(ctx context.Context) []*Address {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AddressCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AddressCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
