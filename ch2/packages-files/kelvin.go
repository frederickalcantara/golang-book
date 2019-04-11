package main

import (
	"./tempconv"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			targ := scanner.Text()
			printTemps(targ)
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	} else {
		for _, args := range os.Args[1:] {
			printTemps(args)
		}
	}
}

func printTemps(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	k := tempconv.Kelvin(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", k, tempconv.KToC(k), c, tempconv.CToK(c))
}
