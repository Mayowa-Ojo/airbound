// Code generated by entc, DO NOT EDIT.

package customer

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFrequentFlyerNumber holds the string denoting the frequent_flyer_number field in the database.
	FieldFrequentFlyerNumber = "frequent_flyer_number"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeIteneraries holds the string denoting the iteneraries edge name in mutations.
	EdgeIteneraries = "iteneraries"
	// Table holds the table name of the customer in the database.
	Table = "customers"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "customers"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_customer"
	// ItenerariesTable is the table that holds the iteneraries relation/edge.
	ItenerariesTable = "iteneraries"
	// ItenerariesInverseTable is the table name for the Itenerary entity.
	// It exists in this package in order to avoid circular dependency with the "itenerary" package.
	ItenerariesInverseTable = "iteneraries"
	// ItenerariesColumn is the table column denoting the iteneraries relation/edge.
	ItenerariesColumn = "customer_id"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldFrequentFlyerNumber,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "customers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_customer",
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
	// FrequentFlyerNumberValidator is a validator for the "frequent_flyer_number" field. It is called by the builders before save.
	FrequentFlyerNumberValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
