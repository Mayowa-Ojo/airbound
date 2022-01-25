package schema

import (
	"airbound/internal/ent/customtypes"
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightInstance holds the schema definition for the FlightInstance entity.
type FlightInstance struct {
	ent.Schema
}

// Fields of the FlightInstance.
func (FlightInstance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("departure_date").GoType(customtypes.Date{}).SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.String("arrival_date").GoType(customtypes.Date{}).SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.Int("departure_gate").NonNegative(),
		field.Int("arrival_gate").NonNegative(),
		field.Enum("flight_status").GoType(enums.FlightStatus("")),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FlightInstance.
func (FlightInstance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flight", Flight.Type).
			Ref("flight_instances").
			Unique(),
		edge.From("flight_schedule", FlightSchedule.Type).
			Ref("flight_instances").
			Unique(),
		edge.To("aircraft", Aircraft.Type).
			StorageKey(edge.Column("flight_instance_id")).
			Unique(),
		edge.To("flight_reservations", FlightReservation.Type).
			StorageKey(edge.Column("flight_instance_id")),
		edge.To("flight_seats", FlightSeat.Type).
			StorageKey(edge.Column("flight_instance_id")),
	}
}
