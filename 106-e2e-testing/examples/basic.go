// Package examples demonstrates E2E Testing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal E2E Testing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== E2E Testing — Basic Example ===\n")
	result := demonstrateE2eTesting()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateE2eTesting() string {
	return "E2E Testing basic demonstration"
}
