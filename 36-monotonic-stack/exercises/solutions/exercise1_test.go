package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/36-monotonic-stack/exercises/solutions"
)

func TestExercise1Core(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{"empty", nil, 0, true},
		{"single", []int{5}, 5, false},
		{"multiple", []int{1, 2, 3}, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solutions.Exercise1Core(tt.input)
			if tt.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestExercise1Transform(t *testing.T) {
	if got := solutions.Exercise1Transform("hello"); got != "olleh" {
		t.Fatalf("got %q", got)
	}
}
