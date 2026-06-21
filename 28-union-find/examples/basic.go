// Package examples demonstrates Union-Find fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Union-Find example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Union-Find — Basic Example ===\n")
	result := demonstrateUnionFind()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateUnionFind() string {
	return "Union-Find basic demonstration"
}
