// Package examples demonstrates Repository Pattern fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Repository Pattern example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Repository Pattern — Basic Example ===\n")
	result := demonstrateRepositoryPattern()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateRepositoryPattern() string {
	return "Repository Pattern basic demonstration"
}
