// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/aircraft"
	"airbound/internal/ent/enums"
	"airbound/internal/ent/flightseat"
	"airbound/internal/ent/seat"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Seat is the model entity for the Seat schema.
type Seat struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SeatNumber holds the value of the "seat_number" field.
	SeatNumber int `json:"seat_number,omitempty"`
	// SeatRow holds the value of the "seat_row" field.
	SeatRow string `json:"seat_row,omitempty"`
	// SeatType holds the value of the "seat_type" field.
	SeatType enums.SeatType `json:"seat_type,omitempty"`
	// SeatClass holds the value of the "seat_class" field.
	SeatClass enums.SeatClass `json:"seat_class,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SeatQuery when eager-loading is set.
	Edges       SeatEdges `json:"edges"`
	aircraft_id *uuid.UUID
}

// SeatEdges holds the relations/edges for other nodes in the graph.
type SeatEdges struct {
	// Aircraft holds the value of the aircraft edge.
	Aircraft *Aircraft `json:"aircraft,omitempty"`
	// FlightSeat holds the value of the flight_seat edge.
	FlightSeat *FlightSeat `json:"flight_seat,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AircraftOrErr returns the Aircraft value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SeatEdges) AircraftOrErr() (*Aircraft, error) {
	if e.loadedTypes[0] {
		if e.Aircraft == nil {
			// The edge aircraft was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: aircraft.Label}
		}
		return e.Aircraft, nil
	}
	return nil, &NotLoadedError{edge: "aircraft"}
}

// FlightSeatOrErr returns the FlightSeat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SeatEdges) FlightSeatOrErr() (*FlightSeat, error) {
	if e.loadedTypes[1] {
		if e.FlightSeat == nil {
			// The edge flight_seat was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flightseat.Label}
		}
		return e.FlightSeat, nil
	}
	return nil, &NotLoadedError{edge: "flight_seat"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Seat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case seat.FieldSeatNumber:
			values[i] = new(sql.NullInt64)
		case seat.FieldSeatRow, seat.FieldSeatType, seat.FieldSeatClass:
			values[i] = new(sql.NullString)
		case seat.FieldCreatedAt, seat.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case seat.FieldID:
			values[i] = new(uuid.UUID)
		case seat.ForeignKeys[0]: // aircraft_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Seat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Seat fields.
func (s *Seat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case seat.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case seat.FieldSeatNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field seat_number", values[i])
			} else if value.Valid {
				s.SeatNumber = int(value.Int64)
			}
		case seat.FieldSeatRow:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seat_row", values[i])
			} else if value.Valid {
				s.SeatRow = value.String
			}
		case seat.FieldSeatType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seat_type", values[i])
			} else if value.Valid {
				s.SeatType = enums.SeatType(value.String)
			}
		case seat.FieldSeatClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seat_class", values[i])
			} else if value.Valid {
				s.SeatClass = enums.SeatClass(value.String)
			}
		case seat.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case seat.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case seat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field aircraft_id", values[i])
			} else if value.Valid {
				s.aircraft_id = new(uuid.UUID)
				*s.aircraft_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryAircraft queries the "aircraft" edge of the Seat entity.
func (s *Seat) QueryAircraft() *AircraftQuery {
	return (&SeatClient{config: s.config}).QueryAircraft(s)
}

// QueryFlightSeat queries the "flight_seat" edge of the Seat entity.
func (s *Seat) QueryFlightSeat() *FlightSeatQuery {
	return (&SeatClient{config: s.config}).QueryFlightSeat(s)
}

// Update returns a builder for updating this Seat.
// Note that you need to call Seat.Unwrap() before calling this method if this Seat
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Seat) Update() *SeatUpdateOne {
	return (&SeatClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Seat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Seat) Unwrap() *Seat {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Seat is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Seat) String() string {
	var builder strings.Builder
	builder.WriteString("Seat(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", seat_number=")
	builder.WriteString(fmt.Sprintf("%v", s.SeatNumber))
	builder.WriteString(", seat_row=")
	builder.WriteString(s.SeatRow)
	builder.WriteString(", seat_type=")
	builder.WriteString(fmt.Sprintf("%v", s.SeatType))
	builder.WriteString(", seat_class=")
	builder.WriteString(fmt.Sprintf("%v", s.SeatClass))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Seats is a parsable slice of Seat.
type Seats []*Seat

func (s Seats) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
