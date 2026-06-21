// Package examples demonstrates Encryption fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Encryption example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Encryption — Basic Example ===\n")
	result := demonstrateEncryption()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateEncryption() string {
	return "Encryption basic demonstration"
}
