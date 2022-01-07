// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/account"
	"airbound/internal/ent/address"
	"airbound/internal/ent/admin"
	"airbound/internal/ent/crew"
	"airbound/internal/ent/customer"
	"airbound/internal/ent/frontdesk"
	"airbound/internal/ent/pilot"
	"airbound/internal/ent/predicate"
	"airbound/internal/ent/role"
	"airbound/internal/ent/user"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetFirstname sets the "firstname" field.
func (uu *UserUpdate) SetFirstname(s string) *UserUpdate {
	uu.mutation.SetFirstname(s)
	return uu
}

// SetLastname sets the "lastname" field.
func (uu *UserUpdate) SetLastname(s string) *UserUpdate {
	uu.mutation.SetLastname(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (uu *UserUpdate) SetAccountID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetAccountID(id)
	return uu
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableAccountID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetAccountID(*id)
	}
	return uu
}

// SetAccount sets the "account" edge to the Account entity.
func (uu *UserUpdate) SetAccount(a *Account) *UserUpdate {
	return uu.SetAccountID(a.ID)
}

// SetAdminID sets the "admin" edge to the Admin entity by ID.
func (uu *UserUpdate) SetAdminID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetAdminID(id)
	return uu
}

// SetNillableAdminID sets the "admin" edge to the Admin entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableAdminID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetAdminID(*id)
	}
	return uu
}

// SetAdmin sets the "admin" edge to the Admin entity.
func (uu *UserUpdate) SetAdmin(a *Admin) *UserUpdate {
	return uu.SetAdminID(a.ID)
}

// SetCrewID sets the "crew" edge to the Crew entity by ID.
func (uu *UserUpdate) SetCrewID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetCrewID(id)
	return uu
}

// SetNillableCrewID sets the "crew" edge to the Crew entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableCrewID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetCrewID(*id)
	}
	return uu
}

// SetCrew sets the "crew" edge to the Crew entity.
func (uu *UserUpdate) SetCrew(c *Crew) *UserUpdate {
	return uu.SetCrewID(c.ID)
}

// SetPilotID sets the "pilot" edge to the Pilot entity by ID.
func (uu *UserUpdate) SetPilotID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetPilotID(id)
	return uu
}

// SetNillablePilotID sets the "pilot" edge to the Pilot entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillablePilotID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetPilotID(*id)
	}
	return uu
}

// SetPilot sets the "pilot" edge to the Pilot entity.
func (uu *UserUpdate) SetPilot(p *Pilot) *UserUpdate {
	return uu.SetPilotID(p.ID)
}

// SetFrontDeskID sets the "front_desk" edge to the FrontDesk entity by ID.
func (uu *UserUpdate) SetFrontDeskID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetFrontDeskID(id)
	return uu
}

// SetNillableFrontDeskID sets the "front_desk" edge to the FrontDesk entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableFrontDeskID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetFrontDeskID(*id)
	}
	return uu
}

// SetFrontDesk sets the "front_desk" edge to the FrontDesk entity.
func (uu *UserUpdate) SetFrontDesk(f *FrontDesk) *UserUpdate {
	return uu.SetFrontDeskID(f.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (uu *UserUpdate) SetCustomerID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetCustomerID(id)
	return uu
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableCustomerID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetCustomerID(*id)
	}
	return uu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (uu *UserUpdate) SetCustomer(c *Customer) *UserUpdate {
	return uu.SetCustomerID(c.ID)
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (uu *UserUpdate) SetAddressID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetAddressID(id)
	return uu
}

// SetNillableAddressID sets the "address" edge to the Address entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableAddressID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetAddressID(*id)
	}
	return uu
}

// SetAddress sets the "address" edge to the Address entity.
func (uu *UserUpdate) SetAddress(a *Address) *UserUpdate {
	return uu.SetAddressID(a.ID)
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (uu *UserUpdate) SetRoleID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetRoleID(id)
	return uu
}

// SetNillableRoleID sets the "role" edge to the Role entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableRoleID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetRoleID(*id)
	}
	return uu
}

