// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/airline"
	"airbound/internal/ent/pilot"
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

// PilotUpdate is the builder for updating Pilot entities.
type PilotUpdate struct {
	config
	hooks    []Hook
	mutation *PilotMutation
}

// Where appends a list predicates to the PilotUpdate builder.
func (pu *PilotUpdate) Where(ps ...predicate.Pilot) *PilotUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetEmployeeID sets the "employee_id" field.
func (pu *PilotUpdate) SetEmployeeID(s string) *PilotUpdate {
	pu.mutation.SetEmployeeID(s)
	return pu
}

// SetLicenceNumber sets the "licence_number" field.
func (pu *PilotUpdate) SetLicenceNumber(s string) *PilotUpdate {
	pu.mutation.SetLicenceNumber(s)
	return pu
}

// SetFlightHours sets the "flight_hours" field.
func (pu *PilotUpdate) SetFlightHours(i int) *PilotUpdate {
	pu.mutation.ResetFlightHours()
	pu.mutation.SetFlightHours(i)
	return pu
}

// AddFlightHours adds i to the "flight_hours" field.
func (pu *PilotUpdate) AddFlightHours(i int) *PilotUpdate {
	pu.mutation.AddFlightHours(i)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PilotUpdate) SetCreatedAt(t time.Time) *PilotUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PilotUpdate) SetNillableCreatedAt(t *time.Time) *PilotUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PilotUpdate) SetUpdatedAt(t time.Time) *PilotUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PilotUpdate) SetUserID(id uuid.UUID) *PilotUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PilotUpdate) SetUser(u *User) *PilotUpdate {
	return pu.SetUserID(u.ID)
}

// SetAirlineID sets the "airline" edge to the Airline entity by ID.
func (pu *PilotUpdate) SetAirlineID(id uuid.UUID) *PilotUpdate {
	pu.mutation.SetAirlineID(id)
	return pu
}

// SetNillableAirlineID sets the "airline" edge to the Airline entity by ID if the given value is not nil.
func (pu *PilotUpdate) SetNillableAirlineID(id *uuid.UUID) *PilotUpdate {
	if id != nil {
		pu = pu.SetAirlineID(*id)
	}
	return pu
}

// SetAirline sets the "airline" edge to the Airline entity.
func (pu *PilotUpdate) SetAirline(a *Airline) *PilotUpdate {
	return pu.SetAirlineID(a.ID)
}

// Mutation returns the PilotMutation object of the builder.
func (pu *PilotUpdate) Mutation() *PilotMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PilotUpdate) ClearUser() *PilotUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearAirline clears the "airline" edge to the Airline entity.
func (pu *PilotUpdate) ClearAirline() *PilotUpdate {
	pu.mutation.ClearAirline()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PilotUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PilotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PilotUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PilotUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PilotUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PilotUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := pilot.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PilotUpdate) check() error {
	if v, ok := pu.mutation.EmployeeID(); ok {
		if err := pilot.EmployeeIDValidator(v); err != nil {
			return &ValidationError{Name: "employee_id", err: fmt.Errorf("ent: validator failed for field \"employee_id\": %w", err)}
		}
	}
	if v, ok := pu.mutation.LicenceNumber(); ok {
		if err := pilot.LicenceNumberValidator(v); err != nil {
			return &ValidationError{Name: "licence_number", err: fmt.Errorf("ent: validator failed for field \"licence_number\": %w", err)}
		}
	}
	if v, ok := pu.mutation.FlightHours(); ok {
		if err := pilot.FlightHoursValidator(v); err != nil {
			return &ValidationError{Name: "flight_hours", err: fmt.Errorf("ent: validator failed for field \"flight_hours\": %w", err)}
		}
	}
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (pu *PilotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pilot.Table,
			Columns: pilot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pilot.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.EmployeeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pilot.FieldEmployeeID,
		})
	}
	if value, ok := pu.mutation.LicenceNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pilot.FieldLicenceNumber,
		})
	}
	if value, ok := pu.mutation.FlightHours(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pilot.FieldFlightHours,
		})
	}
	if value, ok := pu.mutation.AddedFlightHours(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pilot.FieldFlightHours,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pilot.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pilot.FieldUpdatedAt,
		})
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pilot.UserTable,
			Columns: []string{pilot.UserColumn},
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
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pilot.UserTable,
			Columns: []string{pilot.UserColumn},
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
	if pu.mutation.AirlineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pilot.AirlineTable,
			Columns: []string{pilot.AirlineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: airline.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AirlineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pilot.AirlineTable,
			Columns: []string{pilot.AirlineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: airline.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pilot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PilotUpdateOne is the builder for updating a single Pilot entity.
type PilotUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PilotMutation
}

// SetEmployeeID sets the "employee_id" field.
func (puo *PilotUpdateOne) SetEmployeeID(s string) *PilotUpdateOne {
	puo.mutation.SetEmployeeID(s)
	return puo
}

// SetLicenceNumber sets the "licence_number" field.
func (puo *PilotUpdateOne) SetLicenceNumber(s string) *PilotUpdateOne {
	puo.mutation.SetLicenceNumber(s)
	return puo
}

// SetFlightHours sets the "flight_hours" field.
func (puo *PilotUpdateOne) SetFlightHours(i int) *PilotUpdateOne {
	puo.mutation.ResetFlightHours()
	puo.mutation.SetFlightHours(i)
	return puo
}

// AddFlightHours adds i to the "flight_hours" field.
func (puo *PilotUpdateOne) AddFlightHours(i int) *PilotUpdateOne {
	puo.mutation.AddFlightHours(i)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PilotUpdateOne) SetCreatedAt(t time.Time) *PilotUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PilotUpdateOne) SetNillableCreatedAt(t *time.Time) *PilotUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PilotUpdateOne) SetUpdatedAt(t time.Time) *PilotUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PilotUpdateOne) SetUserID(id uuid.UUID) *PilotUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PilotUpdateOne) SetUser(u *User) *PilotUpdateOne {
	return puo.SetUserID(u.ID)
}

