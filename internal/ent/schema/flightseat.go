package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightSeat holds the schema definition for the FlightSeat entity.
type FlightSeat struct {
	ent.Schema
}

// Fields of the FlightSeat.
func (FlightSeat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Float("fare").Positive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FlightSeat.
func (FlightSeat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flight_instance", FlightInstance.Type).
			Ref("flight_seats").
			Unique(),
		edge.From("seat", Seat.Type).
			Ref("flight_seat").
			Unique().
			Required(),
		edge.From("passenger", Passenger.Type).
			Ref("flight_seat").
			Unique(),
	}
}
