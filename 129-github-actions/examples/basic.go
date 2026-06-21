// Package examples demonstrates GitHub Actions fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal GitHub Actions example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== GitHub Actions — Basic Example ===\n")
	result := demonstrateGithubActions()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGithubActions() string {
	return "GitHub Actions basic demonstration"
}
