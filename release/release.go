package release

import (
	"fmt"
	"os"
	"path/filepath"
)

// OS holds the architectures supported by each OS
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
	"all":       []string{"all"},
}

const (
	appName = "app"
	version = "0.1"
)

// Release generates a folder named "_releases" that contains all
// binaries for your requested app
func Release(path string, osys ...string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	if _, err := os.Stat("_release"); os.IsNotExist(err) {
		os.Mkdir("_release", 0777)
	}
	for _, Os := range osys {
		if _, existOS := OS[Os]; !existOS {
			// log error OS not supported
			continue
		}
		for _, ARC := range OS[Os] {
			// build binary
			//exec.CommandContext(ctx, name, arg)
			// move file using os.Rename
			// https://gist.github.com/johnzan/cbc3bf88a8c43b71112d
			err := os.Rename("appName", filepath.Join("_release", fmt.Sprintf("%s-%s.%s-%s.tar.gz", appName, version, Os, ARC)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
