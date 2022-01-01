package schema

import "entgo.io/ent"

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return nil
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return nil
}
