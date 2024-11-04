// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/sundayonah/digital_post_office/ent/notification"
	"github.com/sundayonah/digital_post_office/ent/order"
	"github.com/sundayonah/digital_post_office/ent/schema"
	"github.com/sundayonah/digital_post_office/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescMessage is the schema descriptor for message field.
	notificationDescMessage := notificationFields[0].Descriptor()
	// notification.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	notification.MessageValidator = notificationDescMessage.Validators[0].(func(string) error)
	// notificationDescType is the schema descriptor for type field.
	notificationDescType := notificationFields[1].Descriptor()
	// notification.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	notification.TypeValidator = notificationDescType.Validators[0].(func(string) error)
	// notificationDescIsRead is the schema descriptor for is_read field.
	notificationDescIsRead := notificationFields[2].Descriptor()
	// notification.DefaultIsRead holds the default value on creation for the is_read field.
	notification.DefaultIsRead = notificationDescIsRead.Default.(bool)
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationFields[3].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescTrackingNumber is the schema descriptor for tracking_number field.
	orderDescTrackingNumber := orderFields[0].Descriptor()
	// order.TrackingNumberValidator is a validator for the "tracking_number" field. It is called by the builders before save.
	order.TrackingNumberValidator = orderDescTrackingNumber.Validators[0].(func(string) error)
	// orderDescSafeCode is the schema descriptor for safe_code field.
	orderDescSafeCode := orderFields[1].Descriptor()
	// order.SafeCodeValidator is a validator for the "safe_code" field. It is called by the builders before save.
	order.SafeCodeValidator = orderDescSafeCode.Validators[0].(func(string) error)
	// orderDescPackageDescription is the schema descriptor for package_description field.
	orderDescPackageDescription := orderFields[2].Descriptor()
	// order.PackageDescriptionValidator is a validator for the "package_description" field. It is called by the builders before save.
	order.PackageDescriptionValidator = orderDescPackageDescription.Validators[0].(func(string) error)
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderFields[4].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderFields[5].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFullName is the schema descriptor for full_name field.
	userDescFullName := userFields[0].Descriptor()
	// user.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	user.FullNameValidator = userDescFullName.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[1].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
