// Package examples demonstrates Priority Queues fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal Priority Queues example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== Priority Queues — Basic Example ===\n")
	result := demonstratePriorityQueue()
	fmt.Printf("Result: %v\n", result)
	return nil
}

func demonstratePriorityQueue() string {
	return "Priority Queues basic demonstration"
}
