package main

/////////////
// Main.go //
/////////////

import (
	"github.com/docopt/docopt-go"
	"regexp"
	"strings"
	"fmt"
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
	config := configCommand(args["--config"])
	interf, port := bindCommand(args["--bind"])
	debug = args["--debug"].(bool)

	PrintStuff(interf, port, config, debug)
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
	interf, port := defaultInterface, defaultPort

	if interfaceport != nil {
		str := interfaceport.(string)
		if strings.Contains(str, ":") {
			rgx := regexp.MustCompile("(?P<interface>[a-zA-Z0-9.-]+)?:(?P<port>\\d*)")
			match := rgx.FindStringSubmatch(str)
			names := rgx.SubexpNames()
			aMap := mapSubexpNames(match, names)
			if aMap["interface"] != "" {
				interf = aMap["interface"]
			}
			if aMap["port"] != "" {
				port = aMap["port"]
			}
		} else {
			interf = str
		}
	}

	return interf, port
}

func mapSubexpNames(m, n []string) map[string]string {
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i, _ := range n {
		r[n[i]] = m[i]
	}
	return r
}

func PrintStuff(intf string, port string, config string, debug bool) {
	fmt.Printf("Interface: %s\n", intf)
	fmt.Printf("Port:      %s\n", port)
	fmt.Printf("Config:    %s\n", config)
	fmt.Printf("Debug:     %t\n", debug)
	fmt.Println()
}
