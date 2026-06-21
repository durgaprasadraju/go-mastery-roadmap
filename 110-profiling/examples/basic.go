// Package examples demonstrates Profiling fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Profiling example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Profiling — Basic Example ===\n")
	result := demonstrateProfiling()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateProfiling() string {
	return "Profiling basic demonstration"
}
