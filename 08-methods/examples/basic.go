// Package examples demonstrates Methods fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Methods example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Methods — Basic Example ===\n")
	result := demonstrateMethods()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateMethods() string {
	return "Methods basic demonstration"
}
