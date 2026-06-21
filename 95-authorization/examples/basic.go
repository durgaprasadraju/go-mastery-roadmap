// Package examples demonstrates Authorization fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Authorization example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Authorization — Basic Example ===\n")
	result := demonstrateAuthorization()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateAuthorization() string {
	return "Authorization basic demonstration"
}
