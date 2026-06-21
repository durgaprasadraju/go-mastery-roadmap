// Package examples demonstrates Helm fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Helm example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Helm — Basic Example ===\n")
	result := demonstrateHelm()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHelm() string {
	return "Helm basic demonstration"
}
