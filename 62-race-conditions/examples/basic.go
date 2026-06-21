// Package examples demonstrates Race Conditions fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Race Conditions example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Race Conditions — Basic Example ===\n")
	result := demonstrateRaceConditions()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRaceConditions() string {
	return "Race Conditions basic demonstration"
}
