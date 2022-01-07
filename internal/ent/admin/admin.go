// Code generated by entc, DO NOT EDIT.

package admin

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the admin type in the database.
	Label = "admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTwoFaSecret holds the string denoting the two_fa_secret field in the database.
	FieldTwoFaSecret = "two_fa_secret"
	// FieldTwoFaCompleted holds the string denoting the two_fa_completed field in the database.
	FieldTwoFaCompleted = "two_fa_completed"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the admin in the database.
	Table = "admins"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "admins"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for admin fields.
var Columns = []string{
	FieldID,
	FieldTwoFaSecret,
	FieldTwoFaCompleted,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "admins"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
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
	// TwoFaSecretValidator is a validator for the "two_fa_secret" field. It is called by the builders before save.
	TwoFaSecretValidator func(string) error
	// DefaultTwoFaCompleted holds the default value on creation for the "two_fa_completed" field.
	DefaultTwoFaCompleted bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
