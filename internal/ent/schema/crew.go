package schema

import "entgo.io/ent"

// Crew holds the schema definition for the Crew entity.
type Crew struct {
	ent.Schema
}

// Fields of the Crew.
func (Crew) Fields() []ent.Field {
	return nil
}

// Edges of the Crew.
func (Crew) Edges() []ent.Edge {
	return nil
}
