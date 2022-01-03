package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Crew holds the schema definition for the Crew entity.
type Crew struct {
	ent.Schema
}

// Fields of the Crew.
func (Crew) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("employee_id").MaxLen(50),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Crew.
func (Crew) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("crew").
			Unique().
			Required(),
		edge.From("airline", Airline.Type).
			Ref("crews").
			Unique(),
		edge.From("flights", Flight.Type).Ref("crews"),
	}
}
