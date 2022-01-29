package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Aircraft holds the schema definition for the Aircraft entity.
type Aircraft struct {
	ent.Schema
}

// Fields of the Aircraft.
func (Aircraft) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("tail_number").MaxLen(250),
		field.String("manufacturer").MaxLen(250),
		field.String("model").MaxLen(250),
		field.Int("capacity").NonNegative(),
		field.Int("range").NonNegative(),
		field.Time("manufactured_at"),
		field.Bool("is_grounded").Default(false),
		field.Time("grounded_at").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Aircraft.
func (Aircraft) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("airline", Airline.Type).
			Ref("aircrafts").
			Unique(),
		edge.From("flight_instance", FlightInstance.Type).
			Ref("aircraft").
			Unique(),
		edge.To("seats", Seat.Type).
			StorageKey(edge.Column("aircraft_id")),
	}
}
