// 1. Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount code below)

// 2. Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func differentHash(data string) {

}

func main() {
	hash1 := "x"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "sha384":
			c1 := sha512.Sum384([]byte(hash1))
			fmt.Printf("%x\n", c1)
		case "sha512":
			c1 := sha512.Sum512([]byte(hash1))
			fmt.Printf("%x\n", c1)
		default:
			c1 := sha256.Sum256([]byte(hash1))
			fmt.Printf("%x\n", c1)
		}
	}
}
