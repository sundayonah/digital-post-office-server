package notification

import "fmt"

// TwilioSMSProvider implements the SMSProvider interface
type TwilioSMSProvider struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

func NewTwilioSMSProvider(accountSID, authToken, fromNumber string) *TwilioSMSProvider {
	return &TwilioSMSProvider{
		AccountSID: accountSID,
		AuthToken:  authToken,
		FromNumber: fromNumber,
	}
}

func (t *TwilioSMSProvider) SendSMS(phone, message string) error {
	// TODO: Implement actual Twilio API call
	fmt.Printf("Sending SMS to %s: %s\n", phone, message)
	return nil
}

// For testing/development, you might want a mock provider
type MockSMSProvider struct{}

func NewMockSMSProvider() *MockSMSProvider {
	return &MockSMSProvider{}
}

func (m *MockSMSProvider) SendSMS(phone, message string) error {
	fmt.Printf("Mock SMS to %s: %s\n", phone, message)
	return nil
}
