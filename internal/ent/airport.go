// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/address"
	"airbound/internal/ent/airport"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Airport is the model entity for the Airport schema.
type Airport struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IataCode holds the value of the "iata_code" field.
	IataCode string `json:"iata_code,omitempty"`
	// IcaoCode holds the value of the "icao_code" field.
	IcaoCode string `json:"icao_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AirportQuery when eager-loading is set.
	Edges      AirportEdges `json:"edges"`
	address_id *uuid.UUID
}

// AirportEdges holds the relations/edges for other nodes in the graph.
type AirportEdges struct {
	// Address holds the value of the address edge.
	Address *Address `json:"address,omitempty"`
	// FrontDesks holds the value of the front_desks edge.
	FrontDesks []*FrontDesk `json:"front_desks,omitempty"`
	// DepartureFlights holds the value of the departure_flights edge.
	DepartureFlights []*Flight `json:"departure_flights,omitempty"`
	// ArrivalFlights holds the value of the arrival_flights edge.
	ArrivalFlights []*Flight `json:"arrival_flights,omitempty"`
	// OriginIteneraries holds the value of the origin_iteneraries edge.
	OriginIteneraries []*Itenerary `json:"origin_iteneraries,omitempty"`
	// DestinationIteneraries holds the value of the destination_iteneraries edge.
	DestinationIteneraries []*Itenerary `json:"destination_iteneraries,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// AddressOrErr returns the Address value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AirportEdges) AddressOrErr() (*Address, error) {
	if e.loadedTypes[0] {
		if e.Address == nil {
			// The edge address was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: address.Label}
		}
		return e.Address, nil
	}
	return nil, &NotLoadedError{edge: "address"}
}

// FrontDesksOrErr returns the FrontDesks value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) FrontDesksOrErr() ([]*FrontDesk, error) {
	if e.loadedTypes[1] {
		return e.FrontDesks, nil
	}
	return nil, &NotLoadedError{edge: "front_desks"}
}

// DepartureFlightsOrErr returns the DepartureFlights value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) DepartureFlightsOrErr() ([]*Flight, error) {
	if e.loadedTypes[2] {
		return e.DepartureFlights, nil
	}
	return nil, &NotLoadedError{edge: "departure_flights"}
}

// ArrivalFlightsOrErr returns the ArrivalFlights value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) ArrivalFlightsOrErr() ([]*Flight, error) {
	if e.loadedTypes[3] {
		return e.ArrivalFlights, nil
	}
	return nil, &NotLoadedError{edge: "arrival_flights"}
}

// OriginItenerariesOrErr returns the OriginIteneraries value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) OriginItenerariesOrErr() ([]*Itenerary, error) {
	if e.loadedTypes[4] {
		return e.OriginIteneraries, nil
	}
	return nil, &NotLoadedError{edge: "origin_iteneraries"}
}

// DestinationItenerariesOrErr returns the DestinationIteneraries value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) DestinationItenerariesOrErr() ([]*Itenerary, error) {
	if e.loadedTypes[5] {
		return e.DestinationIteneraries, nil
	}
	return nil, &NotLoadedError{edge: "destination_iteneraries"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Airport) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case airport.FieldName, airport.FieldIataCode, airport.FieldIcaoCode:
			values[i] = new(sql.NullString)
		case airport.FieldCreatedAt, airport.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case airport.FieldID:
			values[i] = new(uuid.UUID)
		case airport.ForeignKeys[0]: // address_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Airport", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Airport fields.
func (a *Airport) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case airport.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case airport.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case airport.FieldIataCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field iata_code", values[i])
			} else if value.Valid {
				a.IataCode = value.String
			}
		case airport.FieldIcaoCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icao_code", values[i])
			} else if value.Valid {
				a.IcaoCode = value.String
			}
		case airport.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case airport.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case airport.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field address_id", values[i])
			} else if value.Valid {
				a.address_id = new(uuid.UUID)
				*a.address_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryAddress queries the "address" edge of the Airport entity.
func (a *Airport) QueryAddress() *AddressQuery {
	return (&AirportClient{config: a.config}).QueryAddress(a)
}

// QueryFrontDesks queries the "front_desks" edge of the Airport entity.
func (a *Airport) QueryFrontDesks() *FrontDeskQuery {
	return (&AirportClient{config: a.config}).QueryFrontDesks(a)
}

// QueryDepartureFlights queries the "departure_flights" edge of the Airport entity.
func (a *Airport) QueryDepartureFlights() *FlightQuery {
	return (&AirportClient{config: a.config}).QueryDepartureFlights(a)
}

// QueryArrivalFlights queries the "arrival_flights" edge of the Airport entity.
func (a *Airport) QueryArrivalFlights() *FlightQuery {
	return (&AirportClient{config: a.config}).QueryArrivalFlights(a)
}

// QueryOriginIteneraries queries the "origin_iteneraries" edge of the Airport entity.
func (a *Airport) QueryOriginIteneraries() *IteneraryQuery {
	return (&AirportClient{config: a.config}).QueryOriginIteneraries(a)
}

// QueryDestinationIteneraries queries the "destination_iteneraries" edge of the Airport entity.
func (a *Airport) QueryDestinationIteneraries() *IteneraryQuery {
	return (&AirportClient{config: a.config}).QueryDestinationIteneraries(a)
}

// Update returns a builder for updating this Airport.
// Note that you need to call Airport.Unwrap() before calling this method if this Airport
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Airport) Update() *AirportUpdateOne {
	return (&AirportClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Airport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Airport) Unwrap() *Airport {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Airport is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Airport) String() string {
	var builder strings.Builder
	builder.WriteString("Airport(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", iata_code=")
	builder.WriteString(a.IataCode)
	builder.WriteString(", icao_code=")
	builder.WriteString(a.IcaoCode)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Airports is a parsable slice of Airport.
type Airports []*Airport

func (a Airports) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}