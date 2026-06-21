// Package examples demonstrates Searching Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Searching Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Searching Algorithms — Basic Example ===\n")
	result := demonstrateSearching()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSearching() string {
	return "Searching Algorithms basic demonstration"
}
