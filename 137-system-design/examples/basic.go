// Package examples demonstrates System Design fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal System Design example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== System Design — Basic Example ===\n")
	result := demonstrateSystemDesign()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSystemDesign() string {
	return "System Design basic demonstration"
}
