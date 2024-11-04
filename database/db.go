package database

import (
	"context"
	"fmt"

	"github.com/sundayonah/digital_post_office/ent"

	_ "github.com/lib/pq"
)

var Client *ent.Client

// InitDB initializes the database connection
func InitDB(ctx context.Context) (*ent.Client, error) {
	fmt.Println("Starting database connection")

	// postgresql://postgres:Encoded.001@localhost:5432/digital-post-office?sslmode=disable

	// Load environment variables
	// connectionString := os.Getenv("DATABASE_URL")
	connectionString := "postgresql://postgres:Encoded.001@localhost:5432/digital-post-office?sslmode=disable"
	if connectionString == "" {
		return nil, fmt.Errorf("database URL not found in environment variables")
	}

	// Open connection to PostgreSQL
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	// Set the global client
	Client = client

	fmt.Println("Successfully connected to the database")
	return client, nil
}
