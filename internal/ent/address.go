// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/address"
	"airbound/internal/ent/airport"
	"airbound/internal/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Address is the model entity for the Address schema.
type Address struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Street holds the value of the "street" field.
	Street string `json:"street,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Zipcode holds the value of the "zipcode" field.
	Zipcode string `json:"zipcode,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AddressQuery when eager-loading is set.
	Edges           AddressEdges `json:"edges"`
	airport_address *uuid.UUID
	user_address    *uuid.UUID
}

// AddressEdges holds the relations/edges for other nodes in the graph.
type AddressEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Airport holds the value of the airport edge.
	Airport *Airport `json:"airport,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AddressEdges) UserOrErr() (*User, error) {
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

// AirportOrErr returns the Airport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AddressEdges) AirportOrErr() (*Airport, error) {
	if e.loadedTypes[1] {
		if e.Airport == nil {
			// The edge airport was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.Airport, nil
	}
	return nil, &NotLoadedError{edge: "airport"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Address) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case address.FieldStreet, address.FieldCity, address.FieldState, address.FieldZipcode:
			values[i] = new(sql.NullString)
		case address.FieldCreatedAt, address.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case address.FieldID:
			values[i] = new(uuid.UUID)
		case address.ForeignKeys[0]: // airport_address
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case address.ForeignKeys[1]: // user_address
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Address", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Address fields.
func (a *Address) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case address.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case address.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field street", values[i])
			} else if value.Valid {
				a.Street = value.String
			}
		case address.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				a.City = value.String
			}
		case address.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				a.State = value.String
			}
		case address.FieldZipcode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zipcode", values[i])
			} else if value.Valid {
				a.Zipcode = value.String
			}
		case address.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case address.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case address.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field airport_address", values[i])
			} else if value.Valid {
				a.airport_address = new(uuid.UUID)
				*a.airport_address = *value.S.(*uuid.UUID)
			}
		case address.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_address", values[i])
			} else if value.Valid {
				a.user_address = new(uuid.UUID)
				*a.user_address = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Address entity.
func (a *Address) QueryUser() *UserQuery {
	return (&AddressClient{config: a.config}).QueryUser(a)
}

// QueryAirport queries the "airport" edge of the Address entity.
func (a *Address) QueryAirport() *AirportQuery {
	return (&AddressClient{config: a.config}).QueryAirport(a)
}

// Update returns a builder for updating this Address.
// Note that you need to call Address.Unwrap() before calling this method if this Address
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Address) Update() *AddressUpdateOne {
	return (&AddressClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Address entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Address) Unwrap() *Address {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Address is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Address) String() string {
	var builder strings.Builder
	builder.WriteString("Address(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", street=")
	builder.WriteString(a.Street)
	builder.WriteString(", city=")
	builder.WriteString(a.City)
	builder.WriteString(", state=")
	builder.WriteString(a.State)
	builder.WriteString(", zipcode=")
	builder.WriteString(a.Zipcode)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Addresses is a parsable slice of Address.
type Addresses []*Address

func (a Addresses) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
