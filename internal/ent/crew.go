// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/airline"
	"airbound/internal/ent/crew"
	"airbound/internal/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Crew is the model entity for the Crew schema.
type Crew struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// EmployeeID holds the value of the "employee_id" field.
	EmployeeID string `json:"employee_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CrewQuery when eager-loading is set.
	Edges      CrewEdges `json:"edges"`
	airline_id *uuid.UUID
	user_id    *uuid.UUID
}

// CrewEdges holds the relations/edges for other nodes in the graph.
type CrewEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Airline holds the value of the airline edge.
	Airline *Airline `json:"airline,omitempty"`
	// Flights holds the value of the flights edge.
	Flights []*Flight `json:"flights,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CrewEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// AirlineOrErr returns the Airline value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CrewEdges) AirlineOrErr() (*Airline, error) {
	if e.loadedTypes[1] {
		if e.Airline == nil {
			// The edge airline was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: airline.Label}
		}
		return e.Airline, nil
	}
	return nil, &NotLoadedError{edge: "airline"}
}

// FlightsOrErr returns the Flights value or an error if the edge
// was not loaded in eager-loading.
func (e CrewEdges) FlightsOrErr() ([]*Flight, error) {
	if e.loadedTypes[2] {
		return e.Flights, nil
	}
	return nil, &NotLoadedError{edge: "flights"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Crew) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case crew.FieldEmployeeID:
			values[i] = new(sql.NullString)
		case crew.FieldCreatedAt, crew.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case crew.FieldID:
			values[i] = new(uuid.UUID)
		case crew.ForeignKeys[0]: // airline_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case crew.ForeignKeys[1]: // user_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Crew", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Crew fields.
func (c *Crew) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case crew.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case crew.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field employee_id", values[i])
			} else if value.Valid {
				c.EmployeeID = value.String
			}
		case crew.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case crew.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case crew.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field airline_id", values[i])
			} else if value.Valid {
				c.airline_id = new(uuid.UUID)
				*c.airline_id = *value.S.(*uuid.UUID)
			}
		case crew.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.user_id = new(uuid.UUID)
				*c.user_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Crew entity.
func (c *Crew) QueryUser() *UserQuery {
	return (&CrewClient{config: c.config}).QueryUser(c)
}

// QueryAirline queries the "airline" edge of the Crew entity.
func (c *Crew) QueryAirline() *AirlineQuery {
	return (&CrewClient{config: c.config}).QueryAirline(c)
}

// QueryFlights queries the "flights" edge of the Crew entity.
func (c *Crew) QueryFlights() *FlightQuery {
	return (&CrewClient{config: c.config}).QueryFlights(c)
}

// Update returns a builder for updating this Crew.
// Note that you need to call Crew.Unwrap() before calling this method if this Crew
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Crew) Update() *CrewUpdateOne {
	return (&CrewClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Crew entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Crew) Unwrap() *Crew {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Crew is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Crew) String() string {
	var builder strings.Builder
	builder.WriteString("Crew(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", employee_id=")
	builder.WriteString(c.EmployeeID)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Crews is a parsable slice of Crew.
type Crews []*Crew

func (c Crews) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
