// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	isMatch := src > 0
	if !isMatch {
		os.Exit(0)
	}
	fmt.Println(src)
}

// match returns true if src matches pattern,
// false otherwise.
func match(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (s4et int, err error) {
	flag.Parse()
	s4et = len(flag.Args())
	if s4et == 0 {
		return 0, errors.New("missing string to world count")
	}
	return s4et, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}