// Package examples demonstrates Interfaces fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Interfaces example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Interfaces — Basic Example ===\n")
	result := demonstrateInterfaces()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateInterfaces() string {
	return "Interfaces basic demonstration"
}
