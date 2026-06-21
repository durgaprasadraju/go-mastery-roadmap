// Package examples demonstrates OAuth2 fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal OAuth2 example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== OAuth2 — Basic Example ===\n")
	result := demonstrateOauth2()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateOauth2() string {
	return "OAuth2 basic demonstration"
}
