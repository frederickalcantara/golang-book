// Write a version of rotate that operates in a single pass.

package main

import "fmt"

func main() {
	r := 1
	data := []string{"Hello", "people", "world"}
	data = append(data[r:], data[:r]...)
	fmt.Println(data)
}
