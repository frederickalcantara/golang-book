package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countWords(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wordfreqL %v\n", err)
				continue
			}
			countWords(f, counts)
			f.Close()
		}
	}

	for word, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, word)
			fmt.Println()
		}
	}
}

func countWords(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		text := input.Text()
		counts[text]++
	}
}
