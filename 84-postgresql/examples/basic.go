// Package examples demonstrates PostgreSQL fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal PostgreSQL example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== PostgreSQL — Basic Example ===\n")
	result := demonstratePostgresql()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstratePostgresql() string {
	return "PostgreSQL basic demonstration"
}
