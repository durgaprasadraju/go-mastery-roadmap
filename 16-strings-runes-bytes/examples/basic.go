// Package examples demonstrates Strings, Runes & Bytes fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Strings, Runes & Bytes example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Strings, Runes & Bytes — Basic Example ===\n")
	result := demonstrateStringsRunesBytes()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstrateStringsRunesBytes() string {
	return "Strings, Runes & Bytes basic demonstration"
}
