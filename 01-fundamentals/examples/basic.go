// Package examples demonstrates Go Fundamentals fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Go Fundamentals example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Go Fundamentals — Basic Example ===\n")
	result := demonstrateFundamentals()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateFundamentals() string {
	return "Go Fundamentals basic demonstration"
}
