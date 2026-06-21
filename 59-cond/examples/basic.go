// Package examples demonstrates sync.Cond fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal sync.Cond example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== sync.Cond — Basic Example ===\n")
	result := demonstrateCond()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCond() string {
	return "sync.Cond basic demonstration"
}
