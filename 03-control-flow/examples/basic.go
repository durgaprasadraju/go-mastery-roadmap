// Package examples demonstrates Control Flow fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Control Flow example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Control Flow — Basic Example ===\n")
	result := demonstrateControlFlow()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateControlFlow() string {
	return "Control Flow basic demonstration"
}
