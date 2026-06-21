// Package examples demonstrates Pipelines fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Pipelines example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Pipelines — Basic Example ===\n")
	result := demonstratePipelines()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstratePipelines() string {
	return "Pipelines basic demonstration"
}
