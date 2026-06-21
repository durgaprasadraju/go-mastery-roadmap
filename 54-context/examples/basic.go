// Package examples demonstrates Context fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Context example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Context — Basic Example ===\n")
	result := demonstrateContext()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateContext() string {
	return "Context basic demonstration"
}
