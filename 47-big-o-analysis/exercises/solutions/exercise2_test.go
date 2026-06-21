package solutions_test

import (
	"context"
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/47-big-o-analysis/exercises/solutions"
)

func TestExercise2Process(t *testing.T) {
	svc := &solutions.Exercise2Service{Topic: "Big-O Analysis"}
	got, err := svc.Process(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}
	if got == "" {
		t.Fatal("expected non-empty result")
	}
}

func TestExercise2Validation(t *testing.T) {
	svc := &solutions.Exercise2Service{Topic: "Big-O Analysis"}
	_, err := svc.Process(context.Background(), "")
	if err == nil {
		t.Fatal("expected validation error")
	}
}
