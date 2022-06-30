package schema

import (
	"SafeSend/pkg/interfaces"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ent.Schema
}

func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.Enum("entity_type").
			GoType(interfaces.EntityType("")).
			Comment("Based on the entity type we can determine how to send messages"),
		field.Time("date_created").Default(time.Now()),
		field.Time("date_modified").Default(time.Now()),
		field.Time("date_deleted").Optional(),
	}
}

func (Entity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("groups", Group.Type),
	}
}
