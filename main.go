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
  dashboard -v | --version
  dashboard -h | --help
  dashboard -b | --bind (<interface>:<port>|<interface>|:<port>)
  dashboard --debug
	dashboard -c | --config <filename>

Options:
  -v, --version  Show version
  -h, --help     Show this message
  -b, --bind     Set the interface:port for the server
	--debug        Turn on debugging output
	-c, --config   Specify a configuration file`

	defaultConfig    = "./static/config.json"
	defaultInterface = "localhost"
	defaultPort      = "8080"
)

func main() {
	args, _ := docopt.Parse(usage, nil, true, Version, false)

	switch cmdName(args) {
	case "--bind":
		fmt.Println("Set the interface:port for the server")
	case "--debug":
		fmt.Println("Turn on debugging output")
	case "--config":
		fmt.Println("Specify a configuration file")
	}

	Server()
}


func cmdName(args map[string]interface{}) string {
	for _, cmd := range []string{"--bind", "--debug", "--config"} {
		if args[cmd].(bool) {
			return cmd
		}
	}

	return ""
}
