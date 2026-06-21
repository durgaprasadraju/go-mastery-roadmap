package solutions_test

import (
	"context"
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/122-grafana/exercises/solutions"
)

func TestServiceProcess(t *testing.T) {
	svc := &solutions.Service{Name: "Grafana"}
	out, err := svc.Process(context.Background(), "test")
	if err != nil || out == "" {
		t.Fatalf("out=%q err=%v", out, err)
	}
}

func TestServiceValidation(t *testing.T) {
	svc := &solutions.Service{Name: "Grafana"}
	_, err := svc.Process(context.Background(), "")
	if err == nil {
		t.Fatal("expected validation error")
	}
}
