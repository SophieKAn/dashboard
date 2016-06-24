package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// main opens and reads config.json into a MyJson type, and then proceeds to
// find the machine statuses for all the labs outlined in the config file.
func main() {

	/* Start the server set up framework of labs from config.json */
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))
	//check(err)

	/* Get initial status of every machine */
	labs := getConfig("./static/config.json")
	lab_statuses := buildAll(labs)

	/* update the status once */
	update(lab_statuses)
}

/* FUNCTIONS */

// buildAll calls build on each lab to obtain the initial status of all the
// machines.
func buildAll(labs []interface{}) map[string][]map[string]int {
	all_labs := make(map[string][]map[string]int)

	for lab := range labs {
		prefix, one_lab := build(labs[lab].(map[string]interface{}))
		all_labs[prefix] = one_lab
	}
	return all_labs
}

// build gets called on one lab. It calls systemStatus on each machine in the
// lab to obtain the initial statuses of all the machines.
func build(lab map[string]interface{}) (string, []map[string]int) {
	prefix := lab["prefix"].(string)
	start := int(lab["start"].(float64))
	end := int(lab["end"].(float64))
	machines_in_lab := make([]map[string]int, 1)

	for i := start; i <= end; i++ {
		hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)

		machine := make(map[string]int)
		machine["machine"], machine["status"] = i, systemStatus(hostname)
		machines_in_lab = append(machines_in_lab, machine)
	}
	return prefix, machines_in_lab
}

// getConfig takes the name of the configuration file (currently "config.json")
// attempts to open/read file then unmarshal it into a list of interfaces.
func getConfig(file_name string) []interface{} {
	config_file, err := ioutil.ReadFile(file_name)
	check(err)

	var labs []interface{}
	err = json.Unmarshal(config_file, &labs)
	check(err)
	return labs
}
