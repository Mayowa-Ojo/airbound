package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Airline holds the schema definition for the Airline entity.
type Airline struct {
	ent.Schema
}

// Fields of the Airline.
func (Airline) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("name").MaxLen(250),
		field.String("iata_code").MaxLen(2).Unique(),
		field.String("icao_code").MaxLen(3).Unique(),
		field.String("call_sign").MaxLen(250).Unique(),
		field.String("country").MaxLen(250),
		field.String("license_code").MaxLen(250).Unique(),
		field.Int("fleet_size").NonNegative(),
		field.Int("ranking").Range(0, 6),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Airline.
func (Airline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("aircrafts", Aircraft.Type).
			StorageKey(edge.Column("airline_id")),
		edge.To("crews", Crew.Type).
			StorageKey(edge.Column("airline_id")),
		edge.To("pilots", Pilot.Type).
			StorageKey(edge.Column("airline_id")),
		edge.To("flights", Flight.Type).
			StorageKey(edge.Column("airline_id")),
	}
}
