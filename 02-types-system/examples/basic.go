// Package examples demonstrates Go Type System fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Go Type System example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Go Type System — Basic Example ===\n")
	result := demonstrateTypesSystem()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTypesSystem() string {
	return "Go Type System basic demonstration"
}
