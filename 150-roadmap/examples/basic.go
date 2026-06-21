// Package examples demonstrates Learning Roadmap fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Learning Roadmap example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Learning Roadmap — Basic Example ===\n")
	result := demonstrateRoadmap()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRoadmap() string {
	return "Learning Roadmap basic demonstration"
}
