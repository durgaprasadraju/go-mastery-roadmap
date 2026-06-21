// Package examples demonstrates Fast & Slow Pointers fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Fast & Slow Pointers example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Fast & Slow Pointers — Basic Example ===\n")
	result := demonstrateFastSlowPointers()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateFastSlowPointers() string {
	return "Fast & Slow Pointers basic demonstration"
}
