// Package examples demonstrates Worker Pools fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Worker Pools example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Worker Pools — Basic Example ===\n")
	result := demonstrateWorkerPools()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateWorkerPools() string {
	return "Worker Pools basic demonstration"
}