// SetRole sets the "role" edge to the Role entity.
func (uu *UserUpdate) SetRole(r *Role) *UserUpdate {
	return uu.SetRoleID(r.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (uu *UserUpdate) ClearAccount() *UserUpdate {
	uu.mutation.ClearAccount()
	return uu
}

// ClearAdmin clears the "admin" edge to the Admin entity.
func (uu *UserUpdate) ClearAdmin() *UserUpdate {
	uu.mutation.ClearAdmin()
	return uu
}

// ClearCrew clears the "crew" edge to the Crew entity.
func (uu *UserUpdate) ClearCrew() *UserUpdate {
	uu.mutation.ClearCrew()
	return uu
}

// ClearPilot clears the "pilot" edge to the Pilot entity.
func (uu *UserUpdate) ClearPilot() *UserUpdate {
	uu.mutation.ClearPilot()
	return uu
}

// ClearFrontDesk clears the "front_desk" edge to the FrontDesk entity.
func (uu *UserUpdate) ClearFrontDesk() *UserUpdate {
	uu.mutation.ClearFrontDesk()
	return uu
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (uu *UserUpdate) ClearCustomer() *UserUpdate {
	uu.mutation.ClearCustomer()
	return uu
}

// ClearAddress clears the "address" edge to the Address entity.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.mutation.ClearAddress()
	return uu
}

// ClearRole clears the "role" edge to the Role entity.
func (uu *UserUpdate) ClearRole() *UserUpdate {
	uu.mutation.ClearRole()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Firstname(); ok {
		if err := user.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf("ent: validator failed for field \"firstname\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Lastname(); ok {
		if err := user.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf("ent: validator failed for field \"lastname\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Firstname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstname,
		})
	}
	if value, ok := uu.mutation.Lastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastname,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AccountTable,
			Columns: []string{user.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AccountTable,
			Columns: []string{user.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AdminCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AdminTable,
			Columns: []string{user.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AdminTable,
			Columns: []string{user.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CrewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CrewTable,
			Columns: []string{user.CrewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: crew.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CrewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CrewTable,
			Columns: []string{user.CrewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: crew.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PilotCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PilotTable,
			Columns: []string{user.PilotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pilot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PilotIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PilotTable,
			Columns: []string{user.PilotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pilot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FrontDeskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FrontDeskTable,
			Columns: []string{user.FrontDeskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: frontdesk.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FrontDeskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FrontDeskTable,
			Columns: []string{user.FrontDeskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: frontdesk.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CustomerTable,
			Columns: []string{user.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CustomerTable,
			Columns: []string{user.CustomerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.AddressTable,
			Columns: []string{user.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.AddressTable,
			Columns: []string{user.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetFirstname sets the "firstname" field.
func (uuo *UserUpdateOne) SetFirstname(s string) *UserUpdateOne {
	uuo.mutation.SetFirstname(s)
	return uuo
}

// SetLastname sets the "lastname" field.
func (uuo *UserUpdateOne) SetLastname(s string) *UserUpdateOne {
	uuo.mutation.SetLastname(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (uuo *UserUpdateOne) SetAccountID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetAccountID(id)
	return uuo
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAccountID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetAccountID(*id)
	}
	return uuo
}

// SetAccount sets the "account" edge to the Account entity.
func (uuo *UserUpdateOne) SetAccount(a *Account) *UserUpdateOne {
	return uuo.SetAccountID(a.ID)
}

// SetAdminID sets the "admin" edge to the Admin entity by ID.
func (uuo *UserUpdateOne) SetAdminID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetAdminID(id)
	return uuo
}

// SetNillableAdminID sets the "admin" edge to the Admin entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAdminID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetAdminID(*id)
	}
	return uuo
}

// SetAdmin sets the "admin" edge to the Admin entity.
func (uuo *UserUpdateOne) SetAdmin(a *Admin) *UserUpdateOne {
	return uuo.SetAdminID(a.ID)
}

// SetCrewID sets the "crew" edge to the Crew entity by ID.
func (uuo *UserUpdateOne) SetCrewID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetCrewID(id)
	return uuo
}

// SetNillableCrewID sets the "crew" edge to the Crew entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCrewID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCrewID(*id)
	}
	return uuo
}

// SetCrew sets the "crew" edge to the Crew entity.
func (uuo *UserUpdateOne) SetCrew(c *Crew) *UserUpdateOne {
	return uuo.SetCrewID(c.ID)
}

// SetPilotID sets the "pilot" edge to the Pilot entity by ID.
func (uuo *UserUpdateOne) SetPilotID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetPilotID(id)
	return uuo
}

// SetNillablePilotID sets the "pilot" edge to the Pilot entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePilotID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPilotID(*id)
	}
	return uuo
}

// SetPilot sets the "pilot" edge to the Pilot entity.
func (uuo *UserUpdateOne) SetPilot(p *Pilot) *UserUpdateOne {
	return uuo.SetPilotID(p.ID)
}

// SetFrontDeskID sets the "front_desk" edge to the FrontDesk entity by ID.
func (uuo *UserUpdateOne) SetFrontDeskID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetFrontDeskID(id)
	return uuo
}

// SetNillableFrontDeskID sets the "front_desk" edge to the FrontDesk entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFrontDeskID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetFrontDeskID(*id)
	}
	return uuo
}

// SetFrontDesk sets the "front_desk" edge to the FrontDesk entity.
func (uuo *UserUpdateOne) SetFrontDesk(f *FrontDesk) *UserUpdateOne {
	return uuo.SetFrontDeskID(f.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (uuo *UserUpdateOne) SetCustomerID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetCustomerID(id)
	return uuo
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCustomerID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCustomerID(*id)
	}
	return uuo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (uuo *UserUpdateOne) SetCustomer(c *Customer) *UserUpdateOne {
	return uuo.SetCustomerID(c.ID)
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (uuo *UserUpdateOne) SetAddressID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetAddressID(id)
	return uuo
}

// SetNillableAddressID sets the "address" edge to the Address entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddressID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetAddressID(*id)
	}
	return uuo
}

// SetAddress sets the "address" edge to the Address entity.
func (uuo *UserUpdateOne) SetAddress(a *Address) *UserUpdateOne {
	return uuo.SetAddressID(a.ID)
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (uuo *UserUpdateOne) SetRoleID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetRoleID(id)
	return uuo
}

// SetNillableRoleID sets the "role" edge to the Role entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRoleID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetRoleID(*id)
	}
	return uuo
}

// SetRole sets the "role" edge to the Role entity.
func (uuo *UserUpdateOne) SetRole(r *Role) *UserUpdateOne {
	return uuo.SetRoleID(r.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (uuo *UserUpdateOne) ClearAccount() *UserUpdateOne {
	uuo.mutation.ClearAccount()
	return uuo
}

// ClearAdmin clears the "admin" edge to the Admin entity.
func (uuo *UserUpdateOne) ClearAdmin() *UserUpdateOne {
	uuo.mutation.ClearAdmin()
	return uuo
}

// ClearCrew clears the "crew" edge to the Crew entity.
func (uuo *UserUpdateOne) ClearCrew() *UserUpdateOne {
	uuo.mutation.ClearCrew()
	return uuo
}

// ClearPilot clears the "pilot" edge to the Pilot entity.
func (uuo *UserUpdateOne) ClearPilot() *UserUpdateOne {
	uuo.mutation.ClearPilot()
	return uuo
}

// ClearFrontDesk clears the "front_desk" edge to the FrontDesk entity.
func (uuo *UserUpdateOne) ClearFrontDesk() *UserUpdateOne {
	uuo.mutation.ClearFrontDesk()
	return uuo
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (uuo *UserUpdateOne) ClearCustomer() *UserUpdateOne {
	uuo.mutation.ClearCustomer()
	return uuo
}

// ClearAddress clears the "address" edge to the Address entity.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.mutation.ClearAddress()
	return uuo
}

// ClearRole clears the "role" edge to the Role entity.
func (uuo *UserUpdateOne) ClearRole() *UserUpdateOne {
	uuo.mutation.ClearRole()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Firstname(); ok {
		if err := user.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf("ent: validator failed for field \"firstname\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Lastname(); ok {
		if err := user.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf("ent: validator failed for field \"lastname\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Firstname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstname,
		})
	}
	if value, ok := uuo.mutation.Lastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastname,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AccountTable,
			Columns: []string{user.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AccountTable,
			Columns: []string{user.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AdminCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AdminTable,
			Columns: []string{user.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AdminTable,
			Columns: []string{user.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CrewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CrewTable,
			Columns: []string{user.CrewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: crew.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CrewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CrewTable,
			Columns: []string{user.CrewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: crew.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PilotCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PilotTable,
			Columns: []string{user.PilotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pilot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PilotIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PilotTable,
			Columns: []string{user.PilotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pilot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FrontDeskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FrontDeskTable,
			Columns: []string{user.FrontDeskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: frontdesk.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FrontDeskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FrontDeskTable,
			Columns: []string{user.FrontDeskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: frontdesk.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CustomerTable,
			Columns: []string{user.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CustomerTable,
			Columns: []string{user.CustomerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.AddressTable,
			Columns: []string{user.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.AddressTable,
			Columns: []string{user.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
