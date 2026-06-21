// Package examples demonstrates Unit Testing fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Unit Testing example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Unit Testing — Basic Example ===\n")
	result := demonstrateUnitTesting()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateUnitTesting() string {
	return "Unit Testing basic demonstration"
}
