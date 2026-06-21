// Package examples demonstrates Generics fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Generics example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Generics — Basic Example ===\n")
	result := demonstrateGenerics()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGenerics() string {
	return "Generics basic demonstration"
}
