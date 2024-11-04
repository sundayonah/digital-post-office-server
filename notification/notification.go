package notification

import (
	"context"
	"fmt"

	"github.com/sundayonah/digital_post_office/ent"
)

type NotificationService struct {
	client      *ent.Client
	smsProvider SMSProvider
}

// SMSProvider is an interface for sending SMS messages
type SMSProvider interface {
	SendSMS(phone, message string) error
}

func NewNotificationService(client *ent.Client, smsProvider SMSProvider) *NotificationService {
	return &NotificationService{
		client:      client,
		smsProvider: smsProvider,
	}
}

func (s *NotificationService) NotifyNewOrder(ctx context.Context, order *ent.Order, sender *ent.User, recipient *ent.User) error {
	// Create notification for sender
	senderMsg := fmt.Sprintf("Your package has been registered. Safe code: %s", order.SafeCode)
	if err := s.createNotification(ctx, sender, order, "new_order", senderMsg); err != nil {
		return err
	}

	// Send SMS to sender
	if err := s.smsProvider.SendSMS(sender.Phone, senderMsg); err != nil {
		return err
	}

	// Create notification for recipient
	recipientMsg := fmt.Sprintf("You have a package coming. Safe code: %s", order.SafeCode)
	if err := s.createNotification(ctx, recipient, order, "new_order", recipientMsg); err != nil {
		return err
	}

	// Send SMS to recipient
	if err := s.smsProvider.SendSMS(recipient.Phone, recipientMsg); err != nil {
		return err
	}

	return nil
}

func (s *NotificationService) createNotification(ctx context.Context, user *ent.User, order *ent.Order, nType, message string) error {
	_, err := s.client.Notification.Create().
		SetUser(user).
		SetOrder(order).
		SetType(nType).
		SetMessage(message).
		Save(ctx)
	return err
}
