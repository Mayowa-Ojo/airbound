package schema

import "entgo.io/ent"

// FlightReservation holds the schema definition for the FlightReservation entity.
type FlightReservation struct {
	ent.Schema
}

// Fields of the FlightReservation.
func (FlightReservation) Fields() []ent.Field {
	return nil
}

// Edges of the FlightReservation.
func (FlightReservation) Edges() []ent.Edge {
	return nil
}
