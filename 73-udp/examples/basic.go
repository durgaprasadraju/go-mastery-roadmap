// Package examples demonstrates UDP fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal UDP example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== UDP — Basic Example ===\n")
	result := demonstrateUdp()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateUdp() string {
	return "UDP basic demonstration"
}
