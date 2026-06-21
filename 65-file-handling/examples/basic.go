// Package examples demonstrates File Handling fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal File Handling example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== File Handling — Basic Example ===\n")
	result := demonstrateFileHandling()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateFileHandling() string {
	return "File Handling basic demonstration"
}
