// Package examples demonstrates Production Projects fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Production Projects example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Production Projects — Basic Example ===\n")
	result := demonstrateProductionProjects()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateProductionProjects() string {
	return "Production Projects basic demonstration"
}
