package main

/////////////
// Main.go //
/////////////

import (
	"github.com/docopt/docopt-go"
)

const Version = "1.0.0"

var (
	usage = `Start a server to display the current usage of all the labs.

Usage:
  dashboard -v | --version
  dashboard -h | --help
  dashboard -b | --bind (<interface>:<port>|<interface>|:<port>)

Options:
  -v, --version  Show version
  -h, --help     Show this message
  -b, --bind     Set the interface:port for the server`

	defaultConfig    = "./static/config.json"
	defaultInterface = "localhost"
	defaultPort      = "8080"
)

func main() {
	_, _ = docopt.Parse(usage, nil, true, Version, false)

	Server()
}
