package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/122-grafana/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.FormatPrometheus("http_requests_total", 42); got != "http_requests_total 42.00" {
		t.Fatalf("got %v want %v", got, "http_requests_total 42.00")
	}
}
