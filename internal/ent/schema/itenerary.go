package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Itenerary holds the schema definition for the Itenerary entity.
type Itenerary struct {
	ent.Schema
}

// Fields of the Itenerary.
func (Itenerary) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Itenerary.
func (Itenerary) Edges() []ent.Edge {
	return nil
}
