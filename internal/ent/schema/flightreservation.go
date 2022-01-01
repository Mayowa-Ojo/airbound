package schema

import (
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightReservation holds the schema definition for the FlightReservation entity.
type FlightReservation struct {
	ent.Schema
}

// Fields of the FlightReservation.
func (FlightReservation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("reservation_number").MaxLen(50),
		field.Enum("reservation_status").GoType(enums.ReservationStatus("")),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FlightReservation.
func (FlightReservation) Edges() []ent.Edge {
	return nil
}
