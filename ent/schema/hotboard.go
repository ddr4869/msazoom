package schema

import "entgo.io/ent"

// HotBoard holds the schema definition for the HotBoard entity.
type HotBoard struct {
	ent.Schema
}

// Fields of the HotBoard.
func (HotBoard) Fields() []ent.Field {
	return nil
}

// Edges of the HotBoard.
func (HotBoard) Edges() []ent.Edge {
	return nil
}
