// Package examples demonstrates Database Migrations fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Database Migrations example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Database Migrations — Basic Example ===\n")
	result := demonstrateMigrations()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMigrations() string {
	return "Database Migrations basic demonstration"
}
