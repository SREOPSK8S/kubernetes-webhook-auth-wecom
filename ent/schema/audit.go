package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Audit holds the schema definition for the Audit entity.
type Audit struct {
	ent.Schema
}

// Fields of the Audit.
func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.String("u_id"),
		field.String("m_id").Immutable(),
		field.Time("certification_time").Optional().SchemaType(map[string]string{dialect.MySQL: "datetime",}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Audit.
func (Audit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages",Message.Type),
	}
}