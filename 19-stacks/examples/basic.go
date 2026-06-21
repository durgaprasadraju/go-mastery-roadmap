// Package examples demonstrates Stacks fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Stacks example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Stacks — Basic Example ===\n")
	result := demonstrateStacks()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateStacks() string {
	return "Stacks basic demonstration"
}
