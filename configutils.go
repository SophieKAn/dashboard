package main

////////////////////
// ConfigUtils.go //
////////////////////

// interfaceToList takes the config file and parses whichever group name it is
// given and expands it into a larger data structure.
func interfaceToList(cfgfile map[string]interface{}, name string) []map[string]interface{} {
	groupInterface := cfgfile[name].([]interface{})
	groupList := make([]map[string]interface{}, 0)
	for i := range groupInterface {
		lab := groupInterface[i].(map[string]interface{})
		groupList = append(groupList, lab)
	}
	return groupList
}

// splitInterfacePort takes a string taken from the command line or an
// environment variable and splits it using a regex.
func splitInterfacePort(inputString string) (string, string) {
	var interf, port string

	if strings.Contains(inputString, ":") {
		rgx := regexp.MustCompile("(?P<interface>[a-zA-Z0-9.-]+)?:(?P<port>\\d{4})?")
		matches := rgx.FindStringSubmatch(inputString)
		matchMap := mapSubexpNames(matches, rgx.SubexpNames())

		if i := matchMap["interface"]; i != "" {
			interf = i
		}

		if p := matchMap["port"]; p != "" {
			port = p
		}

	} else {
		interf = inputString
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

// getTimeInterval finds a time.Duration depending on the number and time units
// given on the command line or from environment variables.
func getTimeInterval(intervalString string) time.Duration {
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

// stringToTime takes the string representing the update interval and converts
// it into a time.Duration type.
func stringToTime(intervalString string, timeUnit string) time.Duration {
	durationString := strings.TrimSuffix(intervalString, timeUnit)
	duration, err := strconv.Atoi(durationString)
	check(err)

	return time.Duration(duration)
}
