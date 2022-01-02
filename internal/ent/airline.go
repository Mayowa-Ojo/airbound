// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/airline"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Airline is the model entity for the Airline schema.
type Airline struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IataCode holds the value of the "iata_code" field.
	IataCode string `json:"iata_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Airline) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case airline.FieldName, airline.FieldIataCode:
			values[i] = new(sql.NullString)
		case airline.FieldCreatedAt, airline.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case airline.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Airline", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Airline fields.
func (a *Airline) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case airline.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case airline.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case airline.FieldIataCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field iata_code", values[i])
			} else if value.Valid {
				a.IataCode = value.String
			}
		case airline.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case airline.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Airline.
// Note that you need to call Airline.Unwrap() before calling this method if this Airline
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Airline) Update() *AirlineUpdateOne {
	return (&AirlineClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Airline entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Airline) Unwrap() *Airline {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Airline is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Airline) String() string {
	var builder strings.Builder
	builder.WriteString("Airline(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", iata_code=")
	builder.WriteString(a.IataCode)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Airlines is a parsable slice of Airline.
type Airlines []*Airline

func (a Airlines) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
