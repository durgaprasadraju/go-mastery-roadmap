// Package examples demonstrates WebSocket fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal WebSocket example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== WebSocket — Basic Example ===\n")
	result := demonstrateWebsocket()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateWebsocket() string {
	return "WebSocket basic demonstration"
}
