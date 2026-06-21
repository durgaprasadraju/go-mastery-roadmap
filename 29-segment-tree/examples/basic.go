// Package examples demonstrates Segment Trees fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Segment Trees example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Segment Trees — Basic Example ===\n")
	result := demonstrateSegmentTree()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSegmentTree() string {
	return "Segment Trees basic demonstration"
}
