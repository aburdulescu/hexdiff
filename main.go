package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

const (
	cReset = "\033[0m"
	cRed   = "\033[31m"
)

func run() error {
	args := os.Args[1:]

	if len(args) < 1 {
		return fmt.Errorf("missing first hex string")
	}
	first := args[0]

	if len(args) < 2 {
		return fmt.Errorf("missing second hex string")
	}
	second := args[1]

	if len(first)%2 != 0 {
		return fmt.Errorf("first arg is not a hex string(len is not even)")
	}

	if len(second)%2 != 0 {
		return fmt.Errorf("second arg is not a hex string(len is not even)")
	}

	var smallestLen int
	if len(first) < len(second) {
		smallestLen = len(first)
	} else {
		smallestLen = len(second)
	}

	firstResult := ""
	secondResult := ""

	for i := 0; i < smallestLen; i += 2 {
		if first[i:i+2] != second[i:i+2] {
			firstResult += fmt.Sprint(cRed, first[i:i+2], cReset)
			secondResult += fmt.Sprint(cRed, second[i:i+2], cReset)
		} else {
			firstResult += fmt.Sprint(first[i : i+2])
			secondResult += fmt.Sprint(second[i : i+2])
		}
	}

	if len(first) > len(second) {
		firstResult += fmt.Sprint(cRed, first[len(second):], cReset)
	}

	if len(second) > len(first) {
		secondResult += fmt.Sprint(cRed, second[len(first):], cReset)
	}

	fmt.Println(firstResult)
	fmt.Println(secondResult)

	return nil
}
