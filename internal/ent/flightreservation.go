// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/enums"
	"airbound/internal/ent/flightinstance"
	"airbound/internal/ent/flightreservation"
	"airbound/internal/ent/itenerary"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// FlightReservation is the model entity for the FlightReservation schema.
type FlightReservation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ReservationNumber holds the value of the "reservation_number" field.
	ReservationNumber string `json:"reservation_number,omitempty"`
	// ReservationStatus holds the value of the "reservation_status" field.
	ReservationStatus enums.ReservationStatus `json:"reservation_status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlightReservationQuery when eager-loading is set.
	Edges              FlightReservationEdges `json:"edges"`
	flight_instance_id *uuid.UUID
	itenerary_id       *uuid.UUID
}

// FlightReservationEdges holds the relations/edges for other nodes in the graph.
type FlightReservationEdges struct {
	// FlightInstance holds the value of the flight_instance edge.
	FlightInstance *FlightInstance `json:"flight_instance,omitempty"`
	// Itenerary holds the value of the itenerary edge.
	Itenerary *Itenerary `json:"itenerary,omitempty"`
	// Passengers holds the value of the passengers edge.
	Passengers []*Passenger `json:"passengers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FlightInstanceOrErr returns the FlightInstance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightReservationEdges) FlightInstanceOrErr() (*FlightInstance, error) {
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

// IteneraryOrErr returns the Itenerary value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightReservationEdges) IteneraryOrErr() (*Itenerary, error) {
	if e.loadedTypes[1] {
		if e.Itenerary == nil {
			// The edge itenerary was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: itenerary.Label}
		}
		return e.Itenerary, nil
	}
	return nil, &NotLoadedError{edge: "itenerary"}
}

// PassengersOrErr returns the Passengers value or an error if the edge
// was not loaded in eager-loading.
func (e FlightReservationEdges) PassengersOrErr() ([]*Passenger, error) {
	if e.loadedTypes[2] {
		return e.Passengers, nil
	}
	return nil, &NotLoadedError{edge: "passengers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlightReservation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flightreservation.FieldReservationNumber, flightreservation.FieldReservationStatus:
			values[i] = new(sql.NullString)
		case flightreservation.FieldCreatedAt, flightreservation.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case flightreservation.FieldID:
			values[i] = new(uuid.UUID)
		case flightreservation.ForeignKeys[0]: // flight_instance_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case flightreservation.ForeignKeys[1]: // itenerary_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlightReservation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlightReservation fields.
func (fr *FlightReservation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flightreservation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fr.ID = *value
			}
		case flightreservation.FieldReservationNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reservation_number", values[i])
			} else if value.Valid {
				fr.ReservationNumber = value.String
			}
		case flightreservation.FieldReservationStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reservation_status", values[i])
			} else if value.Valid {
				fr.ReservationStatus = enums.ReservationStatus(value.String)
			}
		case flightreservation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fr.CreatedAt = value.Time
			}
		case flightreservation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fr.UpdatedAt = value.Time
			}
		case flightreservation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field flight_instance_id", values[i])
			} else if value.Valid {
				fr.flight_instance_id = new(uuid.UUID)
				*fr.flight_instance_id = *value.S.(*uuid.UUID)
			}
		case flightreservation.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field itenerary_id", values[i])
			} else if value.Valid {
				fr.itenerary_id = new(uuid.UUID)
				*fr.itenerary_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFlightInstance queries the "flight_instance" edge of the FlightReservation entity.
func (fr *FlightReservation) QueryFlightInstance() *FlightInstanceQuery {
	return (&FlightReservationClient{config: fr.config}).QueryFlightInstance(fr)
}

// QueryItenerary queries the "itenerary" edge of the FlightReservation entity.
func (fr *FlightReservation) QueryItenerary() *IteneraryQuery {
	return (&FlightReservationClient{config: fr.config}).QueryItenerary(fr)
}

// QueryPassengers queries the "passengers" edge of the FlightReservation entity.
func (fr *FlightReservation) QueryPassengers() *PassengerQuery {
	return (&FlightReservationClient{config: fr.config}).QueryPassengers(fr)
}

// Update returns a builder for updating this FlightReservation.
// Note that you need to call FlightReservation.Unwrap() before calling this method if this FlightReservation
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FlightReservation) Update() *FlightReservationUpdateOne {
	return (&FlightReservationClient{config: fr.config}).UpdateOne(fr)
}

// Unwrap unwraps the FlightReservation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *FlightReservation) Unwrap() *FlightReservation {
	tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("ent: FlightReservation is not a transactional entity")
	}
	fr.config.driver = tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FlightReservation) String() string {
	var builder strings.Builder
	builder.WriteString("FlightReservation(")
	builder.WriteString(fmt.Sprintf("id=%v", fr.ID))
	builder.WriteString(", reservation_number=")
	builder.WriteString(fr.ReservationNumber)
	builder.WriteString(", reservation_status=")
	builder.WriteString(fmt.Sprintf("%v", fr.ReservationStatus))
	builder.WriteString(", created_at=")
	builder.WriteString(fr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(fr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FlightReservations is a parsable slice of FlightReservation.
type FlightReservations []*FlightReservation

func (fr FlightReservations) config(cfg config) {
	for _i := range fr {
		fr[_i].config = cfg
	}
}