package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("firstname").MaxLen(250),
		field.String("lastname").MaxLen(250),
		field.String("email").MaxLen(250),
		field.String("phone").MaxLen(15),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Unique(),
		edge.To("admin", Admin.Type).
			Unique(),
		edge.To("crew", Crew.Type).
			Unique(),
		edge.To("pilot", Pilot.Type).
			Unique(),
		edge.To("front_desk", FrontDesk.Type).
			Unique(),
		edge.To("customer", Customer.Type).
			Unique(),
		edge.To("address", Address.Type).
			Unique(),
		edge.From("role", Role.Type).Ref("users").Unique(),
	}
}
