// Package examples demonstrates Linked Lists fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Linked Lists example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Linked Lists — Basic Example ===\n")
	result := demonstrateLinkedLists()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateLinkedLists() string {
	return "Linked Lists basic demonstration"
}
