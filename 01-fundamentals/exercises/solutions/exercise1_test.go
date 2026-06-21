package solutions_test

import (
	"math"
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/01-fundamentals/exercises/solutions"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		c    float64
		want float64
		err  bool
	}{
		{0, 32, false},
		{100, 212, false},
		{-273.15, -459.67, false},
		{-300, 0, true},
	}
	for _, tt := range tests {
		got, err := solutions.CelsiusToFahrenheit(tt.c)
		if tt.err {
			if err == nil {
				t.Fatalf("expected error for %v", tt.c)
			}
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		if math.Abs(got-tt.want) > 0.01 {
			t.Fatalf("CelsiusToFahrenheit(%v) = %v, want %v", tt.c, got, tt.want)
		}
	}
}

func TestHTTPStatus(t *testing.T) {
	if solutions.StatusOKExplicit.String() != "OK" {
		t.Fatal("unexpected status string")
	}
	if !solutions.StatusOKExplicit.IsSuccess() {
		t.Fatal("200 should be success")
	}
	if solutions.StatusNotFound.IsSuccess() {
		t.Fatal("404 should not be success")
	}
}

func TestSumInts(t *testing.T) {
	sum, err := solutions.SumInts([]int{1, 2, 3})
	if err != nil || sum != 6 {
		t.Fatalf("sum=%d err=%v", sum, err)
	}
}
