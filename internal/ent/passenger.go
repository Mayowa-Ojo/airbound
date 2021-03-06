// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/flightreservation"
	"airbound/internal/ent/flightseat"
	"airbound/internal/ent/passenger"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Passenger is the model entity for the Passenger schema.
type Passenger struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Nationality holds the value of the "nationality" field.
	Nationality string `json:"nationality,omitempty"`
	// PassportNumber holds the value of the "passport_number" field.
	PassportNumber string `json:"passport_number,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PassengerQuery when eager-loading is set.
	Edges                 PassengerEdges `json:"edges"`
	flight_reservation_id *uuid.UUID
}

// PassengerEdges holds the relations/edges for other nodes in the graph.
type PassengerEdges struct {
	// FlightReservation holds the value of the flight_reservation edge.
	FlightReservation *FlightReservation `json:"flight_reservation,omitempty"`
	// FlightSeat holds the value of the flight_seat edge.
	FlightSeat *FlightSeat `json:"flight_seat,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FlightReservationOrErr returns the FlightReservation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PassengerEdges) FlightReservationOrErr() (*FlightReservation, error) {
	if e.loadedTypes[0] {
		if e.FlightReservation == nil {
			// The edge flight_reservation was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flightreservation.Label}
		}
		return e.FlightReservation, nil
	}
	return nil, &NotLoadedError{edge: "flight_reservation"}
}

// FlightSeatOrErr returns the FlightSeat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PassengerEdges) FlightSeatOrErr() (*FlightSeat, error) {
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
func (*Passenger) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case passenger.FieldAge:
			values[i] = new(sql.NullInt64)
		case passenger.FieldFirstname, passenger.FieldLastname, passenger.FieldNationality, passenger.FieldPassportNumber:
			values[i] = new(sql.NullString)
		case passenger.FieldCreatedAt, passenger.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case passenger.FieldID:
			values[i] = new(uuid.UUID)
		case passenger.ForeignKeys[0]: // flight_reservation_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Passenger", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Passenger fields.
func (pa *Passenger) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case passenger.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pa.ID = *value
			}
		case passenger.FieldFirstname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstname", values[i])
			} else if value.Valid {
				pa.Firstname = value.String
			}
		case passenger.FieldLastname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastname", values[i])
			} else if value.Valid {
				pa.Lastname = value.String
			}
		case passenger.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pa.Age = int(value.Int64)
			}
		case passenger.FieldNationality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nationality", values[i])
			} else if value.Valid {
				pa.Nationality = value.String
			}
		case passenger.FieldPassportNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passport_number", values[i])
			} else if value.Valid {
				pa.PassportNumber = value.String
			}
		case passenger.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case passenger.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case passenger.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field flight_reservation_id", values[i])
			} else if value.Valid {
				pa.flight_reservation_id = new(uuid.UUID)
				*pa.flight_reservation_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFlightReservation queries the "flight_reservation" edge of the Passenger entity.
func (pa *Passenger) QueryFlightReservation() *FlightReservationQuery {
	return (&PassengerClient{config: pa.config}).QueryFlightReservation(pa)
}

// QueryFlightSeat queries the "flight_seat" edge of the Passenger entity.
func (pa *Passenger) QueryFlightSeat() *FlightSeatQuery {
	return (&PassengerClient{config: pa.config}).QueryFlightSeat(pa)
}

// Update returns a builder for updating this Passenger.
// Note that you need to call Passenger.Unwrap() before calling this method if this Passenger
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Passenger) Update() *PassengerUpdateOne {
	return (&PassengerClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Passenger entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Passenger) Unwrap() *Passenger {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Passenger is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Passenger) String() string {
	var builder strings.Builder
	builder.WriteString("Passenger(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", firstname=")
	builder.WriteString(pa.Firstname)
	builder.WriteString(", lastname=")
	builder.WriteString(pa.Lastname)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pa.Age))
	builder.WriteString(", nationality=")
	builder.WriteString(pa.Nationality)
	builder.WriteString(", passport_number=")
	builder.WriteString(pa.PassportNumber)
	builder.WriteString(", created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Passengers is a parsable slice of Passenger.
type Passengers []*Passenger

func (pa Passengers) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
