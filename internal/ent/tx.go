// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Address is the client for interacting with the Address builders.
	Address *AddressClient
	// Admin is the client for interacting with the Admin builders.
	Admin *AdminClient
	// Aircraft is the client for interacting with the Aircraft builders.
	Aircraft *AircraftClient
	// Airline is the client for interacting with the Airline builders.
	Airline *AirlineClient
	// Airport is the client for interacting with the Airport builders.
	Airport *AirportClient
	// Crew is the client for interacting with the Crew builders.
	Crew *CrewClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Flight is the client for interacting with the Flight builders.
	Flight *FlightClient
	// FlightInstance is the client for interacting with the FlightInstance builders.
	FlightInstance *FlightInstanceClient
	// FlightReservation is the client for interacting with the FlightReservation builders.
	FlightReservation *FlightReservationClient
	// FlightSchedule is the client for interacting with the FlightSchedule builders.
	FlightSchedule *FlightScheduleClient
	// FlightSeat is the client for interacting with the FlightSeat builders.
	FlightSeat *FlightSeatClient
	// FrontDesk is the client for interacting with the FrontDesk builders.
	FrontDesk *FrontDeskClient
	// Itenerary is the client for interacting with the Itenerary builders.
	Itenerary *IteneraryClient
	// Passenger is the client for interacting with the Passenger builders.
	Passenger *PassengerClient
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Pilot is the client for interacting with the Pilot builders.
	Pilot *PilotClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Seat is the client for interacting with the Seat builders.
	Seat *SeatClient
	// User is the client for interacting with the User builders.
	User *UserClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Account = NewAccountClient(tx.config)
	tx.Address = NewAddressClient(tx.config)
	tx.Admin = NewAdminClient(tx.config)
	tx.Aircraft = NewAircraftClient(tx.config)
	tx.Airline = NewAirlineClient(tx.config)
	tx.Airport = NewAirportClient(tx.config)
	tx.Crew = NewCrewClient(tx.config)
	tx.Customer = NewCustomerClient(tx.config)
	tx.Flight = NewFlightClient(tx.config)
	tx.FlightInstance = NewFlightInstanceClient(tx.config)
	tx.FlightReservation = NewFlightReservationClient(tx.config)
	tx.FlightSchedule = NewFlightScheduleClient(tx.config)
	tx.FlightSeat = NewFlightSeatClient(tx.config)
	tx.FrontDesk = NewFrontDeskClient(tx.config)
	tx.Itenerary = NewIteneraryClient(tx.config)
	tx.Passenger = NewPassengerClient(tx.config)
	tx.Permission = NewPermissionClient(tx.config)
	tx.Pilot = NewPilotClient(tx.config)
	tx.Role = NewRoleClient(tx.config)
	tx.Seat = NewSeatClient(tx.config)
	tx.User = NewUserClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Account.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
