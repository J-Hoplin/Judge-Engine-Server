package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Name of the supported language"),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return nil
}
