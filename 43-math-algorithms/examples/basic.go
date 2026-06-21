// Package examples demonstrates Math Algorithms fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Math Algorithms example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Math Algorithms — Basic Example ===\n")
	result := demonstrateMathAlgorithms()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMathAlgorithms() string {
	return "Math Algorithms basic demonstration"
}
