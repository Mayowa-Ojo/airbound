package schema

import "entgo.io/ent"

// FlightInstance holds the schema definition for the FlightInstance entity.
type FlightInstance struct {
	ent.Schema
}

// Fields of the FlightInstance.
func (FlightInstance) Fields() []ent.Field {
	return nil
}

// Edges of the FlightInstance.
func (FlightInstance) Edges() []ent.Edge {
	return nil
}
