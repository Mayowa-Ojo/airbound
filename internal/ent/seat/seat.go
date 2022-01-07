// Code generated by entc, DO NOT EDIT.

package seat

import (
	"airbound/internal/ent/enums"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the seat type in the database.
	Label = "seat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSeatNumber holds the string denoting the seat_number field in the database.
	FieldSeatNumber = "seat_number"
	// FieldSeatRow holds the string denoting the seat_row field in the database.
	FieldSeatRow = "seat_row"
	// FieldSeatType holds the string denoting the seat_type field in the database.
	FieldSeatType = "seat_type"
	// FieldSeatClass holds the string denoting the seat_class field in the database.
	FieldSeatClass = "seat_class"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAircraft holds the string denoting the aircraft edge name in mutations.
	EdgeAircraft = "aircraft"
	// EdgeFlightSeat holds the string denoting the flight_seat edge name in mutations.
	EdgeFlightSeat = "flight_seat"
	// Table holds the table name of the seat in the database.
	Table = "seats"
	// AircraftTable is the table that holds the aircraft relation/edge.
	AircraftTable = "seats"
	// AircraftInverseTable is the table name for the Aircraft entity.
	// It exists in this package in order to avoid circular dependency with the "aircraft" package.
	AircraftInverseTable = "aircrafts"
	// AircraftColumn is the table column denoting the aircraft relation/edge.
	AircraftColumn = "aircraft_id"
	// FlightSeatTable is the table that holds the flight_seat relation/edge.
	FlightSeatTable = "flight_seats"
	// FlightSeatInverseTable is the table name for the FlightSeat entity.
	// It exists in this package in order to avoid circular dependency with the "flightseat" package.
	FlightSeatInverseTable = "flight_seats"
	// FlightSeatColumn is the table column denoting the flight_seat relation/edge.
	FlightSeatColumn = "seat_id"
)

// Columns holds all SQL columns for seat fields.
var Columns = []string{
	FieldID,
	FieldSeatNumber,
	FieldSeatRow,
	FieldSeatType,
	FieldSeatClass,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "seats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"aircraft_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// SeatNumberValidator is a validator for the "seat_number" field. It is called by the builders before save.
	SeatNumberValidator func(int) error
	// SeatRowValidator is a validator for the "seat_row" field. It is called by the builders before save.
	SeatRowValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// SeatTypeValidator is a validator for the "seat_type" field enum values. It is called by the builders before save.
func SeatTypeValidator(st enums.SeatType) error {
	switch st {
	case "REGULAR", "EMERGENCY_EXIT", "ACCESSIBLE":
		return nil
	default:
		return fmt.Errorf("seat: invalid enum value for seat_type field: %q", st)
	}
}

// SeatClassValidator is a validator for the "seat_class" field enum values. It is called by the builders before save.
func SeatClassValidator(sc enums.SeatClass) error {
	switch sc {
	case "ECONOMY", "ECONOMY_PLUS", "BUSINESS", "FIRST":
		return nil
	default:
		return fmt.Errorf("seat: invalid enum value for seat_class field: %q", sc)
	}
}
