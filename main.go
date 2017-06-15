package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/catkins/pfx/prefixer"
)

func main() {
	flag.Parse()

	prefix := flag.Arg(0)

	if prefix == "" {
		fmt.Fprintln(os.Stderr, "usage: pfx <PREFIX>")
		os.Exit(2)
	}
	scanner := bufio.NewScanner(os.Stdin)

	p := prefixer.NewPrefixer(prefix, scanner, os.Stdout)
	p.PrefixLines()
}
