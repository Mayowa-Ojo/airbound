package schema

import (
	"time"

	"entgo.io/ent"
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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Aircraft.
func (Aircraft) Edges() []ent.Edge {
	return nil
}
