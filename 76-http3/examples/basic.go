// Package examples demonstrates HTTP/3 fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal HTTP/3 example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== HTTP/3 — Basic Example ===\n")
	result := demonstrateHttp3()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateHttp3() string {
	return "HTTP/3 basic demonstration"
}
