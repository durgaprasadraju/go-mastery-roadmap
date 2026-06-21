// Package examples demonstrates YAML fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal YAML example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== YAML — Basic Example ===\n")
	result := demonstrateYaml()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateYaml() string {
	return "YAML basic demonstration"
}
