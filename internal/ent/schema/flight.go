package schema

import (
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return nil
}
