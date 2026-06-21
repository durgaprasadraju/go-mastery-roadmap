// Package logger provides production-grade structured logging for all modules.
package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

// New creates a production-ready JSON logger.
func New(service string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.InitialFields = map[string]interface{}{
		"service": service,
	}
	return cfg.Build()
}

// NewDevelopment creates a human-readable development logger.
func NewDevelopment(service string) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.InitialFields = map[string]interface{}{
		"service": service,
	}
	return cfg.Build()
}

// Must creates a logger or panics.
func Must(service string) *zap.Logger {
	l, err := New(service)
	if err != nil {
		panic(err)
	}
	return l
}

// WithContext attaches a logger to context for propagation.
func WithContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, l)
}

// FromContext retrieves the logger from context or returns a default.
func FromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok && l != nil {
		return l
	}
	l, _ := NewDevelopment("default")
	return l
}

// Default returns a logger writing to stderr.
func Default() *zap.Logger {
	l, err := NewDevelopment(os.Getenv("SERVICE_NAME"))
	if err != nil {
		panic(err)
	}
	return l
}
