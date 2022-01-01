package schema

import "entgo.io/ent"

// Seat holds the schema definition for the Seat entity.
type Seat struct {
	ent.Schema
}

// Fields of the Seat.
func (Seat) Fields() []ent.Field {
	return nil
}

// Edges of the Seat.
func (Seat) Edges() []ent.Edge {
	return nil
}
