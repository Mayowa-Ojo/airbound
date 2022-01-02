// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/enums"
	"airbound/internal/ent/flightreservation"
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
		}
	}
	return nil
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
