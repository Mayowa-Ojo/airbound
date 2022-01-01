package schema

import "entgo.io/ent"

// Itenerary holds the schema definition for the Itenerary entity.
type Itenerary struct {
	ent.Schema
}

// Fields of the Itenerary.
func (Itenerary) Fields() []ent.Field {
	return nil
}

// Edges of the Itenerary.
func (Itenerary) Edges() []ent.Edge {
	return nil
}
