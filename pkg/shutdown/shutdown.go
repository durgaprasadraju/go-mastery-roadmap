// Package shutdown provides graceful shutdown coordination.
package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Graceful waits for SIGINT/SIGTERM then runs cleanup with timeout.
func Graceful(log *zap.Logger, timeout time.Duration, cleanup func(ctx context.Context) error) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	log.Info("shutdown signal received", zap.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := cleanup(ctx); err != nil {
		log.Error("shutdown cleanup failed", zap.Error(err))
		os.Exit(1)
	}
	log.Info("shutdown complete")
}
