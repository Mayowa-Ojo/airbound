// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/admin"
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

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// Where appends a list predicates to the AdminUpdate builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetTwoFaSecret sets the "two_fa_secret" field.
func (au *AdminUpdate) SetTwoFaSecret(s string) *AdminUpdate {
	au.mutation.SetTwoFaSecret(s)
	return au
}

// SetNillableTwoFaSecret sets the "two_fa_secret" field if the given value is not nil.
func (au *AdminUpdate) SetNillableTwoFaSecret(s *string) *AdminUpdate {
	if s != nil {
		au.SetTwoFaSecret(*s)
	}
	return au
}

// ClearTwoFaSecret clears the value of the "two_fa_secret" field.
func (au *AdminUpdate) ClearTwoFaSecret() *AdminUpdate {
	au.mutation.ClearTwoFaSecret()
	return au
}

// SetTwoFaCompleted sets the "two_fa_completed" field.
func (au *AdminUpdate) SetTwoFaCompleted(b bool) *AdminUpdate {
	au.mutation.SetTwoFaCompleted(b)
	return au
}

// SetNillableTwoFaCompleted sets the "two_fa_completed" field if the given value is not nil.
func (au *AdminUpdate) SetNillableTwoFaCompleted(b *bool) *AdminUpdate {
	if b != nil {
		au.SetTwoFaCompleted(*b)
	}
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AdminUpdate) SetCreatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AdminUpdate) SetNillableCreatedAt(t *time.Time) *AdminUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AdminUpdate) SetUpdatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *AdminUpdate) SetUserID(id uuid.UUID) *AdminUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *AdminUpdate) SetUser(u *User) *AdminUpdate {
	return au.SetUserID(u.ID)
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *AdminUpdate) ClearUser() *AdminUpdate {
	au.mutation.ClearUser()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*AdminMutation)
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
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AdminUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := admin.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AdminUpdate) check() error {
	if v, ok := au.mutation.TwoFaSecret(); ok {
		if err := admin.TwoFaSecretValidator(v); err != nil {
			return &ValidationError{Name: "two_fa_secret", err: fmt.Errorf("ent: validator failed for field \"two_fa_secret\": %w", err)}
		}
	}
	if _, ok := au.mutation.UserID(); au.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: admin.FieldID,
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
	if value, ok := au.mutation.TwoFaSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldTwoFaSecret,
		})
	}
	if au.mutation.TwoFaSecretCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: admin.FieldTwoFaSecret,
		})
	}
	if value, ok := au.mutation.TwoFaCompleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: admin.FieldTwoFaCompleted,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldUpdatedAt,
		})
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   admin.UserTable,
			Columns: []string{admin.UserColumn},
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
			Table:   admin.UserTable,
			Columns: []string{admin.UserColumn},
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
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminMutation
}

// SetTwoFaSecret sets the "two_fa_secret" field.
func (auo *AdminUpdateOne) SetTwoFaSecret(s string) *AdminUpdateOne {
	auo.mutation.SetTwoFaSecret(s)
	return auo
}

// SetNillableTwoFaSecret sets the "two_fa_secret" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableTwoFaSecret(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetTwoFaSecret(*s)
	}
	return auo
}

// ClearTwoFaSecret clears the value of the "two_fa_secret" field.
func (auo *AdminUpdateOne) ClearTwoFaSecret() *AdminUpdateOne {
	auo.mutation.ClearTwoFaSecret()
	return auo
}

// SetTwoFaCompleted sets the "two_fa_completed" field.
func (auo *AdminUpdateOne) SetTwoFaCompleted(b bool) *AdminUpdateOne {
	auo.mutation.SetTwoFaCompleted(b)
	return auo
}

// SetNillableTwoFaCompleted sets the "two_fa_completed" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableTwoFaCompleted(b *bool) *AdminUpdateOne {
	if b != nil {
		auo.SetTwoFaCompleted(*b)
	}
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AdminUpdateOne) SetCreatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableCreatedAt(t *time.Time) *AdminUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AdminUpdateOne) SetUpdatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *AdminUpdateOne) SetUserID(id uuid.UUID) *AdminUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *AdminUpdateOne) SetUser(u *User) *AdminUpdateOne {
	return auo.SetUserID(u.ID)
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AdminUpdateOne) ClearUser() *AdminUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdminUpdateOne) Select(field string, fields ...string) *AdminUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Admin entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	var (
		err  error
		node *Admin
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
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
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AdminUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := admin.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AdminUpdateOne) check() error {
	if v, ok := auo.mutation.TwoFaSecret(); ok {
		if err := admin.TwoFaSecretValidator(v); err != nil {
			return &ValidationError{Name: "two_fa_secret", err: fmt.Errorf("ent: validator failed for field \"two_fa_secret\": %w", err)}
		}
	}
	if _, ok := auo.mutation.UserID(); auo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (_node *Admin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: admin.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Admin.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, admin.FieldID)
		for _, f := range fields {
			if !admin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != admin.FieldID {
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
	if value, ok := auo.mutation.TwoFaSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldTwoFaSecret,
		})
	}
	if auo.mutation.TwoFaSecretCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: admin.FieldTwoFaSecret,
		})
	}
	if value, ok := auo.mutation.TwoFaCompleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: admin.FieldTwoFaCompleted,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldUpdatedAt,
		})
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   admin.UserTable,
			Columns: []string{admin.UserColumn},
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
			Table:   admin.UserTable,
			Columns: []string{admin.UserColumn},
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
	_node = &Admin{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
