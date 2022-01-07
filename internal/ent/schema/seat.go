package schema

import (
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Seat holds the schema definition for the Seat entity.
type Seat struct {
	ent.Schema
}

// Fields of the Seat.
func (Seat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Int("seat_number").NonNegative(),
		field.String("seat_row").MaxLen(1),
		field.Enum("seat_type").GoType(enums.SeatType("")),
		field.Enum("seat_class").GoType(enums.SeatClass("")),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Seat.
func (Seat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("aircraft", Aircraft.Type).
			Ref("seats").
			Unique(),
		edge.To("flight_seat", FlightSeat.Type).
			StorageKey(edge.Column("seat_id")).
			Unique(),
	}
}
