package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/01-fundamentals/exercises/solutions"
)

func BenchmarkCelsiusToFahrenheit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solutions.CelsiusToFahrenheit(100)
	}
}

func BenchmarkSumInts(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solutions.SumInts(data)
	}
}
