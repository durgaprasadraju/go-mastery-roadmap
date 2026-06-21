package solutions_test

import (
	"context"
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/43-math-algorithms/exercises/solutions"
)

func TestServiceProcess(t *testing.T) {
	svc := &solutions.Service{Name: "Math Algorithms"}
	out, err := svc.Process(context.Background(), "test")
	if err != nil || out == "" {
		t.Fatalf("out=%q err=%v", out, err)
	}
}

func TestServiceValidation(t *testing.T) {
	svc := &solutions.Service{Name: "Math Algorithms"}
	_, err := svc.Process(context.Background(), "")
	if err == nil {
		t.Fatal("expected validation error")
	}
}
