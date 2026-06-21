// Package examples demonstrates Structs fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Structs example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Structs — Basic Example ===\n")
	result := demonstrateStructs()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateStructs() string {
	return "Structs basic demonstration"
}
