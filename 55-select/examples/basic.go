// Package examples demonstrates Select Statement fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Select Statement example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Select Statement — Basic Example ===\n")
	result := demonstrateSelect()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSelect() string {
	return "Select Statement basic demonstration"
}
