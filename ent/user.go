// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sundayonah/digital_post_office/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ClerkUserID holds the value of the "clerk_user_id" field.
	ClerkUserID string `json:"clerk_user_id,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// SentOrders holds the value of the sent_orders edge.
	SentOrders []*Order `json:"sent_orders,omitempty"`
	// ReceivedOrders holds the value of the received_orders edge.
	ReceivedOrders []*Order `json:"received_orders,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SentOrdersOrErr returns the SentOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SentOrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[0] {
		return e.SentOrders, nil
	}
	return nil, &NotLoadedError{edge: "sent_orders"}
}

// ReceivedOrdersOrErr returns the ReceivedOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReceivedOrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[1] {
		return e.ReceivedOrders, nil
	}
	return nil, &NotLoadedError{edge: "received_orders"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[2] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldClerkUserID, user.FieldFullName, user.FieldPhone, user.FieldEmail:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldClerkUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field clerk_user_id", values[i])
			} else if value.Valid {
				u.ClerkUserID = value.String
			}
		case user.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				u.FullName = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySentOrders queries the "sent_orders" edge of the User entity.
func (u *User) QuerySentOrders() *OrderQuery {
	return NewUserClient(u.config).QuerySentOrders(u)
}

// QueryReceivedOrders queries the "received_orders" edge of the User entity.
func (u *User) QueryReceivedOrders() *OrderQuery {
	return NewUserClient(u.config).QueryReceivedOrders(u)
}

// QueryNotifications queries the "notifications" edge of the User entity.
func (u *User) QueryNotifications() *NotificationQuery {
	return NewUserClient(u.config).QueryNotifications(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("clerk_user_id=")
	builder.WriteString(u.ClerkUserID)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(u.FullName)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
