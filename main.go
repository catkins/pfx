package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/catkins/pfx/prefixer"
)

func main() {
	var prefix string

	flag.StringVar(&prefix, "P", "", "string to prefix output with")
	flag.Parse()

	if prefix == "" {
		fmt.Fprintln(os.Stderr, "--prefix PREFIX required")
		os.Exit(2)
	}
	scanner := bufio.NewScanner(os.Stdin)

	p := prefixer.NewPrefixer(prefix, scanner, os.Stdout)
	p.PrefixLines()
}
