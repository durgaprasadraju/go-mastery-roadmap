// Package examples demonstrates Cheat Sheets fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Cheat Sheets example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Cheat Sheets — Basic Example ===\n")
	result := demonstrateCheatSheets()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateCheatSheets() string {
	return "Cheat Sheets basic demonstration"
}
