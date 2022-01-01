package schema

import "entgo.io/ent"

// Aircraft holds the schema definition for the Aircraft entity.
type Aircraft struct {
	ent.Schema
}

// Fields of the Aircraft.
func (Aircraft) Fields() []ent.Field {
	return nil
}

// Edges of the Aircraft.
func (Aircraft) Edges() []ent.Edge {
	return nil
}
