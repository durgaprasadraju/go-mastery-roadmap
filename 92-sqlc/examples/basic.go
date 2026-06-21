// Package examples demonstrates sqlc fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal sqlc example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== sqlc — Basic Example ===\n")
	result := demonstrateSqlc()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSqlc() string {
	return "sqlc basic demonstration"
}
