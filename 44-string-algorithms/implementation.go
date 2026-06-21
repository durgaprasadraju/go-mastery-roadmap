// Package string_algorithms implements String Algorithms from scratch for learning purposes.
package string_algorithms

import "errors"

var ErrEmpty = errors.New("string_algorithms: empty structure")

// DataStructure is a placeholder — see module-specific implementations.
type DataStructure struct {
	data []int
}

// New creates a new String Algorithms.
func New() *DataStructure {
	return &DataStructure{data: make([]int, 0)}
}

// Size returns the number of elements.
func (ds *DataStructure) Size() int {
	return len(ds.data)
}

// IsEmpty reports whether the structure has no elements.
func (ds *DataStructure) IsEmpty() bool {
	return len(ds.data) == 0
}
