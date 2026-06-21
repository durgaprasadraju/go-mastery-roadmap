// Package examples demonstrates Binary Search Trees fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Binary Search Trees example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Binary Search Trees — Basic Example ===\n")
	result := demonstrateBinarySearchTree()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateBinarySearchTree() string {
	return "Binary Search Trees basic demonstration"
}
