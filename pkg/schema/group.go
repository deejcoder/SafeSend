package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("groupName"),
		field.Int32("maxParticipants").
			Min(2).
			Max(10).
			Positive(),
		field.Bool("inviteOnly").Default(false),
		field.Time("dateCreated").Default(time.Now()),
		field.Time("dateModified").Default(time.Now()),
		field.Time("dateDeleted").Optional(),
	}
}

func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Unique(),
		edge.From("entities", Entity.Type).
			Ref("groups").
			Unique(),
	}
}
