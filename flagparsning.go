package phony

import (
	"flag"
	"fmt"
	"os"
)

func SetupAndParseFlags(config *string, port *int) {
	var help bool
	flag.Usage = printUsage

	// Flag declarations
	flag.BoolVar(&help, "help", false, "")
	flag.StringVar(config, "config", "./phonyconfig.json", "")
	flag.IntVar(port, "port", 9191, "")

	flag.Parse()

	if help == true {
		printUsage()
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Printf(`
PHONY -- A simple JSON server for serving predefined json responses.	
--------------------------------------------------------------------

USAGE:
    phony [options]

OPTIONS:
    --help:     Show usage and documentation.
    --config:   Path to phony config file (phonyconfig.json).
    --port:     Network port on which to run Phony server.
`)
}
