package solutions

import "fmt"

// MetricPoint represents a single metric sample.
type MetricPoint struct {
	Name  string
	Value float64
	Tags  map[string]string
}

// FormatPrometheus formats a counter metric line.
func FormatPrometheus(name string, value float64) string {
	return fmt.Sprintf("%s %.2f", name, value)
}