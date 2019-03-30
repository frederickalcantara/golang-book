package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	printFunction()
	printRangeLoop()
	printLoop()
}

func printFunction() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Time it took to run join function: %.7fs ", time.Since(start).Seconds())
}

func printRangeLoop() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Time it took to run range function: %.7fs ", time.Since(start).Seconds())
}

func printLoop() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Time it took loop function: %.7fs ", time.Since(start).Seconds())
}
