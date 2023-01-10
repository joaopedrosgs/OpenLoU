package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Construction holds the schema definition for the Construction entity.
type Construction struct {
	ent.Schema
}

// Fields of the Construction.
func (Construction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x"),
		field.Int("y"),
		field.Int("raw_production").Default(0),
		field.Int("type").Default(0),
		field.Int("level").Default(0),
		field.Float("modifier").Default(1),
		field.Bool("need_refresh").Default(true),
	}
}

// Edges of the Construction.
func (Construction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("city", City.Type).Ref("constructions").Unique(),
		edge.To("queue", Queue.Type),
	}
}
func (Construction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("x", "y").
			Edges("city").
			Unique(),
	}
}
