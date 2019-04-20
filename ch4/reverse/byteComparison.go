// Equal function compares different types of slices
// bytes std library comes with a highly optimized bytes.Equal function for comparing two slices of bytes ([]byte)

// The only lefal slice comparison is against nil,
// if summer (sliced array) == nil { ... }

package main

import (
	"fmt"
)

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	equal()
}
