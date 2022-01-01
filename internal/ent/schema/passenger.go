package schema

import "entgo.io/ent"

// Passenger holds the schema definition for the Passenger entity.
type Passenger struct {
	ent.Schema
}

// Fields of the Passenger.
func (Passenger) Fields() []ent.Field {
	return nil
}

// Edges of the Passenger.
func (Passenger) Edges() []ent.Edge {
	return nil
}
