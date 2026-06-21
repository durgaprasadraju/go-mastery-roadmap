// Package examples demonstrates Replication fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Replication example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Replication — Basic Example ===\n")
	result := demonstrateReplication()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateReplication() string {
	return "Replication basic demonstration"
}
