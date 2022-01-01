package schema

import "entgo.io/ent"

// Airport holds the schema definition for the Airport entity.
type Airport struct {
	ent.Schema
}

// Fields of the Airport.
func (Airport) Fields() []ent.Field {
	return nil
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return nil
}
