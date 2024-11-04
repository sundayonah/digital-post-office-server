// schema/user.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("clerk_user_id").Unique().NotEmpty(),
		field.String("full_name").NotEmpty(),
		field.String("phone").Unique().NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sent_orders", Order.Type),
		edge.To("received_orders", Order.Type),
		edge.To("notifications", Notification.Type),
	}
}
