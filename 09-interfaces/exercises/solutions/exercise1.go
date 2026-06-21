package solutions

import "fmt"

// Person implements fmt.Stringer.
type Person struct{ First, Last string }

func (p Person) String() string {
	return fmt.Sprintf("%s, %s", p.Last, p.First)
}