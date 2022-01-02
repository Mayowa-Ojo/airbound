// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/enums"
	"airbound/internal/ent/flightinstance"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// FlightInstance is the model entity for the FlightInstance schema.
type FlightInstance struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// DepartureGate holds the value of the "departure_gate" field.
	DepartureGate int `json:"departure_gate,omitempty"`
	// ArrivalGate holds the value of the "arrival_gate" field.
	ArrivalGate int `json:"arrival_gate,omitempty"`
	// FlightStatus holds the value of the "flight_status" field.
	FlightStatus enums.FlightStatus `json:"flight_status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FlightInstance) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flightinstance.FieldDepartureGate, flightinstance.FieldArrivalGate:
			values[i] = new(sql.NullInt64)
		case flightinstance.FieldFlightStatus:
			values[i] = new(sql.NullString)
		case flightinstance.FieldCreatedAt, flightinstance.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case flightinstance.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FlightInstance", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FlightInstance fields.
func (fi *FlightInstance) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flightinstance.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fi.ID = *value
			}
		case flightinstance.FieldDepartureGate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field departure_gate", values[i])
			} else if value.Valid {
				fi.DepartureGate = int(value.Int64)
			}
		case flightinstance.FieldArrivalGate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field arrival_gate", values[i])
			} else if value.Valid {
				fi.ArrivalGate = int(value.Int64)
			}
		case flightinstance.FieldFlightStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flight_status", values[i])
			} else if value.Valid {
				fi.FlightStatus = enums.FlightStatus(value.String)
			}
		case flightinstance.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fi.CreatedAt = value.Time
			}
		case flightinstance.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fi.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FlightInstance.
// Note that you need to call FlightInstance.Unwrap() before calling this method if this FlightInstance
// was returned from a transaction, and the transaction was committed or rolled back.
func (fi *FlightInstance) Update() *FlightInstanceUpdateOne {
	return (&FlightInstanceClient{config: fi.config}).UpdateOne(fi)
}

// Unwrap unwraps the FlightInstance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fi *FlightInstance) Unwrap() *FlightInstance {
	tx, ok := fi.config.driver.(*txDriver)
	if !ok {
		panic("ent: FlightInstance is not a transactional entity")
	}
	fi.config.driver = tx.drv
	return fi
}

// String implements the fmt.Stringer.
func (fi *FlightInstance) String() string {
	var builder strings.Builder
	builder.WriteString("FlightInstance(")
	builder.WriteString(fmt.Sprintf("id=%v", fi.ID))
	builder.WriteString(", departure_gate=")
	builder.WriteString(fmt.Sprintf("%v", fi.DepartureGate))
	builder.WriteString(", arrival_gate=")
	builder.WriteString(fmt.Sprintf("%v", fi.ArrivalGate))
	builder.WriteString(", flight_status=")
	builder.WriteString(fmt.Sprintf("%v", fi.FlightStatus))
	builder.WriteString(", created_at=")
	builder.WriteString(fi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(fi.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FlightInstances is a parsable slice of FlightInstance.
type FlightInstances []*FlightInstance

func (fi FlightInstances) config(cfg config) {
	for _i := range fi {
		fi[_i].config = cfg
	}
}
