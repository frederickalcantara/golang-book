// 1. Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount code below)

// 2. Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"fmt"
)

func differentBits(h1, h2 *[32]byte) int {
	counter := 0
	for i := range h1 {
		x := h1[i] ^ h2[i]
		for x != 0 {
			x &= x - 1
			counter++
		}
	}
	return counter
}

func main() {
	hash1 := "x"
	hash2 := "X"
	c1 := sha256.Sum256([]byte(hash1))
	c2 := sha256.Sum256([]byte(hash2))
	fmt.Printf("%x\n%x\n", c1, c2)
	fmt.Printf("Number of different bits: %d\n", differentBits(&c1, &c2))
}

/*
^ Bitwise excluse (OR) : Golang bitwise exclusive
If the bits are the same then you return 0
Else if the bits aren't the same then you return 1
*/
