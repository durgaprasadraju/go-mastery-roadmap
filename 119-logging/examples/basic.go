// Package examples demonstrates Logging fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Logging example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Logging — Basic Example ===\n")
	result := demonstrateLogging()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateLogging() string {
	return "Logging basic demonstration"
}
