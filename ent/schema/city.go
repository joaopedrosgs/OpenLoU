package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x").Default(0),
		field.Int("y").Default(0),
		field.String("name").NotEmpty().Default("New city"),
		field.Int("points").Default(3),

		field.Int("wood_production").Default(300),
		field.Int("stone_production").Default(0),
		field.Int("iron_production").Default(0),
		field.Int("food_production").Default(0),

		field.Int("wood_stored").Default(300),
		field.Int("stone_stored").Default(0),
		field.Int("iron_stored").Default(0),
		field.Int("food_stored").Default(0),

		field.Int("wood_limit").Default(300),
		field.Int("stone_limit").Default(0),
		field.Int("iron_limit").Default(0),
		field.Int("food_limit").Default(0),

		field.Time("queue_time").Default(time.Now),
		field.Int("construction_speed").Default(1),
	}
}
func (City) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("x", "y").
			Unique(),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Unique().Ref("cities"),
		edge.To("constructions", Construction.Type),
		edge.To("queue", Queue.Type),
	}
}
