// Write an in-place function to eliminate adjacent duplicates in a []string slice.

package main

import "fmt"

func nonempty3(strings []string) []string {
	out := make([]string, 0) // zero-length slice or this will work as well: []string{}
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "one", "", "three", "one"}
	fmt.Printf("%q\n", nonempty3(data))
}
