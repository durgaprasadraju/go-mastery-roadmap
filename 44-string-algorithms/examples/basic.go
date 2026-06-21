// Package examples demonstrates String Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal String Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== String Algorithms — Basic Example ===\n")
	result := demonstrateStringAlgorithms()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateStringAlgorithms() string {
	return "String Algorithms basic demonstration"
}
