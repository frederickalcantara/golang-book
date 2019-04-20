package main

import "fmt"

// remove elements while preserving the order
func removeOrderElements(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// remove elements without needing to preserve the order
func removeUnorderElements(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(removeOrderElements(s, 2))
	fmt.Println(removeUnorderElements(s, 2))
}
