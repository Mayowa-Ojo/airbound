// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "account_status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "CLOSED", "BLACKLISTED", "BLOCKED", "NONE"}},
		{Name: "password", Type: field.TypeBytes},
		{Name: "salt", Type: field.TypeBytes},
		{Name: "two_fa_secret", Type: field.TypeString, Nullable: true, Size: 250},
		{Name: "two_fa_completed", Type: field.TypeBool, Default: false},
		{Name: "verification_token", Type: field.TypeString, Nullable: true},
		{Name: "forgot_password_token", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_account",
				Columns:    []*schema.Column{AccountsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "street", Type: field.TypeString, Size: 150},
		{Name: "city", Type: field.TypeString, Size: 150},
		{Name: "state", Type: field.TypeString, Size: 150},
		{Name: "zipcode", Type: field.TypeString, Size: 50},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
	}
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "level", Type: field.TypeInt, Default: 1},
		{Name: "security_question", Type: field.TypeString, Nullable: true, Size: 500},
		{Name: "security_answer", Type: field.TypeString, Nullable: true, Size: 250},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "admins_users_admin",
				Columns:    []*schema.Column{AdminsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AircraftsColumns holds the columns for the "aircrafts" table.
	AircraftsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "tail_number", Type: field.TypeString, Size: 250},
		{Name: "manufacturer", Type: field.TypeString, Size: 250},
		{Name: "model", Type: field.TypeString, Size: 250},
		{Name: "capacity", Type: field.TypeInt},
		{Name: "range", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "airline_id", Type: field.TypeUUID, Nullable: true},
		{Name: "flight_instance_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// AircraftsTable holds the schema information for the "aircrafts" table.
	AircraftsTable = &schema.Table{
		Name:       "aircrafts",
		Columns:    AircraftsColumns,
		PrimaryKey: []*schema.Column{AircraftsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "aircrafts_airlines_aircrafts",
				Columns:    []*schema.Column{AircraftsColumns[8]},
				RefColumns: []*schema.Column{AirlinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "aircrafts_flight_instances_aircraft",
				Columns:    []*schema.Column{AircraftsColumns[9]},
				RefColumns: []*schema.Column{FlightInstancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AirlinesColumns holds the columns for the "airlines" table.
	AirlinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 250},
		{Name: "iata_code", Type: field.TypeString, Size: 2},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AirlinesTable holds the schema information for the "airlines" table.
	AirlinesTable = &schema.Table{
		Name:       "airlines",
		Columns:    AirlinesColumns,
		PrimaryKey: []*schema.Column{AirlinesColumns[0]},
	}
	// AirportsColumns holds the columns for the "airports" table.
	AirportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 250},
		{Name: "iata_code", Type: field.TypeString, Size: 3},
		{Name: "icao_code", Type: field.TypeString, Size: 4},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "address_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// AirportsTable holds the schema information for the "airports" table.
	AirportsTable = &schema.Table{
		Name:       "airports",
		Columns:    AirportsColumns,
		PrimaryKey: []*schema.Column{AirportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "airports_addresses_airport",
				Columns:    []*schema.Column{AirportsColumns[6]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CrewsColumns holds the columns for the "crews" table.
	CrewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "employee_id", Type: field.TypeString, Size: 50},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "airline_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// CrewsTable holds the schema information for the "crews" table.
	CrewsTable = &schema.Table{
		Name:       "crews",
		Columns:    CrewsColumns,
		PrimaryKey: []*schema.Column{CrewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "crews_airlines_crews",
				Columns:    []*schema.Column{CrewsColumns[4]},
				RefColumns: []*schema.Column{AirlinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "crews_users_crew",
				Columns:    []*schema.Column{CrewsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "frequent_flyer_number", Type: field.TypeString, Size: 50},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customers_users_customer",
				Columns:    []*schema.Column{CustomersColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlightsColumns holds the columns for the "flights" table.
	FlightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "flight_number", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "duration", Type: field.TypeInt},
		{Name: "distance", Type: field.TypeInt},
		{Name: "boarding_policy", Type: field.TypeEnum, Enums: []string{"GROUP_BASED", "ZONE_BASED"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "depature_airport_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arrival_airport_id", Type: field.TypeUUID, Nullable: true},
	}
	// FlightsTable holds the schema information for the "flights" table.
	FlightsTable = &schema.Table{
		Name:       "flights",
		Columns:    FlightsColumns,
		PrimaryKey: []*schema.Column{FlightsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flights_airports_departure_flights",
				Columns:    []*schema.Column{FlightsColumns[7]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "flights_airports_arrival_flights",
				Columns:    []*schema.Column{FlightsColumns[8]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlightInstancesColumns holds the columns for the "flight_instances" table.
	FlightInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "departure_gate", Type: field.TypeInt},
		{Name: "arrival_gate", Type: field.TypeInt},
		{Name: "flight_status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "SCHEDULED", "DELAYED", "DEPARTED", "ARRIVED", "CANCELED"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "flight_id", Type: field.TypeUUID, Nullable: true},
	}
	// FlightInstancesTable holds the schema information for the "flight_instances" table.
	FlightInstancesTable = &schema.Table{
		Name:       "flight_instances",
		Columns:    FlightInstancesColumns,
		PrimaryKey: []*schema.Column{FlightInstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flight_instances_flights_flight_instances",
				Columns:    []*schema.Column{FlightInstancesColumns[6]},
				RefColumns: []*schema.Column{FlightsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlightReservationsColumns holds the columns for the "flight_reservations" table.
	FlightReservationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "reservation_number", Type: field.TypeString, Size: 50},
		{Name: "reservation_status", Type: field.TypeEnum, Enums: []string{"REQUESTED", "PENDING", "CONFIRMED", "CANCELED", "CHECKED_IN", "ABANDONED"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "flight_instance_id", Type: field.TypeUUID, Nullable: true},
		{Name: "itenerary_id", Type: field.TypeUUID, Nullable: true},
	}
	// FlightReservationsTable holds the schema information for the "flight_reservations" table.
	FlightReservationsTable = &schema.Table{
		Name:       "flight_reservations",
		Columns:    FlightReservationsColumns,
		PrimaryKey: []*schema.Column{FlightReservationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flight_reservations_flight_instances_flight_reservations",
				Columns:    []*schema.Column{FlightReservationsColumns[5]},
				RefColumns: []*schema.Column{FlightInstancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "flight_reservations_iteneraries_flight_reservations",
				Columns:    []*schema.Column{FlightReservationsColumns[6]},
				RefColumns: []*schema.Column{ItenerariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlightSchedulesColumns holds the columns for the "flight_schedules" table.
	FlightSchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "weekday", Type: field.TypeEnum, Nullable: true, Enums: []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}},
		{Name: "schedule_type", Type: field.TypeEnum, Enums: []string{"WEEKLY", "CUSTOM"}},
		{Name: "custom_date", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "departs_at", Type: field.TypeString},
		{Name: "arrives_at", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "flight_id", Type: field.TypeUUID, Nullable: true},
	}
	// FlightSchedulesTable holds the schema information for the "flight_schedules" table.
	FlightSchedulesTable = &schema.Table{
		Name:       "flight_schedules",
		Columns:    FlightSchedulesColumns,
		PrimaryKey: []*schema.Column{FlightSchedulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flight_schedules_flights_flight_schedules",
				Columns:    []*schema.Column{FlightSchedulesColumns[8]},
				RefColumns: []*schema.Column{FlightsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlightSeatsColumns holds the columns for the "flight_seats" table.
	FlightSeatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "fare", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "flight_instance_id", Type: field.TypeUUID, Nullable: true},
		{Name: "passenger_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "seat_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// FlightSeatsTable holds the schema information for the "flight_seats" table.
	FlightSeatsTable = &schema.Table{
		Name:       "flight_seats",
		Columns:    FlightSeatsColumns,
		PrimaryKey: []*schema.Column{FlightSeatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flight_seats_flight_instances_flight_seats",
				Columns:    []*schema.Column{FlightSeatsColumns[4]},
				RefColumns: []*schema.Column{FlightInstancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "flight_seats_passengers_flight_seat",
				Columns:    []*schema.Column{FlightSeatsColumns[5]},
				RefColumns: []*schema.Column{PassengersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "flight_seats_seats_flight_seat",
				Columns:    []*schema.Column{FlightSeatsColumns[6]},
				RefColumns: []*schema.Column{SeatsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FrontDesksColumns holds the columns for the "front_desks" table.
	FrontDesksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "employee_id", Type: field.TypeString, Size: 50},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "airport_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// FrontDesksTable holds the schema information for the "front_desks" table.
	FrontDesksTable = &schema.Table{
		Name:       "front_desks",
		Columns:    FrontDesksColumns,
		PrimaryKey: []*schema.Column{FrontDesksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "front_desks_airports_front_desks",
				Columns:    []*schema.Column{FrontDesksColumns[4]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "front_desks_users_front_desk",
				Columns:    []*schema.Column{FrontDesksColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ItenerariesColumns holds the columns for the "iteneraries" table.
	ItenerariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "origin_airport_id", Type: field.TypeUUID, Nullable: true},
		{Name: "destination_airport_id", Type: field.TypeUUID, Nullable: true},
		{Name: "customer_id", Type: field.TypeUUID, Nullable: true},
	}
	// ItenerariesTable holds the schema information for the "iteneraries" table.
	ItenerariesTable = &schema.Table{
		Name:       "iteneraries",
		Columns:    ItenerariesColumns,
		PrimaryKey: []*schema.Column{ItenerariesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "iteneraries_airports_origin_iteneraries",
				Columns:    []*schema.Column{ItenerariesColumns[3]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "iteneraries_airports_destination_iteneraries",
				Columns:    []*schema.Column{ItenerariesColumns[4]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "iteneraries_customers_iteneraries",
				Columns:    []*schema.Column{ItenerariesColumns[5]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PassengersColumns holds the columns for the "passengers" table.
	PassengersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "firstname", Type: field.TypeString, Size: 250},
		{Name: "lastname", Type: field.TypeString, Size: 250},
		{Name: "age", Type: field.TypeInt},
		{Name: "passport_number", Type: field.TypeString, Size: 50},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "flight_reservation_id", Type: field.TypeUUID, Nullable: true},
	}
	// PassengersTable holds the schema information for the "passengers" table.
	PassengersTable = &schema.Table{
		Name:       "passengers",
		Columns:    PassengersColumns,
		PrimaryKey: []*schema.Column{PassengersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "passengers_flight_reservations_passengers",
				Columns:    []*schema.Column{PassengersColumns[7]},
				RefColumns: []*schema.Column{FlightReservationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "permission", Type: field.TypeString, Size: 1000},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
	}
	// PilotsColumns holds the columns for the "pilots" table.
	PilotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "employee_id", Type: field.TypeString, Size: 50},
		{Name: "licence_number", Type: field.TypeString, Size: 50},
		{Name: "flight_hours", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "airline_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// PilotsTable holds the schema information for the "pilots" table.
	PilotsTable = &schema.Table{
		Name:       "pilots",
		Columns:    PilotsColumns,
		PrimaryKey: []*schema.Column{PilotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pilots_airlines_pilots",
				Columns:    []*schema.Column{PilotsColumns[6]},
				RefColumns: []*schema.Column{AirlinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "pilots_users_pilot",
				Columns:    []*schema.Column{PilotsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeEnum, Enums: []string{"ADMIN", "PILOT", "CREW", "FRONT_DESK", "CUSTOMER"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SeatsColumns holds the columns for the "seats" table.
	SeatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "seat_number", Type: field.TypeInt},
		{Name: "seat_row", Type: field.TypeString, Size: 1},
		{Name: "seat_type", Type: field.TypeEnum, Enums: []string{"REGULAR", "EMERGENCY_EXIT", "ACCESSIBLE"}},
		{Name: "seat_class", Type: field.TypeEnum, Enums: []string{"ECONOMY", "ECONOMY_PLUS", "BUSINESS", "FIRST"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "aircraft_id", Type: field.TypeUUID, Nullable: true},
	}
	// SeatsTable holds the schema information for the "seats" table.
	SeatsTable = &schema.Table{
		Name:       "seats",
		Columns:    SeatsColumns,
		PrimaryKey: []*schema.Column{SeatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "seats_aircrafts_seats",
				Columns:    []*schema.Column{SeatsColumns[7]},
				RefColumns: []*schema.Column{AircraftsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "firstname", Type: field.TypeString, Size: 250},
		{Name: "lastname", Type: field.TypeString, Size: 250},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 250},
		{Name: "phone", Type: field.TypeString, Size: 15},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "address_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "role_id", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_addresses_user",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_roles_users",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FlightCrewColumns holds the columns for the "flight_crew" table.
	FlightCrewColumns = []*schema.Column{
		{Name: "flight_id", Type: field.TypeUUID},
		{Name: "crew_id", Type: field.TypeUUID},
	}
	// FlightCrewTable holds the schema information for the "flight_crew" table.
	FlightCrewTable = &schema.Table{
		Name:       "flight_crew",
		Columns:    FlightCrewColumns,
		PrimaryKey: []*schema.Column{FlightCrewColumns[0], FlightCrewColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flight_crew_flight_id",
				Columns:    []*schema.Column{FlightCrewColumns[0]},
				RefColumns: []*schema.Column{FlightsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "flight_crew_crew_id",
				Columns:    []*schema.Column{FlightCrewColumns[1]},
				RefColumns: []*schema.Column{CrewsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RolePermissionColumns holds the columns for the "role_permission" table.
	RolePermissionColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "permission_id", Type: field.TypeUUID},
	}
	// RolePermissionTable holds the schema information for the "role_permission" table.
	RolePermissionTable = &schema.Table{
		Name:       "role_permission",
		Columns:    RolePermissionColumns,
		PrimaryKey: []*schema.Column{RolePermissionColumns[0], RolePermissionColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_permission_role_id",
				Columns:    []*schema.Column{RolePermissionColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_permission_permission_id",
				Columns:    []*schema.Column{RolePermissionColumns[1]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AddressesTable,
		AdminsTable,
		AircraftsTable,
		AirlinesTable,
		AirportsTable,
		CrewsTable,
		CustomersTable,
		FlightsTable,
		FlightInstancesTable,
		FlightReservationsTable,
		FlightSchedulesTable,
		FlightSeatsTable,
		FrontDesksTable,
		ItenerariesTable,
		PassengersTable,
		PermissionsTable,
		PilotsTable,
		RolesTable,
		SeatsTable,
		UsersTable,
		FlightCrewTable,
		RolePermissionTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	AdminsTable.ForeignKeys[0].RefTable = UsersTable
	AircraftsTable.ForeignKeys[0].RefTable = AirlinesTable
	AircraftsTable.ForeignKeys[1].RefTable = FlightInstancesTable
	AirportsTable.ForeignKeys[0].RefTable = AddressesTable
	CrewsTable.ForeignKeys[0].RefTable = AirlinesTable
	CrewsTable.ForeignKeys[1].RefTable = UsersTable
	CustomersTable.ForeignKeys[0].RefTable = UsersTable
	FlightsTable.ForeignKeys[0].RefTable = AirportsTable
	FlightsTable.ForeignKeys[1].RefTable = AirportsTable
	FlightInstancesTable.ForeignKeys[0].RefTable = FlightsTable
	FlightReservationsTable.ForeignKeys[0].RefTable = FlightInstancesTable
	FlightReservationsTable.ForeignKeys[1].RefTable = ItenerariesTable
	FlightSchedulesTable.ForeignKeys[0].RefTable = FlightsTable
	FlightSeatsTable.ForeignKeys[0].RefTable = FlightInstancesTable
	FlightSeatsTable.ForeignKeys[1].RefTable = PassengersTable
	FlightSeatsTable.ForeignKeys[2].RefTable = SeatsTable
	FrontDesksTable.ForeignKeys[0].RefTable = AirportsTable
	FrontDesksTable.ForeignKeys[1].RefTable = UsersTable
	ItenerariesTable.ForeignKeys[0].RefTable = AirportsTable
	ItenerariesTable.ForeignKeys[1].RefTable = AirportsTable
	ItenerariesTable.ForeignKeys[2].RefTable = CustomersTable
	PassengersTable.ForeignKeys[0].RefTable = FlightReservationsTable
	PilotsTable.ForeignKeys[0].RefTable = AirlinesTable
	PilotsTable.ForeignKeys[1].RefTable = UsersTable
	SeatsTable.ForeignKeys[0].RefTable = AircraftsTable
	UsersTable.ForeignKeys[0].RefTable = AddressesTable
	UsersTable.ForeignKeys[1].RefTable = RolesTable
	FlightCrewTable.ForeignKeys[0].RefTable = FlightsTable
	FlightCrewTable.ForeignKeys[1].RefTable = CrewsTable
	RolePermissionTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionTable.ForeignKeys[1].RefTable = PermissionsTable
}
