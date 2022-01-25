package schema

import (
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("flight_number").MaxLen(20).Unique(),
		field.Int("duration").NonNegative(),
		field.Int("distance").NonNegative(),
		field.Enum("boarding_policy").GoType(enums.BoardingPolicy("")),
		field.Enum("trip_type").GoType(enums.TripType("")),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("flight_instances", FlightInstance.Type).
			StorageKey(edge.Column("flight_id")),
		edge.To("flight_schedules", FlightSchedule.Type).
			StorageKey(edge.Column("flight_id")),
		edge.To("crews", Crew.Type).
			StorageKey(edge.Table("flight_crew"), edge.Columns("flight_id", "crew_id")),
		edge.From("departure_airport", Airport.Type).Ref("departure_flights").Unique(),
		edge.From("arrival_airport", Airport.Type).Ref("arrival_flights").Unique(),
		edge.From("airline", Airline.Type).Ref("flights").Unique(),
	}
}
