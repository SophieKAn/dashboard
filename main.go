package main

/////////////
// Main.go //
/////////////

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"regexp"
	"strings"
)

const Version = "dashboard 1.0.0"

var (
	usage = `Start a web server to display current usage of labs.

Usage:
  dashboard [options]
  dashboard --version
  dashboard -h | --help

Options:
  --debug                                             Turn on debugging output.
  -b, --bind=(<interface>:<port>|<interface>|:<port>) Set the interface and port for the server.
  -c, --config=<file>                                 Specify a configuration file.`

	defaultConfig    = "./static/config.json"
	defaultInterface = "localhost"
	defaultPort      = "8080"
	debug            = false
)

func main() {
	args, err := docopt.Parse(usage, nil, true, Version, false)
	Check(err)
	config := configCommand(args["--config"])
	interf, port := bindCommand(args["--bind"])
	debug = args["--debug"].(bool)

	PrintStuff(interf, port, config, debug)
	Server(interf, port, config, debug)
}

func configCommand(filename interface{}) string {
	config := defaultConfig

	if filename != nil {
		config = filename.(string)
	}

	return config
}

func bindCommand(input interface{}) (string, string) {
	interf, port := defaultInterface, defaultPort

	if input != nil {
		inputString := input.(string)

		if strings.Contains(inputString, ":") {
			rgx := regexp.MustCompile("(?P<interface>[a-zA-Z0-9.-]+)?:(?P<port>\\d{4})?")
			matches := rgx.FindStringSubmatch(inputString)
			matchMap := mapSubexpNames(matches, rgx.SubexpNames())

			if inf := matchMap["interface"]; inf != "" {
				interf = inf
			}

			if p := matchMap["port"]; p != "" {
				port = p
			}
		} else {
			interf = inputString
		}
	}
	return interf, port
}

func mapSubexpNames(m, n []string) map[string]string {
	/* http://stackoverflow.com/a/30483899/6279238 */
	/* Code found in comment on main answer */
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
