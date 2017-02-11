package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arxdsilva/go-release/release"
)

// TO DO: make it buildable by a file like .gorelease

// var ARC = map[string]struct{}{
// 	"386":      struct{}{},
// 	"amd64":    struct{}{},
// 	"arm":      struct{}{},
// 	"arm64":    struct{}{},
// 	"ppc64":    struct{}{},
// 	"ppc64le":  struct{}{},
// 	"mips64":   struct{}{},
// 	"mips64le": struct{}{},
// 	"all":      struct{}{},
// }

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
