package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arxdsilva/go-release/release"
)

// TO DO: make it buildable by a file like .gorelease

func main() {
	var (
		path = flag.String("p", "", "Path to your app (required)")
		OSys = flag.String("os", "", "Systems to build your app (required)")
		// architecture = flag.String("a", "", "Kind of architecture wanted (required)")
	)
	flag.Parse()
	args := os.Args
	_, existARC := release.OS[*OSys]
	switch {
	case len(args) < 2 || args[1] == "-h":
		flag.Usage()
		os.Exit(1)
	case len(*OSys) == 0 || len(*path) == 0:
		flag.Usage()
		os.Exit(1)
	case !existARC:
		fmt.Println("Architecture not valid")
		flag.Usage()
		os.Exit(1)
	}
	Release(*path, *OSys)
}
