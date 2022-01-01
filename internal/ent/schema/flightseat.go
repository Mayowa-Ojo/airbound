package schema

import "entgo.io/ent"

// FlightSeat holds the schema definition for the FlightSeat entity.
type FlightSeat struct {
	ent.Schema
}

// Fields of the FlightSeat.
func (FlightSeat) Fields() []ent.Field {
	return nil
}

// Edges of the FlightSeat.
func (FlightSeat) Edges() []ent.Edge {
	return nil
}
