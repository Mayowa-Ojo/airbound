package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Passenger holds the schema definition for the Passenger entity.
type Passenger struct {
	ent.Schema
}

// Fields of the Passenger.
func (Passenger) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("firstname").MaxLen(250),
		field.String("lastname").MaxLen(250),
		field.Int("age").NonNegative(),
		field.String("passport_number").MaxLen(50),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Passenger.
func (Passenger) Edges() []ent.Edge {
	return nil
}
