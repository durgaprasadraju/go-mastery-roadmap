// Package examples demonstrates Integration Testing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Integration Testing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Integration Testing — Basic Example ===\n")
	result := demonstrateIntegrationTesting()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateIntegrationTesting() string {
	return "Integration Testing basic demonstration"
}
