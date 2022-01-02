// Code generated by entc, DO NOT EDIT.

package aircraft

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the aircraft type in the database.
	Label = "aircraft"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTailNumber holds the string denoting the tail_number field in the database.
	FieldTailNumber = "tail_number"
	// FieldManufacturer holds the string denoting the manufacturer field in the database.
	FieldManufacturer = "manufacturer"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldCapacity holds the string denoting the capacity field in the database.
	FieldCapacity = "capacity"
	// FieldRange holds the string denoting the range field in the database.
	FieldRange = "range"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the aircraft in the database.
	Table = "aircrafts"
)

// Columns holds all SQL columns for aircraft fields.
var Columns = []string{
	FieldID,
	FieldTailNumber,
	FieldManufacturer,
	FieldModel,
	FieldCapacity,
	FieldRange,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TailNumberValidator is a validator for the "tail_number" field. It is called by the builders before save.
	TailNumberValidator func(string) error
	// ManufacturerValidator is a validator for the "manufacturer" field. It is called by the builders before save.
	ManufacturerValidator func(string) error
	// ModelValidator is a validator for the "model" field. It is called by the builders before save.
	ModelValidator func(string) error
	// CapacityValidator is a validator for the "capacity" field. It is called by the builders before save.
	CapacityValidator func(int) error
	// RangeValidator is a validator for the "range" field. It is called by the builders before save.
	RangeValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
