package schema

import "entgo.io/ent"

// Airline holds the schema definition for the Airline entity.
type Airline struct {
	ent.Schema
}

// Fields of the Airline.
func (Airline) Fields() []ent.Field {
	return nil
}

// Edges of the Airline.
func (Airline) Edges() []ent.Edge {
	return nil
}
