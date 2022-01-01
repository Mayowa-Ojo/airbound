package schema

import "entgo.io/ent"

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return nil
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return nil
}
