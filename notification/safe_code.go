package notification

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateSafeCode() string {
	// Generate a random 6-digit code
	code := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", code)
}

// Function to generate a random tracking number
func GenerateTrackingNumber() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	return fmt.Sprintf("TRACK-%d", rand.Intn(1000000))
}
