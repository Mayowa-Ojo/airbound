package schema

import (
	"airbound/internal/ent/enums"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Enum("account_status").GoType(enums.AccountStatus("")),
		field.Bytes("password"),
		field.Bytes("salt"),
		field.String("two_fa_secret").MaxLen(250).Optional(),
		field.Bool("two_fa_completed").Default(false),
		field.String("verification_token").Optional(),
		field.String("forgot_password_token").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("account").
			Unique(),
	}
}
