package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Board holds the schema definition for the Board entity.
type Board struct {
	ent.Schema
}

// Fields of the Board.
func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.String("board_name").
			Default("unknown"),
		field.String("board_admin").
			Default("unknown"),
		field.String("board_password").
			Default("unknown"),
		field.Int("board_star").
			Default(0).
			Optional(),
		field.Time("createdAt").
			Default(time.Now),
		field.Time("updatedAt").
			Default(time.Now),
	}
}

// Edges of the Board.
func (Board) Edges() []ent.Edge {
	return nil
}
