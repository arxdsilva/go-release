package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// ARC is all possible architectures avalible for golang
var ARC = map[string]struct{}{
	"386":      struct{}{},
	"amd64":    struct{}{},
	"arm":      struct{}{},
	"arm64":    struct{}{},
	"ppc64":    struct{}{},
	"ppc64le":  struct{}{},
	"mips64":   struct{}{},
	"mips64le": struct{}{},
}

// OS is
var OS = map[string][]string{
	"android":   []string{"arm"},
	"darwin":    []string{"386", "amd64", "arm", "arm64"},
	"dragonfly": []string{"amd64"},
	"freebsd":   []string{"386", "amd64", "arm"},
	"linux":     []string{"all"},
	"netbsd":    []string{"386", "amd64", "arm"},
	"openbsd":   []string{"386", "amd64", "arm"},
	"plan9":     []string{"386", "amd64"},
	"solaris":   []string{"amd64"},
	"windows":   []string{"386", "amd64"},
}

const (
	appName = "app"
	version = "0.1"
)

type AppInfo struct {
	Path         string
	Version      string
	OpSystem     string
	Architecture string
}

// Release generates a folder named "_releases" that contains all
// binaries for your requested app
func Release(app *AppInfo) error {
	// build binary with defer func
	// move file using os.Rename
	// https://gist.github.com/johnzan/cbc3bf88a8c43b71112d
	err := os.Rename("appName", filepath.Join("_release", fmt.Sprintf("%s-%s.%s-%s", appName, version, Os, arc)))
	if err != nil {
		return err
	}
	return nil
}
