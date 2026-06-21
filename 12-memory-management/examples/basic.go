// Package examples demonstrates Memory Management fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Memory Management example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Memory Management — Basic Example ===\n")
	result := demonstrateMemoryManagement()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMemoryManagement() string {
	return "Memory Management basic demonstration"
}
