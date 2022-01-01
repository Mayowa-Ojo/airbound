package schema

import "entgo.io/ent"

// Pilot holds the schema definition for the Pilot entity.
type Pilot struct {
	ent.Schema
}

// Fields of the Pilot.
func (Pilot) Fields() []ent.Field {
	return nil
}

// Edges of the Pilot.
func (Pilot) Edges() []ent.Edge {
	return nil
}
