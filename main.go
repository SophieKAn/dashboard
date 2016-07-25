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
	-h, --help
	-v, --version
  -b, --bind=<interface>:<port> Set the interface and port for the server
  --debug                       Turn on debugging output
  -c, --config=<file>           Specify a configuration file`

	defaultConfig    = "./static/config.json"
	defaultInterface = "localhost"
	defaultPort      = "8080"
	debug            = false
)

func main() {
	args, _ := docopt.Parse(usage, nil, true, Version, false)
	config       := configCommand(args["--config"])
	interf, port := bindCommand(args["--bind"])
	debug = args["--debug"].(bool)

	printStuff(interf, port, config, debug)
	Server(interf, port, config, debug)
}

func configCommand(filename interface{}) string {
	var config = ""
	if filename != nil {
		config = filename.(string)
	} else {
		config = defaultConfig
	}
	return config
}

func bindCommand(interfaceport interface{}) (string, string) {
	var interf, port string

	if interfaceport != nil {
		_ = interfaceport.(string)
	}

	interf = defaultInterface
	port   = defaultPort
	return interf, port
}

func printStuff(interf string, port string, config string, debug bool) {
	fmt.Printf("Interface is %s\n", interf)
	fmt.Printf("Port is %s\n", port)
	fmt.Printf("Config file is %s\n", config)
	fmt.Printf("Debug is %t\n", debug)
}
