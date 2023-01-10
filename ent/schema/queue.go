package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Queue holds the schema definition for the Queue entity.
type Queue struct {
	ent.Schema
}

// Fields of the Queue.
func (Queue) Fields() []ent.Field {
	return []ent.Field{
		field.Time("completion"),
		field.Int("action"),
		field.Int("order"),
	}
}

// Edges of the Queue.
func (Queue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("city", City.Type).Ref("queue").Unique(),
		edge.From("construction", Construction.Type).Ref("queue").Unique(),
	}
}

func (Queue) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order").
			Edges("city").
			Unique(),
	}
}
