// Package examples demonstrates Kubernetes fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Kubernetes example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Kubernetes — Basic Example ===\n")
	result := demonstrateKubernetes()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateKubernetes() string {
	return "Kubernetes basic demonstration"
}
