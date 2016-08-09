package main

///////////////
// Config.go //
///////////////

import (
	"time"
	"os"
	"strings"
	"strconv"
	"fmt"
	"github.com/docopt/docopt-go"
	"regexp"
)

type Config struct {
	Configfile string
	Interface  string
	Port       string
	Debug      bool
	Interval   time.Duration
}

const (
	defaultConfig     = "./static/config.json"
	linuxConfigPath   = "/etc/dashboard/config.json"
	freeBSDConfigPath = "usr/local/etc/dashboard/config.json"
)

//
//
func (c *Config) Configure() {
	parseArgs(c, getArgs())
	parseEnvs(c, getEnVars())
	parseConfig(c, c.Configfile)
}

//
//
func parseArgs(c *Config, args map[string]interface{}) {
	c.Configfile = configCommand(args["--config"])
	c.Interface, c.Port = bindCommand(args["--bind"])
	c.Interval = intervalCommand(args["--interval"])
	c.Debug = args["--debug"].(bool)
}

//
//
func parseEnvs(c *Config, enVars map[string]string) {
	if c.Interface == "" || c.Port == "" {
		c.Port, c.Interface = splitInterfacePort(enVars["DASHBOARD_BIND"])
	}
	dbg, err := strconv.ParseBool(enVars["DASHBOARD_DEBUG"]); Check(err)

	c.Debug = c.Debug || dbg

	if c.Interval == 0 {
		c.Interval = getInterval(enVars["DASHBOARD_INTERVAL"])
	}
}

//
//
func parseConfig(c *Config, cfgFile string) {
	cfgfile := GetConfig(cfgFile)

	if c.Interface == "" {
		c.Interface = cfgfile["interface"].(string)
	}
	if c.Port == "" {
		c.Port = cfgfile["port"].(string)
	}
	if c.Interval == 0 {
		c.Interval = getInterval(cfgfile["interval"].(string))
	}
}

//
//
func getArgs() map[string]interface{} {
	args, err := docopt.Parse(Usage, nil, true, Version, false); Check(err)

	return args
}

//
//
func getEnVars() map[string]string {
	envMap := make(map[string]string)

	envMap["DASHBOARD_BIND"] = os.Getenv("DASHBOARD_BIND")
	envMap["DASHBOARD_INTERVAL"] = os.Getenv("DASHBOARD_INTERVAL")
	envMap["DASHBOARD_DEBUG"] = os.Getenv("DASHBOARD_DEBUG")

	return envMap
}

//
//
func configCommand(filename interface{}) string {
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
		fmt.Println("This program requires a config file to run. See documentation")
		os.Exit(1)
	}

	return config
}

//
//
func bindCommand(input interface{}) (string, string) {
	var interf, port string

	if input != nil {
		inputString := input.(string)
		interf, port = splitInterfacePort(inputString)
	}

	return interf, port
}

//
//
func intervalCommand(input interface{}) time.Duration {
	var interval time.Duration

	if input != nil {
		intervalString := input.(string)
		interval = getInterval(intervalString)
	}

	return interval
}

//
//
func splitInterfacePort(inputString string) (string, string) {
	var intf, prt string

	if strings.Contains(inputString, ":") {
		rgx := regexp.MustCompile("(?P<interface>[a-zA-Z0-9.-]+)?:(?P<port>\\d{4})?")
		matches := rgx.FindStringSubmatch(inputString)
		matchMap := mapSubexpNames(matches, rgx.SubexpNames())

		if i := matchMap["interface"]; i != "" {
			intf = i
		}

		if p := matchMap["port"]; p != "" {
			prt = p
		}

	} else {
			intf = inputString
	}
	return intf, prt
}

//
//
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

//
//
func getInterval(intervalString string) time.Duration {
	var interval time.Duration

	if strings.Contains(intervalString, "s") {
		interval = time.Second * stringToTime(intervalString, "s")

	} else if strings.Contains(intervalString, "m") {
		interval = time.Minute * stringToTime(intervalString, "m")

	} else if strings.Contains(intervalString, "h") {
		interval = time.Hour * stringToTime(intervalString, "h")
	}

	return interval
}

//
//
func stringToTime(intervalString string, timeUnit string) time.Duration {
	number := strings.TrimSuffix(intervalString, timeUnit)
	theTime, err := strconv.Atoi(number); Check(err)

	return time.Duration(theTime)
}
