// Package examples demonstrates CQRS fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal CQRS example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== CQRS — Basic Example ===\n")
	result := demonstrateCqrs()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCqrs() string {
	return "CQRS basic demonstration"
}
