// Package examples demonstrates Grafana fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Grafana example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Grafana — Basic Example ===\n")
	result := demonstrateGrafana()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateGrafana() string {
	return "Grafana basic demonstration"
}
