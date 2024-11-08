// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sundayonah/digital_post_office/ent/order"
	"github.com/sundayonah/digital_post_office/ent/user"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TrackingNumber holds the value of the "tracking_number" field.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// SafeCode holds the value of the "safe_code" field.
	SafeCode string `json:"-"`
	// PackageDescription holds the value of the "package_description" field.
	PackageDescription string `json:"package_description,omitempty"`
	// Status holds the value of the "status" field.
	Status order.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges                OrderEdges `json:"edges"`
	user_sent_orders     *int
	user_received_orders *int
	selectValues         sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Sender holds the value of the sender edge.
	Sender *User `json:"sender,omitempty"`
	// Recipient holds the value of the recipient edge.
	Recipient *User `json:"recipient,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) SenderOrErr() (*User, error) {
	if e.Sender != nil {
		return e.Sender, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// RecipientOrErr returns the Recipient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) RecipientOrErr() (*User, error) {
	if e.Recipient != nil {
		return e.Recipient, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "recipient"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[2] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			values[i] = new(sql.NullInt64)
		case order.FieldTrackingNumber, order.FieldSafeCode, order.FieldPackageDescription, order.FieldStatus:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case order.ForeignKeys[0]: // user_sent_orders
			values[i] = new(sql.NullInt64)
		case order.ForeignKeys[1]: // user_received_orders
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldTrackingNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tracking_number", values[i])
			} else if value.Valid {
				o.TrackingNumber = value.String
			}
		case order.FieldSafeCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field safe_code", values[i])
			} else if value.Valid {
				o.SafeCode = value.String
			}
		case order.FieldPackageDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field package_description", values[i])
			} else if value.Valid {
				o.PackageDescription = value.String
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = order.Status(value.String)
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_sent_orders", value)
			} else if value.Valid {
				o.user_sent_orders = new(int)
				*o.user_sent_orders = int(value.Int64)
			}
		case order.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_received_orders", value)
			} else if value.Valid {
				o.user_received_orders = new(int)
				*o.user_received_orders = int(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QuerySender queries the "sender" edge of the Order entity.
func (o *Order) QuerySender() *UserQuery {
	return NewOrderClient(o.config).QuerySender(o)
}

// QueryRecipient queries the "recipient" edge of the Order entity.
func (o *Order) QueryRecipient() *UserQuery {
	return NewOrderClient(o.config).QueryRecipient(o)
}

// QueryNotifications queries the "notifications" edge of the Order entity.
func (o *Order) QueryNotifications() *NotificationQuery {
	return NewOrderClient(o.config).QueryNotifications(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("tracking_number=")
	builder.WriteString(o.TrackingNumber)
	builder.WriteString(", ")
	builder.WriteString("safe_code=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("package_description=")
	builder.WriteString(o.PackageDescription)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
