// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/account"
	"airbound/internal/ent/enums"
	"airbound/internal/ent/predicate"
	"airbound/internal/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAccountStatus sets the "account_status" field.
func (au *AccountUpdate) SetAccountStatus(es enums.AccountStatus) *AccountUpdate {
	au.mutation.SetAccountStatus(es)
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(b []byte) *AccountUpdate {
	au.mutation.SetPassword(b)
	return au
}

// SetSalt sets the "salt" field.
func (au *AccountUpdate) SetSalt(b []byte) *AccountUpdate {
	au.mutation.SetSalt(b)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccountUpdate) SetCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *AccountUpdate) SetUserID(id uuid.UUID) *AccountUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *AccountUpdate) SetUser(u *User) *AccountUpdate {
	return au.SetUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *AccountUpdate) ClearUser() *AccountUpdate {
	au.mutation.ClearUser()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.AccountStatus(); ok {
		if err := account.AccountStatusValidator(v); err != nil {
			return &ValidationError{Name: "account_status", err: fmt.Errorf("ent: validator failed for field \"account_status\": %w", err)}
		}
	}
	if _, ok := au.mutation.UserID(); au.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.AccountStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: account.FieldAccountStatus,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := au.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: account.FieldSalt,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   account.UserTable,
			Columns: []string{account.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   account.UserTable,
			Columns: []string{account.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetAccountStatus sets the "account_status" field.
func (auo *AccountUpdateOne) SetAccountStatus(es enums.AccountStatus) *AccountUpdateOne {
	auo.mutation.SetAccountStatus(es)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(b []byte) *AccountUpdateOne {
	auo.mutation.SetPassword(b)
	return auo
}

// SetSalt sets the "salt" field.
func (auo *AccountUpdateOne) SetSalt(b []byte) *AccountUpdateOne {
	auo.mutation.SetSalt(b)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccountUpdateOne) SetCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *AccountUpdateOne) SetUserID(id uuid.UUID) *AccountUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *AccountUpdateOne) SetUser(u *User) *AccountUpdateOne {
	return auo.SetUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AccountUpdateOne) ClearUser() *AccountUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.AccountStatus(); ok {
		if err := account.AccountStatusValidator(v); err != nil {
			return &ValidationError{Name: "account_status", err: fmt.Errorf("ent: validator failed for field \"account_status\": %w", err)}
		}
	}
	if _, ok := auo.mutation.UserID(); auo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Account.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.AccountStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: account.FieldAccountStatus,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := auo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: account.FieldSalt,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   account.UserTable,
			Columns: []string{account.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   account.UserTable,
			Columns: []string{account.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
