// Package examples demonstrates Clean Architecture fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Clean Architecture example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Clean Architecture — Basic Example ===\n")
	result := demonstrateCleanArchitecture()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCleanArchitecture() string {
	return "Clean Architecture basic demonstration"
}
