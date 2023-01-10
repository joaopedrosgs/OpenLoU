package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.String("password_hash").NotEmpty(),
		field.Int("gold").Default(0),
		field.Int("diamonds").Default(0),
		field.Int("darkwood").Default(0),
		field.Int("runestone").Default(0),
		field.Int("veritium").Default(0),
		field.Int("trueseed").Default(0),
		field.Int("rank").Default(0),
		field.Int("alliance_rank").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cities", City.Type),
	}
}
