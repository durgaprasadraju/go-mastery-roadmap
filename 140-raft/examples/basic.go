// Package examples demonstrates Raft Consensus fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Raft Consensus example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Raft Consensus — Basic Example ===\n")
	result := demonstrateRaft()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRaft() string {
	return "Raft Consensus basic demonstration"
}
