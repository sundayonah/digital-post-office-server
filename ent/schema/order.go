// schema/order.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("tracking_number").Unique().NotEmpty(),
		field.String("safe_code").Sensitive().NotEmpty(),
		field.String("package_description").NotEmpty(),
		field.Enum("status").
			Values("pending", "delivered", "cancelled").
			Default("pending"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).
			Ref("sent_orders").
			Unique().
			Required(),
		edge.From("recipient", User.Type).
			Ref("received_orders").
			Unique().
			Required(),
		edge.To("notifications", Notification.Type),
	}
}
