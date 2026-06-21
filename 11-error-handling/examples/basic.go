// Package examples demonstrates Error Handling fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Error Handling example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Error Handling — Basic Example ===\n")
	result := demonstrateErrorHandling()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateErrorHandling() string {
	return "Error Handling basic demonstration"
}
