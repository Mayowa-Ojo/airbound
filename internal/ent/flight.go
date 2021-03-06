// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/airline"
	"airbound/internal/ent/airport"
	"airbound/internal/ent/enums"
	"airbound/internal/ent/flight"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Flight is the model entity for the Flight schema.
type Flight struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FlightNumber holds the value of the "flight_number" field.
	FlightNumber string `json:"flight_number,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// Distance holds the value of the "distance" field.
	Distance int `json:"distance,omitempty"`
	// BoardingPolicy holds the value of the "boarding_policy" field.
	BoardingPolicy enums.BoardingPolicy `json:"boarding_policy,omitempty"`
	// TripType holds the value of the "trip_type" field.
	TripType enums.TripType `json:"trip_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlightQuery when eager-loading is set.
	Edges               FlightEdges `json:"edges"`
	airline_id          *uuid.UUID
	depature_airport_id *uuid.UUID
	arrival_airport_id  *uuid.UUID
}

// FlightEdges holds the relations/edges for other nodes in the graph.
type FlightEdges struct {
	// FlightInstances holds the value of the flight_instances edge.
	FlightInstances []*FlightInstance `json:"flight_instances,omitempty"`
	// FlightSchedules holds the value of the flight_schedules edge.
	FlightSchedules []*FlightSchedule `json:"flight_schedules,omitempty"`
	// Crews holds the value of the crews edge.
	Crews []*Crew `json:"crews,omitempty"`
	// DepartureAirport holds the value of the departure_airport edge.
	DepartureAirport *Airport `json:"departure_airport,omitempty"`
	// ArrivalAirport holds the value of the arrival_airport edge.
	ArrivalAirport *Airport `json:"arrival_airport,omitempty"`
	// Airline holds the value of the airline edge.
	Airline *Airline `json:"airline,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// FlightInstancesOrErr returns the FlightInstances value or an error if the edge
// was not loaded in eager-loading.
func (e FlightEdges) FlightInstancesOrErr() ([]*FlightInstance, error) {
	if e.loadedTypes[0] {
		return e.FlightInstances, nil
	}
	return nil, &NotLoadedError{edge: "flight_instances"}
}

// FlightSchedulesOrErr returns the FlightSchedules value or an error if the edge
// was not loaded in eager-loading.
func (e FlightEdges) FlightSchedulesOrErr() ([]*FlightSchedule, error) {
	if e.loadedTypes[1] {
		return e.FlightSchedules, nil
	}
	return nil, &NotLoadedError{edge: "flight_schedules"}
}

// CrewsOrErr returns the Crews value or an error if the edge
// was not loaded in eager-loading.
func (e FlightEdges) CrewsOrErr() ([]*Crew, error) {
	if e.loadedTypes[2] {
		return e.Crews, nil
	}
	return nil, &NotLoadedError{edge: "crews"}
}

// DepartureAirportOrErr returns the DepartureAirport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightEdges) DepartureAirportOrErr() (*Airport, error) {
	if e.loadedTypes[3] {
		if e.DepartureAirport == nil {
			// The edge departure_airport was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.DepartureAirport, nil
	}
	return nil, &NotLoadedError{edge: "departure_airport"}
}

// ArrivalAirportOrErr returns the ArrivalAirport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightEdges) ArrivalAirportOrErr() (*Airport, error) {
	if e.loadedTypes[4] {
		if e.ArrivalAirport == nil {
			// The edge arrival_airport was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.ArrivalAirport, nil
	}
	return nil, &NotLoadedError{edge: "arrival_airport"}
}

// AirlineOrErr returns the Airline value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightEdges) AirlineOrErr() (*Airline, error) {
	if e.loadedTypes[5] {
		if e.Airline == nil {
			// The edge airline was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: airline.Label}
		}
		return e.Airline, nil
	}
	return nil, &NotLoadedError{edge: "airline"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Flight) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case flight.FieldDuration, flight.FieldDistance:
			values[i] = new(sql.NullInt64)
		case flight.FieldFlightNumber, flight.FieldBoardingPolicy, flight.FieldTripType:
			values[i] = new(sql.NullString)
		case flight.FieldCreatedAt, flight.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case flight.FieldID:
			values[i] = new(uuid.UUID)
		case flight.ForeignKeys[0]: // airline_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case flight.ForeignKeys[1]: // depature_airport_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case flight.ForeignKeys[2]: // arrival_airport_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Flight", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Flight fields.
func (f *Flight) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flight.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case flight.FieldFlightNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flight_number", values[i])
			} else if value.Valid {
				f.FlightNumber = value.String
			}
		case flight.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				f.Duration = int(value.Int64)
			}
		case flight.FieldDistance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field distance", values[i])
			} else if value.Valid {
				f.Distance = int(value.Int64)
			}
		case flight.FieldBoardingPolicy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field boarding_policy", values[i])
			} else if value.Valid {
				f.BoardingPolicy = enums.BoardingPolicy(value.String)
			}
		case flight.FieldTripType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trip_type", values[i])
			} else if value.Valid {
				f.TripType = enums.TripType(value.String)
			}
		case flight.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case flight.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case flight.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field airline_id", values[i])
			} else if value.Valid {
				f.airline_id = new(uuid.UUID)
				*f.airline_id = *value.S.(*uuid.UUID)
			}
		case flight.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field depature_airport_id", values[i])
			} else if value.Valid {
				f.depature_airport_id = new(uuid.UUID)
				*f.depature_airport_id = *value.S.(*uuid.UUID)
			}
		case flight.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field arrival_airport_id", values[i])
			} else if value.Valid {
				f.arrival_airport_id = new(uuid.UUID)
				*f.arrival_airport_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFlightInstances queries the "flight_instances" edge of the Flight entity.
func (f *Flight) QueryFlightInstances() *FlightInstanceQuery {
	return (&FlightClient{config: f.config}).QueryFlightInstances(f)
}

// QueryFlightSchedules queries the "flight_schedules" edge of the Flight entity.
func (f *Flight) QueryFlightSchedules() *FlightScheduleQuery {
	return (&FlightClient{config: f.config}).QueryFlightSchedules(f)
}

// QueryCrews queries the "crews" edge of the Flight entity.
func (f *Flight) QueryCrews() *CrewQuery {
	return (&FlightClient{config: f.config}).QueryCrews(f)
}

// QueryDepartureAirport queries the "departure_airport" edge of the Flight entity.
func (f *Flight) QueryDepartureAirport() *AirportQuery {
	return (&FlightClient{config: f.config}).QueryDepartureAirport(f)
}

// QueryArrivalAirport queries the "arrival_airport" edge of the Flight entity.
func (f *Flight) QueryArrivalAirport() *AirportQuery {
	return (&FlightClient{config: f.config}).QueryArrivalAirport(f)
}

// QueryAirline queries the "airline" edge of the Flight entity.
func (f *Flight) QueryAirline() *AirlineQuery {
	return (&FlightClient{config: f.config}).QueryAirline(f)
}

// Update returns a builder for updating this Flight.
// Note that you need to call Flight.Unwrap() before calling this method if this Flight
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Flight) Update() *FlightUpdateOne {
	return (&FlightClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Flight entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Flight) Unwrap() *Flight {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Flight is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Flight) String() string {
	var builder strings.Builder
	builder.WriteString("Flight(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", flight_number=")
	builder.WriteString(f.FlightNumber)
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", f.Duration))
	builder.WriteString(", distance=")
	builder.WriteString(fmt.Sprintf("%v", f.Distance))
	builder.WriteString(", boarding_policy=")
	builder.WriteString(fmt.Sprintf("%v", f.BoardingPolicy))
	builder.WriteString(", trip_type=")
	builder.WriteString(fmt.Sprintf("%v", f.TripType))
	builder.WriteString(", created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Flights is a parsable slice of Flight.
type Flights []*Flight

func (f Flights) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
