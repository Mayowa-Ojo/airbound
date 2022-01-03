// Code generated by entc, DO NOT EDIT.

package role

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "role_id"
	// PermissionsTable is the table that holds the permissions relation/edge. The primary key declared below.
	PermissionsTable = "role_permission"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permissions"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// PermissionsPrimaryKey and PermissionsColumn2 are the table columns denoting the
	// primary key for the permissions relation (M2M).
	PermissionsPrimaryKey = []string{"role_id", "permission_id"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
