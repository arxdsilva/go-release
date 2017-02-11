package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		path    = flag.String("p", ".", "Path to your app")
		OSys    = flag.String("os", "all", "Systems to build your app")
		version = flag.String("v", "", "App version")
	)
	flag.Parse()
	args := os.Args
	_, existOS := OS[*OSys]
	switch {
	case args[1] == "-h":
		flag.Usage()
		os.Exit(1)
	case !existOS:
		fmt.Println("Architecture not valid")
		flag.Usage()
		os.Exit(1)
	}
	for osystem, arc := range OS {
		for _, a := range arc {
			app := &AppInfo{
				Path:         *path,
				Version:      *version,
				OpSystem:     osystem,
				Architecture: a,
			}
			Release(app)
		}
	}
}

func (app *AppInfo) createReleaseDir() error {
	err := os.Chdir(app.Path)
	if err != nil {
		return err
	}
	if _, err := os.Stat("_release"); os.IsNotExist(err) {
		os.Mkdir("_release", 0777)
	}
	return nil
}
