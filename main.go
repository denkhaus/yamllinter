package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/denkhaus/tcgl/applog"
	"gopkg.in/yaml.v2"
)

//############################################################################
func pushError(f string, args ...interface{}) {
	applog.Errorf(fmt.Sprintf("yamllinter: %s", f), args...)
	os.Exit(1)
}

//############################################################################
func pushInfo(f string, args ...interface{}) {
	applog.Infof(fmt.Sprintf("yamllinter: %s", f), args...)
}

//############################################################################
func main() {

	if len(os.Args) < 2 {
		pushError("unrecognized input arguments: [usage: yamllinter <filetolint.yaml>]")
	}

	var filePath string
	verbose := false

	if os.Args[1] == "-v" {
		verbose = true
		filePath = os.Args[2]
	} else {
		filePath = os.Args[1]
	}

	if verbose {
		pushInfo("input file is %q", filePath)
	}

	if _, err := os.Stat(filePath); err != nil {
		pushError("input file not found")
	}

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		pushError("could not read input file")
	}

	t := make(map[string]interface{})
	if err := yaml.Unmarshal(buf, &t); err != nil {
		pushError("syntax error in %q: %s", filePath, err)

	}

	os.Exit(0)
}
