// Package examples demonstrates Two Pointers fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Two Pointers example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Two Pointers — Basic Example ===\n")
	result := demonstrateTwoPointers()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateTwoPointers() string {
	return "Two Pointers basic demonstration"
}
