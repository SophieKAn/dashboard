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

// build builds one Lab and returns it, complete with a list of pointers to all
// of the machines in the lab.
func build(lab map[string]interface{}) Lab {
	var new_lab Lab
	new_lab.name = lab["prefix"].(string)

	machines_in_lab := make([]*Machine, 1)
	start := int(lab["start"].(float64))
	end := int(lab["end"].(float64))
	for i := start; i <= end; i++ {
		machines_in_lab = append(machines_in_lab, &Machine{i, 1})
	}
	new_lab.machines = machines_in_lab

	return new_lab
}
