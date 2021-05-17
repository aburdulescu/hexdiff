package main

import (
	"fmt"
	"os"
	"strings"
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
		return fmt.Errorf("first arg is not a hex string(length is not even)")
	}

	if len(second)%2 != 0 {
		return fmt.Errorf("second arg is not a hex string(length is not even)")
	}

	firstResult, secondResult := hexdiff(first, second)

	fmt.Println(firstResult)
	fmt.Println(secondResult)

	return nil
}

func hexdiff(first, second string) (string, string) {
	var smallestLen int
	if len(first) < len(second) {
		smallestLen = len(first)
	} else {
		smallestLen = len(second)
	}

	var firstResult strings.Builder
	var secondResult strings.Builder

	firstResult.Grow(len(first))
	secondResult.Grow(len(second))

	for i := 0; i < smallestLen; i += 2 {
		if !strings.EqualFold(first[i:i+2], second[i:i+2]) {
			firstResult.WriteString(cRed)
			firstResult.WriteString(first[i : i+2])
			firstResult.WriteString(cReset)

			secondResult.WriteString(cRed)
			secondResult.WriteString(second[i : i+2])
			secondResult.WriteString(cReset)
		} else {
			firstResult.WriteString(first[i : i+2])
			secondResult.WriteString(second[i : i+2])
		}
	}

	if len(first) > len(second) {
		firstResult.WriteString(cRed)
		firstResult.WriteString(first[len(second):])
		firstResult.WriteString(cReset)
	}

	if len(second) > len(first) {
		secondResult.WriteString(cRed)
		secondResult.WriteString(second[len(first):])
		secondResult.WriteString(cReset)
	}

	return firstResult.String(), secondResult.String()
}
