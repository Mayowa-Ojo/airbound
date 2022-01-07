// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// EdgeCrew holds the string denoting the crew edge name in mutations.
	EdgeCrew = "crew"
	// EdgePilot holds the string denoting the pilot edge name in mutations.
	EdgePilot = "pilot"
	// EdgeFrontDesk holds the string denoting the front_desk edge name in mutations.
	EdgeFrontDesk = "front_desk"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeAddress holds the string denoting the address edge name in mutations.
	EdgeAddress = "address"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "accounts"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "user_id"
	// AdminTable is the table that holds the admin relation/edge.
	AdminTable = "admins"
	// AdminInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminInverseTable = "admins"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "user_id"
	// CrewTable is the table that holds the crew relation/edge.
	CrewTable = "crews"
	// CrewInverseTable is the table name for the Crew entity.
	// It exists in this package in order to avoid circular dependency with the "crew" package.
	CrewInverseTable = "crews"
	// CrewColumn is the table column denoting the crew relation/edge.
	CrewColumn = "user_id"
	// PilotTable is the table that holds the pilot relation/edge.
	PilotTable = "pilots"
	// PilotInverseTable is the table name for the Pilot entity.
	// It exists in this package in order to avoid circular dependency with the "pilot" package.
	PilotInverseTable = "pilots"
	// PilotColumn is the table column denoting the pilot relation/edge.
	PilotColumn = "user_id"
	// FrontDeskTable is the table that holds the front_desk relation/edge.
	FrontDeskTable = "front_desks"
	// FrontDeskInverseTable is the table name for the FrontDesk entity.
	// It exists in this package in order to avoid circular dependency with the "frontdesk" package.
	FrontDeskInverseTable = "front_desks"
	// FrontDeskColumn is the table column denoting the front_desk relation/edge.
	FrontDeskColumn = "user_id"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "customers"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "user_id"
	// AddressTable is the table that holds the address relation/edge.
	AddressTable = "users"
	// AddressInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressInverseTable = "addresses"
	// AddressColumn is the table column denoting the address relation/edge.
	AddressColumn = "address_id"
	// RoleTable is the table that holds the role relation/edge.
	RoleTable = "users"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "role_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFirstname,
	FieldLastname,
	FieldEmail,
	FieldPhone,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"address_id",
	"role_id",
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
	// FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	FirstnameValidator func(string) error
	// LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	LastnameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)