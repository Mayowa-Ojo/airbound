package schema

import (
	"regexp"
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
		field.String("email").MaxLen(250).Match(regexp.MustCompile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`)).Unique(),
		field.String("phone").MaxLen(15),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			StorageKey(edge.Column("user_id")).
			Unique(),
		edge.To("admin", Admin.Type).
			StorageKey(edge.Column("user_id")).
			Unique(),
		edge.To("crew", Crew.Type).
			StorageKey(edge.Column("user_id")).
			Unique(),
		edge.To("pilot", Pilot.Type).
			StorageKey(edge.Column("user_id")).
			Unique(),
		edge.To("front_desk", FrontDesk.Type).
			StorageKey(edge.Column("user_id")).
			Unique(),
		edge.To("customer", Customer.Type).
			StorageKey(edge.Column("user_id")).
			Unique(),
		edge.From("address", Address.Type).
			Ref("user").
			Unique(),
		edge.From("role", Role.Type).Ref("users").Unique(),
	}
}
