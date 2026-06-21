// Package examples demonstrates HTTP/2 fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal HTTP/2 example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== HTTP/2 — Basic Example ===\n")
	result := demonstrateHttp2()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHttp2() string {
	return "HTTP/2 basic demonstration"
}
