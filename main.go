package main

/////////////
// Main.go //
/////////////

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const Version = "dashboard 1.0.0"

type Configs struct {
	Configfile string
	Interface  string
	Port       string
	Interval   time.Duration
	Debug      bool
}

var (
	usage = `Start a web server to display current usage of labs.

Usage:
  dashboard [options]
  dashboard --version
  dashboard -h | --help

Options:
  --debug                                             Turn on debugging output.
  -b, --bind=(<interface>:<port>|<interface>|:<port>) Set the interface and port for the server.
  -c, --config=<file>                                 Specify a configuration file.
  -i, --interval=(<sec>s|<min>m|<hr>h)`

	defaultConfig     = "./static/config.json"
	defaultInterval   = 5 * time.Second
	linuxConfigPath   = "/etc/dashboard/config.json"
	freeBSDConfigPath = "/usr/local/etc/dashboard/config.json"
)

func main() {

	var configs Configs = Configs{"", "", "", 0, false}

	/* > Get configs from command-line arguments */

	args, err := docopt.Parse(usage, nil, true, Version, false)
	Check(err)

	configs.Configfile = configCommand(args["--config"])
	configs.Interface, configs.Port = bindCommand(args["--bind"])
	configs.Interval = intervalCommand(args["--interval"])
	configs.Debug = args["--debug"].(bool)

	/* > Get configs from config file */

	PrintConfigs(configs)
	//Server(configs)
}

func configCommand(filename interface{}) string {
	var config string
	environmentConfig := os.Getenv("DASHBOARD_CONFIG")

	if filename != nil {
		config = filename.(string)
	} else if environmentConfig != "" {
		config = os.Getenv("DASHBOARD_CONFIG")
	} else if _, err := os.Stat(linuxConfigPath); err == nil {
		config = linuxConfigPath
	} else if _, err := os.Stat(freeBSDConfigPath); err == nil {
		config = freeBSDConfigPath
	} else {
		fmt.Println("This program requires a config file to run. See documentation.")
	}

	return config
}

func bindCommand(input interface{}) (string, string) {
	var interf, port, inputString string

	if input != nil {
		inputString = input.(string)
	} else {
		inputString = os.Getenv("DASHBOARD_BIND")
	}

	if inputString != "" {
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

func intervalCommand(input interface{}) time.Duration {
	var intervalString string
	var interval time.Duration

	if input != nil {
		intervalString = input.(string)
	} else {
		intervalString = os.Getenv("DASHBOARD_INTERVAL")
	}

	if intervalString != "" {
		if strings.Contains(intervalString, "s") {
			number := strings.TrimSuffix(intervalString, "s")
			theTime, err := strconv.Atoi(number)
			Check(err)
			interval = time.Second * time.Duration(theTime)
		}

		if strings.Contains(intervalString, "m") {
			number := strings.TrimSuffix(intervalString, "m")
			theTime, err := strconv.Atoi(number)
			Check(err)
			interval = time.Minute * time.Duration(theTime)
		}

		if strings.Contains(intervalString, "h") {
			number := strings.TrimSuffix(intervalString, "h")
			theTime, err := strconv.Atoi(number)
			Check(err)
			interval = time.Hour * time.Duration(theTime)

		}
	}
	return interval
}

func PrintConfigs(configs Configs) {
	fmt.Printf("Interface:  %s\n", configs.Interface)
	fmt.Printf("Port:       %s\n", configs.Port)
	fmt.Printf("Configfile: %s\n", configs.Configfile)
	fmt.Printf("Interval:   %q\n", configs.Interval)
	fmt.Printf("Debug:      %t\n", configs.Debug)
	fmt.Println()
}
