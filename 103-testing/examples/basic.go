// Package examples demonstrates Testing Fundamentals fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Testing Fundamentals example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Testing Fundamentals — Basic Example ===\n")
	result := demonstrateTesting()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTesting() string {
	return "Testing Fundamentals basic demonstration"
}