// SetAirlineID sets the "airline" edge to the Airline entity by ID.
func (puo *PilotUpdateOne) SetAirlineID(id uuid.UUID) *PilotUpdateOne {
	puo.mutation.SetAirlineID(id)
	return puo
}

// SetNillableAirlineID sets the "airline" edge to the Airline entity by ID if the given value is not nil.
func (puo *PilotUpdateOne) SetNillableAirlineID(id *uuid.UUID) *PilotUpdateOne {
	if id != nil {
		puo = puo.SetAirlineID(*id)
	}
	return puo
}

// SetAirline sets the "airline" edge to the Airline entity.
func (puo *PilotUpdateOne) SetAirline(a *Airline) *PilotUpdateOne {
	return puo.SetAirlineID(a.ID)
}

// Mutation returns the PilotMutation object of the builder.
func (puo *PilotUpdateOne) Mutation() *PilotMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PilotUpdateOne) ClearUser() *PilotUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearAirline clears the "airline" edge to the Airline entity.
func (puo *PilotUpdateOne) ClearAirline() *PilotUpdateOne {
	puo.mutation.ClearAirline()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PilotUpdateOne) Select(field string, fields ...string) *PilotUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pilot entity.
func (puo *PilotUpdateOne) Save(ctx context.Context) (*Pilot, error) {
	var (
		err  error
		node *Pilot
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PilotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PilotUpdateOne) SaveX(ctx context.Context) *Pilot {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PilotUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PilotUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PilotUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := pilot.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PilotUpdateOne) check() error {
	if v, ok := puo.mutation.EmployeeID(); ok {
		if err := pilot.EmployeeIDValidator(v); err != nil {
			return &ValidationError{Name: "employee_id", err: fmt.Errorf("ent: validator failed for field \"employee_id\": %w", err)}
		}
	}
	if v, ok := puo.mutation.LicenceNumber(); ok {
		if err := pilot.LicenceNumberValidator(v); err != nil {
			return &ValidationError{Name: "licence_number", err: fmt.Errorf("ent: validator failed for field \"licence_number\": %w", err)}
		}
	}
	if v, ok := puo.mutation.FlightHours(); ok {
		if err := pilot.FlightHoursValidator(v); err != nil {
			return &ValidationError{Name: "flight_hours", err: fmt.Errorf("ent: validator failed for field \"flight_hours\": %w", err)}
		}
	}
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (puo *PilotUpdateOne) sqlSave(ctx context.Context) (_node *Pilot, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pilot.Table,
			Columns: pilot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pilot.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Pilot.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pilot.FieldID)
		for _, f := range fields {
			if !pilot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pilot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.EmployeeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pilot.FieldEmployeeID,
		})
	}
	if value, ok := puo.mutation.LicenceNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pilot.FieldLicenceNumber,
		})
	}
	if value, ok := puo.mutation.FlightHours(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pilot.FieldFlightHours,
		})
	}
	if value, ok := puo.mutation.AddedFlightHours(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pilot.FieldFlightHours,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pilot.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pilot.FieldUpdatedAt,
		})
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pilot.UserTable,
			Columns: []string{pilot.UserColumn},
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
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pilot.UserTable,
			Columns: []string{pilot.UserColumn},
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
	if puo.mutation.AirlineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pilot.AirlineTable,
			Columns: []string{pilot.AirlineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: airline.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AirlineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pilot.AirlineTable,
			Columns: []string{pilot.AirlineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: airline.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Pilot{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pilot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}