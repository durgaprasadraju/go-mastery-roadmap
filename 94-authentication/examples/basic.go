// Package examples demonstrates Authentication fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Authentication example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Authentication — Basic Example ===\n")
	result := demonstrateAuthentication()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateAuthentication() string {
	return "Authentication basic demonstration"
}
