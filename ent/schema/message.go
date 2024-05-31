package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int("board_id"),
		field.String("message"),
		field.String("writer"),
		field.Time("createdAt").
			Default(time.Now),
		field.Time("updatedAt").
			Default(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
