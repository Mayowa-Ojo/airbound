package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Pilot holds the schema definition for the Pilot entity.
type Pilot struct {
	ent.Schema
}

// Fields of the Pilot.
func (Pilot) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("employee_id").MaxLen(50),
		field.String("licence_number").MaxLen(50),
		field.Int("flight_hours").NonNegative(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Pilot.
func (Pilot) Edges() []ent.Edge {
	return nil
}
