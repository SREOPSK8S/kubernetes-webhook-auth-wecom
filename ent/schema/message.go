package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("m_id").Unique().Comment("msg id"),
		field.String("content").Comment("msg content"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("msg create time"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner",Audit.Type).Ref("messages").Unique(),
	}
}
