// Package examples demonstrates Secrets Management fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Secrets Management example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Secrets Management — Basic Example ===\n")
	result := demonstrateSecretsManagement()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateSecretsManagement() string {
	return "Secrets Management basic demonstration"
}
