package schema

import (
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
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
		field.Int("departure_gate").NonNegative(),
		field.Int("arrival_gate").NonNegative(),
		field.Time("departs_at"),
		field.Time("arrives_at"),
		field.Enum("flight_status").GoType(enums.FlightStatus("")),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FlightInstance.
func (FlightInstance) Edges() []ent.Edge {
	return nil
}
