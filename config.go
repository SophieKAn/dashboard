package main

///////////////
// Config.go //
///////////////

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Configfile         string                   `json:"-"`
	Interface          string                   `json:"interface"`
	Port               string                   `json:"port"`
	Debug              bool                     `json:"-"`
	Interval           time.Duration            `json:"-"`
	MachineRanges      []map[string]interface{} `json:"machineRanges"`
	MachineIdentifiers []map[string]interface{} `json:"machineIdentifiers"`
}

const (
	linuxConfigPath   = "/etc/dashboard/config.json"
	freeBSDConfigPath = "/usr/local/etc/dashboard/config.json"

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

	version = "dashboard 1.0.0"
)

// Configure method on struct Config parses command-line arguments, environment
// variables, and the config file in that order, to glean settings for running
// the server.
func (c *Config) Configure() {
	parseArgs(c, getArgs())
	parseEnvs(c, getEnVars())
	parseConfig(c, c.Configfile)
}

// getArgs parses flags from the command line.
func getArgs() map[string]interface{} {
	args, err := docopt.Parse(usage, nil, true, version, false)
	check(err)
	return args
}

// parseArgs parses the command-line arguments and calls functions to interpret
// each one of them.
func parseArgs(c *Config, args map[string]interface{}) {
	c.Configfile = getConfigfile(args["--config"])
	c.Interface, c.Port = bindArg(args["--bind"])
	c.Interval = intervalArg(args["--interval"])
	c.Debug = args["--debug"].(bool)
}

// getConfigfile gets passed the name for the config file that was or wasn't
// given on the command line. If there wasn't one there it check environment
// variables, then default linux and freeBSD path. If it can't find a config
// file, the program will not proceed.
func getConfigfile(filename interface{}) string {
	var config string

	if filename != nil {
		config = filename.(string)
	} else if envConf := os.Getenv("DASHBOARD_CONFIG"); envConf != "" {
		config = envConf
	} else if _, err := os.Stat(linuxConfigPath); err == nil {
		config = linuxConfigPath
	} else if _, err := os.Stat(freeBSDConfigPath); err == nil {
		config = freeBSDConfigPath
	} else {
		fmt.Println("This program requires a config file to run. Please refer to documentation.")
		os.Exit(1)
	}

	return config
}

// bindArg parses the command-line argument for binding interface to port,
// then returns the acquired interface and port.
func bindArg(input interface{}) (string, string) {
	var interf, port string

	if input != nil {
		inputString := input.(string)
		interf, port = splitInterfacePort(inputString)
	}

	return interf, port
}

// intervalArg interpretes the command-line interval argument if there is one.
// If not, it returns zero.
func intervalArg(input interface{}) time.Duration {
	var interval time.Duration

	if input != nil {
		intervalString := input.(string)
		interval = getTimeInterval(intervalString)
	}

	return interval
}

// getEnvars creates a map of all the necessary environment variables for the
// program.
func getEnVars() map[string]string {
	envMap := make(map[string]string)

	envMap["BIND"] = os.Getenv("DASHBOARD_BIND")
	envMap["INTERVAL"] = os.Getenv("DASHBOARD_INTERVAL")
	envMap["DEBUG"] = os.Getenv("DASHBOARD_DEBUG")

	return envMap
}

// parseEnvs parses a map of relevant environment variables and uses the values
// if they aren't already present from the command line.
func parseEnvs(c *Config, enVars map[string]string) {
	i, p := splitInterfacePort(enVars["BIND"])
	if c.Interface == "" {
		c.Interface = i
	}

	if c.Port == "" {
		c.Port = p
	}

	dbg, err := strconv.ParseBool(enVars["DEBUG"])
	if err != nil {
		dbg = false
	}

	if c.Interval == 0 {
		c.Interval = getTimeInterval(enVars["INTERVAL"])
	}

	c.Debug = c.Debug || dbg
}

// parseConfig grabs the remaining settings from the config file, including the
// machine identifiers and ranges.
func parseConfig(c *Config, cfgFile string) {
	cfgfile := getConfig(cfgFile)

	if c.Interface == "" {
		c.Interface = cfgfile["interface"].(string)
	}
	if c.Port == "" {
		c.Port = cfgfile["port"].(string)
	}
	if c.Interval == 0 {
		c.Interval = getTimeInterval(cfgfile["interval"].(string))
	}

	c.MachineRanges = interfaceToList(cfgfile, "machineRanges")
	c.MachineIdentifiers = interfaceToList(cfgfile, "machineIdentifiers")
}
