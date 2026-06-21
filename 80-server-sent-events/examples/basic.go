// Package examples demonstrates Server-Sent Events fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Server-Sent Events example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Server-Sent Events — Basic Example ===\n")
	result := demonstrateServerSentEvents()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateServerSentEvents() string {
	return "Server-Sent Events basic demonstration"
}
