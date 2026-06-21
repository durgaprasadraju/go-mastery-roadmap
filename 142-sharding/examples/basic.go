// Package examples demonstrates Sharding fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Sharding example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Sharding — Basic Example ===\n")
	result := demonstrateSharding()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSharding() string {
	return "Sharding basic demonstration"
}
