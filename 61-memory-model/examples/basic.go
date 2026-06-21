// Package examples demonstrates Go Memory Model fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Go Memory Model example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Go Memory Model — Basic Example ===\n")
	result := demonstrateMemoryModel()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMemoryModel() string {
	return "Go Memory Model basic demonstration"
}
