// Package examples demonstrates Consensus Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Consensus Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Consensus Algorithms — Basic Example ===\n")
	result := demonstrateConsensus()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateConsensus() string {
	return "Consensus Algorithms basic demonstration"
}
