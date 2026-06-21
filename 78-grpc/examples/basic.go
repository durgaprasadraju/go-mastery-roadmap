// Package examples demonstrates gRPC fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal gRPC example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== gRPC — Basic Example ===\n")
	result := demonstrateGrpc()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGrpc() string {
	return "gRPC basic demonstration"
}
