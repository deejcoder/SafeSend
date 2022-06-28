package schema

import (
	"SafeSend/pkg/interfaces"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type AccessToken struct {
	ent.Schema
}

func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.Enum("token_provider").
			GoType(interfaces.TokenProvider("")).
			Comment("The vendor that provided the token."),
		field.String("access_token").NotEmpty(),
		field.String("refresh_token").NotEmpty(),
		field.Time("expiry"),
		field.Time("date_created").Default(time.Now()),
		field.Time("date_modified").Default(time.Now()),
	}
}

func (AccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("access_tokens"),
	}
}
