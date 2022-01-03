// Code generated by entc, DO NOT EDIT.

package flightseat

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the flightseat type in the database.
	Label = "flight_seat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFare holds the string denoting the fare field in the database.
	FieldFare = "fare"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeFlightInstance holds the string denoting the flight_instance edge name in mutations.
	EdgeFlightInstance = "flight_instance"
	// EdgeSeat holds the string denoting the seat edge name in mutations.
	EdgeSeat = "seat"
	// EdgePassenger holds the string denoting the passenger edge name in mutations.
	EdgePassenger = "passenger"
	// Table holds the table name of the flightseat in the database.
	Table = "flight_seats"
	// FlightInstanceTable is the table that holds the flight_instance relation/edge.
	FlightInstanceTable = "flight_seats"
	// FlightInstanceInverseTable is the table name for the FlightInstance entity.
	// It exists in this package in order to avoid circular dependency with the "flightinstance" package.
	FlightInstanceInverseTable = "flight_instances"
	// FlightInstanceColumn is the table column denoting the flight_instance relation/edge.
	FlightInstanceColumn = "flight_instance_id"
	// SeatTable is the table that holds the seat relation/edge.
	SeatTable = "flight_seats"
	// SeatInverseTable is the table name for the Seat entity.
	// It exists in this package in order to avoid circular dependency with the "seat" package.
	SeatInverseTable = "seats"
	// SeatColumn is the table column denoting the seat relation/edge.
	SeatColumn = "seat_flight_seat"
	// PassengerTable is the table that holds the passenger relation/edge.
	PassengerTable = "flight_seats"
	// PassengerInverseTable is the table name for the Passenger entity.
	// It exists in this package in order to avoid circular dependency with the "passenger" package.
	PassengerInverseTable = "passengers"
	// PassengerColumn is the table column denoting the passenger relation/edge.
	PassengerColumn = "passenger_flight_seat"
)

// Columns holds all SQL columns for flightseat fields.
var Columns = []string{
	FieldID,
	FieldFare,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "flight_seats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"flight_instance_id",
	"passenger_flight_seat",
	"seat_flight_seat",
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
	// FareValidator is a validator for the "fare" field. It is called by the builders before save.
	FareValidator func(float64) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
