// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/enums"
	"airbound/internal/ent/predicate"
	"airbound/internal/ent/seat"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SeatUpdate is the builder for updating Seat entities.
type SeatUpdate struct {
	config
	hooks    []Hook
	mutation *SeatMutation
}

// Where appends a list predicates to the SeatUpdate builder.
func (su *SeatUpdate) Where(ps ...predicate.Seat) *SeatUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSeatNumber sets the "seat_number" field.
func (su *SeatUpdate) SetSeatNumber(i int) *SeatUpdate {
	su.mutation.ResetSeatNumber()
	su.mutation.SetSeatNumber(i)
	return su
}

// AddSeatNumber adds i to the "seat_number" field.
func (su *SeatUpdate) AddSeatNumber(i int) *SeatUpdate {
	su.mutation.AddSeatNumber(i)
	return su
}

// SetSeatRow sets the "seat_row" field.
func (su *SeatUpdate) SetSeatRow(s string) *SeatUpdate {
	su.mutation.SetSeatRow(s)
	return su
}

// SetSeatType sets the "seat_type" field.
func (su *SeatUpdate) SetSeatType(et enums.SeatType) *SeatUpdate {
	su.mutation.SetSeatType(et)
	return su
}

// SetSeatClass sets the "seat_class" field.
func (su *SeatUpdate) SetSeatClass(ec enums.SeatClass) *SeatUpdate {
	su.mutation.SetSeatClass(ec)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SeatUpdate) SetCreatedAt(t time.Time) *SeatUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SeatUpdate) SetNillableCreatedAt(t *time.Time) *SeatUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SeatUpdate) SetUpdatedAt(t time.Time) *SeatUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// Mutation returns the SeatMutation object of the builder.
func (su *SeatUpdate) Mutation() *SeatMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeatUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeatUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeatUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SeatUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := seat.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SeatUpdate) check() error {
	if v, ok := su.mutation.SeatNumber(); ok {
		if err := seat.SeatNumberValidator(v); err != nil {
			return &ValidationError{Name: "seat_number", err: fmt.Errorf("ent: validator failed for field \"seat_number\": %w", err)}
		}
	}
	if v, ok := su.mutation.SeatRow(); ok {
		if err := seat.SeatRowValidator(v); err != nil {
			return &ValidationError{Name: "seat_row", err: fmt.Errorf("ent: validator failed for field \"seat_row\": %w", err)}
		}
	}
	if v, ok := su.mutation.SeatType(); ok {
		if err := seat.SeatTypeValidator(v); err != nil {
			return &ValidationError{Name: "seat_type", err: fmt.Errorf("ent: validator failed for field \"seat_type\": %w", err)}
		}
	}
	if v, ok := su.mutation.SeatClass(); ok {
		if err := seat.SeatClassValidator(v); err != nil {
			return &ValidationError{Name: "seat_class", err: fmt.Errorf("ent: validator failed for field \"seat_class\": %w", err)}
		}
	}
	return nil
}

func (su *SeatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seat.Table,
			Columns: seat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: seat.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.SeatNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: seat.FieldSeatNumber,
		})
	}
	if value, ok := su.mutation.AddedSeatNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: seat.FieldSeatNumber,
		})
	}
	if value, ok := su.mutation.SeatRow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seat.FieldSeatRow,
		})
	}
	if value, ok := su.mutation.SeatType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: seat.FieldSeatType,
		})
	}
	if value, ok := su.mutation.SeatClass(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: seat.FieldSeatClass,
		})
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: seat.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: seat.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SeatUpdateOne is the builder for updating a single Seat entity.
type SeatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeatMutation
}

// SetSeatNumber sets the "seat_number" field.
func (suo *SeatUpdateOne) SetSeatNumber(i int) *SeatUpdateOne {
	suo.mutation.ResetSeatNumber()
	suo.mutation.SetSeatNumber(i)
	return suo
}

// AddSeatNumber adds i to the "seat_number" field.
func (suo *SeatUpdateOne) AddSeatNumber(i int) *SeatUpdateOne {
	suo.mutation.AddSeatNumber(i)
	return suo
}

// SetSeatRow sets the "seat_row" field.
func (suo *SeatUpdateOne) SetSeatRow(s string) *SeatUpdateOne {
	suo.mutation.SetSeatRow(s)
	return suo
}

// SetSeatType sets the "seat_type" field.
func (suo *SeatUpdateOne) SetSeatType(et enums.SeatType) *SeatUpdateOne {
	suo.mutation.SetSeatType(et)
	return suo
}

// SetSeatClass sets the "seat_class" field.
func (suo *SeatUpdateOne) SetSeatClass(ec enums.SeatClass) *SeatUpdateOne {
	suo.mutation.SetSeatClass(ec)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SeatUpdateOne) SetCreatedAt(t time.Time) *SeatUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SeatUpdateOne) SetNillableCreatedAt(t *time.Time) *SeatUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SeatUpdateOne) SetUpdatedAt(t time.Time) *SeatUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// Mutation returns the SeatMutation object of the builder.
func (suo *SeatUpdateOne) Mutation() *SeatMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeatUpdateOne) Select(field string, fields ...string) *SeatUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Seat entity.
func (suo *SeatUpdateOne) Save(ctx context.Context) (*Seat, error) {
	var (
		err  error
		node *Seat
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeatUpdateOne) SaveX(ctx context.Context) *Seat {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeatUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeatUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SeatUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := seat.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SeatUpdateOne) check() error {
	if v, ok := suo.mutation.SeatNumber(); ok {
		if err := seat.SeatNumberValidator(v); err != nil {
			return &ValidationError{Name: "seat_number", err: fmt.Errorf("ent: validator failed for field \"seat_number\": %w", err)}
		}
	}
	if v, ok := suo.mutation.SeatRow(); ok {
		if err := seat.SeatRowValidator(v); err != nil {
			return &ValidationError{Name: "seat_row", err: fmt.Errorf("ent: validator failed for field \"seat_row\": %w", err)}
		}
	}
	if v, ok := suo.mutation.SeatType(); ok {
		if err := seat.SeatTypeValidator(v); err != nil {
			return &ValidationError{Name: "seat_type", err: fmt.Errorf("ent: validator failed for field \"seat_type\": %w", err)}
		}
	}
	if v, ok := suo.mutation.SeatClass(); ok {
		if err := seat.SeatClassValidator(v); err != nil {
			return &ValidationError{Name: "seat_class", err: fmt.Errorf("ent: validator failed for field \"seat_class\": %w", err)}
		}
	}
	return nil
}

func (suo *SeatUpdateOne) sqlSave(ctx context.Context) (_node *Seat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seat.Table,
			Columns: seat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: seat.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Seat.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seat.FieldID)
		for _, f := range fields {
			if !seat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != seat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.SeatNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: seat.FieldSeatNumber,
		})
	}
	if value, ok := suo.mutation.AddedSeatNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: seat.FieldSeatNumber,
		})
	}
	if value, ok := suo.mutation.SeatRow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: seat.FieldSeatRow,
		})
	}
	if value, ok := suo.mutation.SeatType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: seat.FieldSeatType,
		})
	}
	if value, ok := suo.mutation.SeatClass(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: seat.FieldSeatClass,
		})
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: seat.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: seat.FieldUpdatedAt,
		})
	}
	_node = &Seat{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
