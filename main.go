package main

/////////////
// Main.go //
/////////////

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

const Version = "1.0.0"

var (
	usage = `Start a server to display the current usage of all the labs.

Usage:
  dashboard [options]
	dashboard -v | --version
	dashboard -h | --help

Options:
  -b, --bind=<interface>:<port> Set the interface and port for the server
  --debug                       Turn on debugging output
  -c, --config=<file>           Specify a configuration file`

	defaultConfig    = "./static/config.json"
	defaultInterface = "localhost"
	defaultPort      = "8080"
)

func main() {
	args, _ := docopt.Parse(usage, nil, true, Version, false)
	for opt := range args {
		if args[opt] != false {
			fmt.Println("Option: ", opt)
		}

	}
}
