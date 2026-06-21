package logger_test

import (
	"context"
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

func TestNew(t *testing.T) {
	l, err := logger.New("test-service")
	if err != nil {
		t.Fatal(err)
	}
	l.Info("test message")
}

func TestContextPropagation(t *testing.T) {
	l, err := logger.NewDevelopment("test")
	if err != nil {
		t.Fatal(err)
	}
	ctx := logger.WithContext(context.Background(), l)
	got := logger.FromContext(ctx)
	if got == nil {
		t.Fatal("expected logger from context")
	}
}
