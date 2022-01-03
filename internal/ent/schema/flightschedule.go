package schema

import (
	"airbound/internal/ent/customtypes"
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightSchedule holds the schema definition for the FlightSchedule entity.
type FlightSchedule struct {
	ent.Schema
}

// Fields of the FlightSchedule.
func (FlightSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Enum("weekday").GoType(enums.WeekDay(0)).Optional(),
		field.Enum("schedule_type").GoType(enums.FlightScheduleType("")),
		field.String("custom_date").GoType(customtypes.Date{}).SchemaType(map[string]string{dialect.Postgres: "date"}).Optional(),
		// temporarily storing this as 'string' till 'time without timezone' is added to the scanColumns method - ref: https://github.com/ent/ent/issues/2244
		field.String("departs_at").GoType(customtypes.Time{}),
		// field.String("departs_at").GoType(customtypes.Time{}).SchemaType(map[string]string{dialect.Postgres: "time"}),
		field.String("arrives_at").GoType(customtypes.Time{}),
		// field.String("arrives_at").GoType(customtypes.Time{}).SchemaType(map[string]string{dialect.Postgres: "time"}),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FlightSchedule.
func (FlightSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flight", Flight.Type).
			Ref("flight_schedules").
			Unique(),
	}
}
