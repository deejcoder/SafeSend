package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("email").Unique(),
		field.String("displayName"),
		field.Time("dateAccessed").Optional(),
		field.Time("dateCreated").Default(time.Now()),
		field.Time("dateModified").Default(time.Now()),
		field.Time("deletedDate").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("users"),
		edge.From("entities", Entity.Type).
			Ref("users").
			Unique(),
		edge.To("access_tokens", AccessToken.Type),
	}
}
