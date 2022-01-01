package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FrontDesk holds the schema definition for the FrontDesk entity.
type FrontDesk struct {
	ent.Schema
}

// Fields of the FrontDesk.
func (FrontDesk) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("employee_id").MaxLen(50),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FrontDesk.
func (FrontDesk) Edges() []ent.Edge {
	return nil
}
