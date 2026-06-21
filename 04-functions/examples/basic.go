// Package examples demonstrates Functions & Closures fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Functions & Closures example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Functions & Closures — Basic Example ===\n")
	result := demonstrateFunctions()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateFunctions() string {
	return "Functions & Closures basic demonstration"
}
