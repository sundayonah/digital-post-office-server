package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("message").NotEmpty(),
		field.String("type").NotEmpty(),
		field.Bool("is_read").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("notifications").
			Unique().
			Required(),
		edge.From("order", Order.Type).
			Ref("notifications").
			Unique().
			Required(),
	}
}
