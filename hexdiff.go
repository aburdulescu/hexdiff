package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const usage = `Usage: hexdiff [options] input1 input2

Compare the two given inputs as hex and print their differences, if any.

Options:
    -h    print this message and exit
    -v    print version and exit
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
	fInputsAsFiles := flag.Bool("f", false, "treat inputs as files")
	fConvertToHex := flag.Bool("x", false, "when -i is active, convert file content to hex")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

	flag.Parse()

	_ = fVersion
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

	fmt.Print(hexdiff(first, second))

	return nil
}

func hexdiff(first, second string) string {
	var minLen int
	if len(first) < len(second) {
		minLen = len(first)
	} else {
		minLen = len(second)
	}

	result := new(strings.Builder)
	result.Grow(len(first) + len(second))

	i := 0
	for ; i < minLen; i += 2 {
		l := first[i : i+2]
		r := second[i : i+2]
		if !strings.EqualFold(l, r) {
			fmt.Fprintf(result, "%s%s %s%s\n", cRed, l, r, cReset)
		} else {
			fmt.Fprintf(result, "%s %s\n", l, r)
		}
	}

	if len(first) == len(second) {
		return result.String()
	}

	extra := ""
	padding := ""

	if len(first) > len(second) {
		extra = first[i:]
	} else {
		extra = second[i:]
		padding = "   "
	}

	for j := 0; j < len(extra); j += 2 {
		fmt.Fprintf(result, "%s%s%s%s\n", cRed, padding, extra[j:j+2], cReset)
	}

	return result.String()
}
