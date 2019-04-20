// Example of an in-place slice algorithm

package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// works the same as the first nonempty function
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data)
}

//we can push a new value onto the end of the slice with append:
stack = append(stack, v) // push v

//The top of the stack is the last element:
top := stack[len(stack)-1] // top of stack

// and shrinking the stack by popping that element is
stack = stack[:len(stack)-1] // pop