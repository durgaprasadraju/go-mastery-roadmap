// Package examples demonstrates Tree Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Tree Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Tree Algorithms — Basic Example ===\n")
	result := demonstrateTreeAlgorithms()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTreeAlgorithms() string {
	return "Tree Algorithms basic demonstration"
}
