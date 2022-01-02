// Code generated by entc, DO NOT EDIT.

package hook

import (
	"airbound/internal/ent"
	"context"
	"fmt"
)

// The AccountFunc type is an adapter to allow the use of ordinary
// function as Account mutator.
type AccountFunc func(context.Context, *ent.AccountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AccountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountMutation", m)
	}
	return f(ctx, mv)
}

// The AddressFunc type is an adapter to allow the use of ordinary
// function as Address mutator.
type AddressFunc func(context.Context, *ent.AddressMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AddressFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AddressMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AddressMutation", m)
	}
	return f(ctx, mv)
}

// The AdminFunc type is an adapter to allow the use of ordinary
// function as Admin mutator.
type AdminFunc func(context.Context, *ent.AdminMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AdminFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AdminMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AdminMutation", m)
	}
	return f(ctx, mv)
}

// The AircraftFunc type is an adapter to allow the use of ordinary
// function as Aircraft mutator.
type AircraftFunc func(context.Context, *ent.AircraftMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AircraftFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AircraftMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AircraftMutation", m)
	}
	return f(ctx, mv)
}

// The AirlineFunc type is an adapter to allow the use of ordinary
// function as Airline mutator.
type AirlineFunc func(context.Context, *ent.AirlineMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AirlineFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AirlineMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AirlineMutation", m)
	}
	return f(ctx, mv)
}

// The AirportFunc type is an adapter to allow the use of ordinary
// function as Airport mutator.
type AirportFunc func(context.Context, *ent.AirportMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AirportFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AirportMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AirportMutation", m)
	}
	return f(ctx, mv)
}

// The CrewFunc type is an adapter to allow the use of ordinary
// function as Crew mutator.
type CrewFunc func(context.Context, *ent.CrewMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CrewFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CrewMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CrewMutation", m)
	}
	return f(ctx, mv)
}

// The CustomerFunc type is an adapter to allow the use of ordinary
// function as Customer mutator.
type CustomerFunc func(context.Context, *ent.CustomerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CustomerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CustomerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CustomerMutation", m)
	}
	return f(ctx, mv)
}

// The FlightFunc type is an adapter to allow the use of ordinary
// function as Flight mutator.
type FlightFunc func(context.Context, *ent.FlightMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FlightFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FlightMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FlightMutation", m)
	}
	return f(ctx, mv)
}

// The FlightInstanceFunc type is an adapter to allow the use of ordinary
// function as FlightInstance mutator.
type FlightInstanceFunc func(context.Context, *ent.FlightInstanceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FlightInstanceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FlightInstanceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FlightInstanceMutation", m)
	}
	return f(ctx, mv)
}

// The FlightReservationFunc type is an adapter to allow the use of ordinary
// function as FlightReservation mutator.
type FlightReservationFunc func(context.Context, *ent.FlightReservationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FlightReservationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FlightReservationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FlightReservationMutation", m)
	}
	return f(ctx, mv)
}

// The FlightScheduleFunc type is an adapter to allow the use of ordinary
// function as FlightSchedule mutator.
type FlightScheduleFunc func(context.Context, *ent.FlightScheduleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FlightScheduleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FlightScheduleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FlightScheduleMutation", m)
	}
	return f(ctx, mv)
}

// The FlightSeatFunc type is an adapter to allow the use of ordinary
// function as FlightSeat mutator.
type FlightSeatFunc func(context.Context, *ent.FlightSeatMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FlightSeatFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FlightSeatMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FlightSeatMutation", m)
	}
	return f(ctx, mv)
}

// The FrontDeskFunc type is an adapter to allow the use of ordinary
// function as FrontDesk mutator.
type FrontDeskFunc func(context.Context, *ent.FrontDeskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FrontDeskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FrontDeskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FrontDeskMutation", m)
	}
	return f(ctx, mv)
}

// The IteneraryFunc type is an adapter to allow the use of ordinary
// function as Itenerary mutator.
type IteneraryFunc func(context.Context, *ent.IteneraryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IteneraryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.IteneraryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IteneraryMutation", m)
	}
	return f(ctx, mv)
}

// The PassengerFunc type is an adapter to allow the use of ordinary
// function as Passenger mutator.
type PassengerFunc func(context.Context, *ent.PassengerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PassengerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PassengerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PassengerMutation", m)
	}
	return f(ctx, mv)
}

// The PermissionFunc type is an adapter to allow the use of ordinary
// function as Permission mutator.
type PermissionFunc func(context.Context, *ent.PermissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PermissionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PermissionMutation", m)
	}
	return f(ctx, mv)
}

// The PilotFunc type is an adapter to allow the use of ordinary
// function as Pilot mutator.
type PilotFunc func(context.Context, *ent.PilotMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PilotFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PilotMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PilotMutation", m)
	}
	return f(ctx, mv)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
	}
	return f(ctx, mv)
}

// The SeatFunc type is an adapter to allow the use of ordinary
// function as Seat mutator.
type SeatFunc func(context.Context, *ent.SeatMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SeatFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SeatMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SeatMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
