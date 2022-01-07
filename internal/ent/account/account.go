// Code generated by entc, DO NOT EDIT.

package account

import (
	"airbound/internal/ent/enums"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountStatus holds the string denoting the account_status field in the database.
	FieldAccountStatus = "account_status"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "accounts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldAccountStatus,
	FieldPassword,
	FieldSalt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "accounts"
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// AccountStatusValidator is a validator for the "account_status" field enum values. It is called by the builders before save.
func AccountStatusValidator(as enums.AccountStatus) error {
	switch as {
	case "ACTIVE", "CLOSED", "BLACKLISTED", "BLOCKED", "NONE":
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for account_status field: %q", as)
	}
}