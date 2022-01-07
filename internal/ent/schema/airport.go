package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Airport holds the schema definition for the Airport entity.
type Airport struct {
	ent.Schema
}

// Fields of the Airport.
func (Airport) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("name").MaxLen(250),
		field.String("iata_code").MaxLen(3),
		field.String("icao_code").MaxLen(4),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("address", Address.Type).
			Ref("airport").
			Unique(),
		edge.To("front_desks", FrontDesk.Type).
			StorageKey(edge.Column("airport_id")),
		edge.To("departure_flights", Flight.Type).
			StorageKey(edge.Column("depature_airport_id")),
		edge.To("arrival_flights", Flight.Type).
			StorageKey(edge.Column("arrival_airport_id")),
		edge.To("origin_iteneraries", Itenerary.Type).
			StorageKey(edge.Column("origin_airport_id")),
		edge.To("destination_iteneraries", Itenerary.Type).
			StorageKey(edge.Column("destination_airport_id")),
	}
}
