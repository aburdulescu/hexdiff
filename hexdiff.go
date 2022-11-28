package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const usage = `Usage: chexdiff [options] input1 input2

Compare the two given inputs as hex and print their differences, if any.

Options:
    -h    print this message and exit
    -v    print version and exit
    -i    case insensitive comparison
    -f    treat inputs as files and compare their contents
    -x    if -f is active, convert file contents to hex before comparing them
`

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
	fVersion := flag.Bool("v", false, "version")
	fIgnoreCase := flag.Bool("i", false, "ignore case")
	fInputsAsFiles := flag.Bool("f", false, "treat inputs as files")
	fConvertToHex := flag.Bool("x", false, "when -i is active, convert file content to hex")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	flag.Parse()

	_ = fVersion
	_ = fIgnoreCase
	_ = fInputsAsFiles
	_ = fConvertToHex

	args := flag.Args()

	if len(args) < 2 {
		flag.Usage()
	}
	first := args[0]
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
