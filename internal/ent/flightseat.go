// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/flightinstance"
	"airbound/internal/ent/flightseat"
	"airbound/internal/ent/passenger"
	"airbound/internal/ent/seat"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// FlightSeat is the model entity for the FlightSeat schema.
type FlightSeat struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Fare holds the value of the "fare" field.
	Fare float64 `json:"fare,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlightSeatQuery when eager-loading is set.
	Edges                 FlightSeatEdges `json:"edges"`
	flight_instance_id    *uuid.UUID
	passenger_flight_seat *uuid.UUID
	seat_flight_seat      *uuid.UUID
}

// FlightSeatEdges holds the relations/edges for other nodes in the graph.
type FlightSeatEdges struct {
	// FlightInstance holds the value of the flight_instance edge.
	FlightInstance *FlightInstance `json:"flight_instance,omitempty"`
	// Seat holds the value of the seat edge.
	Seat *Seat `json:"seat,omitempty"`
	// Passenger holds the value of the passenger edge.
	Passenger *Passenger `json:"passenger,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FlightInstanceOrErr returns the FlightInstance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightSeatEdges) FlightInstanceOrErr() (*FlightInstance, error) {
	if e.loadedTypes[0] {
		if e.FlightInstance == nil {
			// The edge flight_instance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flightinstance.Label}
		}
		return e.FlightInstance, nil
	}
	return nil, &NotLoadedError{edge: "flight_instance"}
}

// SeatOrErr returns the Seat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightSeatEdges) SeatOrErr() (*Seat, error) {
	if e.loadedTypes[1] {
		if e.Seat == nil {
			// The edge seat was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: seat.Label}
		}
		return e.Seat, nil
	}
	return nil, &NotLoadedError{edge: "seat"}
}

// PassengerOrErr returns the Passenger value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightSeatEdges) PassengerOrErr() (*Passenger, error) {
	if e.loadedTypes[2] {
		if e.Passenger == nil {
			// The edge passenger was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: passenger.Label}
		}
		return e.Passenger, nil
	}
	return nil, &NotLoadedError{edge: "passenger"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlightSeat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flightseat.FieldFare:
			values[i] = new(sql.NullFloat64)
		case flightseat.FieldCreatedAt, flightseat.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case flightseat.FieldID:
			values[i] = new(uuid.UUID)
		case flightseat.ForeignKeys[0]: // flight_instance_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case flightseat.ForeignKeys[1]: // passenger_flight_seat
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case flightseat.ForeignKeys[2]: // seat_flight_seat
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlightSeat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlightSeat fields.
func (fs *FlightSeat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flightseat.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fs.ID = *value
			}
		case flightseat.FieldFare:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fare", values[i])
			} else if value.Valid {
				fs.Fare = value.Float64
			}
		case flightseat.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fs.CreatedAt = value.Time
			}
		case flightseat.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fs.UpdatedAt = value.Time
			}
		case flightseat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field flight_instance_id", values[i])
			} else if value.Valid {
				fs.flight_instance_id = new(uuid.UUID)
				*fs.flight_instance_id = *value.S.(*uuid.UUID)
			}
		case flightseat.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field passenger_flight_seat", values[i])
			} else if value.Valid {
				fs.passenger_flight_seat = new(uuid.UUID)
				*fs.passenger_flight_seat = *value.S.(*uuid.UUID)
			}
		case flightseat.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field seat_flight_seat", values[i])
			} else if value.Valid {
				fs.seat_flight_seat = new(uuid.UUID)
				*fs.seat_flight_seat = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFlightInstance queries the "flight_instance" edge of the FlightSeat entity.
func (fs *FlightSeat) QueryFlightInstance() *FlightInstanceQuery {
	return (&FlightSeatClient{config: fs.config}).QueryFlightInstance(fs)
}

// QuerySeat queries the "seat" edge of the FlightSeat entity.
func (fs *FlightSeat) QuerySeat() *SeatQuery {
	return (&FlightSeatClient{config: fs.config}).QuerySeat(fs)
}

// QueryPassenger queries the "passenger" edge of the FlightSeat entity.
func (fs *FlightSeat) QueryPassenger() *PassengerQuery {
	return (&FlightSeatClient{config: fs.config}).QueryPassenger(fs)
}

// Update returns a builder for updating this FlightSeat.
// Note that you need to call FlightSeat.Unwrap() before calling this method if this FlightSeat
// was returned from a transaction, and the transaction was committed or rolled back.
func (fs *FlightSeat) Update() *FlightSeatUpdateOne {
	return (&FlightSeatClient{config: fs.config}).UpdateOne(fs)
}

// Unwrap unwraps the FlightSeat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fs *FlightSeat) Unwrap() *FlightSeat {
	tx, ok := fs.config.driver.(*txDriver)
	if !ok {
		panic("ent: FlightSeat is not a transactional entity")
	}
	fs.config.driver = tx.drv
	return fs
}

// String implements the fmt.Stringer.
func (fs *FlightSeat) String() string {
	var builder strings.Builder
	builder.WriteString("FlightSeat(")
	builder.WriteString(fmt.Sprintf("id=%v", fs.ID))
	builder.WriteString(", fare=")
	builder.WriteString(fmt.Sprintf("%v", fs.Fare))
	builder.WriteString(", created_at=")
	builder.WriteString(fs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(fs.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FlightSeats is a parsable slice of FlightSeat.
type FlightSeats []*FlightSeat

func (fs FlightSeats) config(cfg config) {
	for _i := range fs {
		fs[_i].config = cfg
	}
}
