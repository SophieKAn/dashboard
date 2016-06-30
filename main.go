package main

import (
	"fmt"
	"net/http"
)

// Machine represents a single machine in a lab.
type Machine struct {
	hostname int
	status int
}

// Lab represents one lab. It contains a title and a list of Machines.
type Lab struct {
	name string
	machines []*Machine
}

func main() {

	/* Start the server and set up framework of labs from config.json */
	go http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))

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
