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

	/* Set up framework for labs. */
	lab_frameworks := buildAll(labs)

	/* Update the statuses continually */
}

/* FUNCTIONS */

// buildAll takes the unmarshaled json and calls build on all the labs. It
// returns a list of Lab structs.
func buildAll(labs []interface{}) []Lab {
	all_labs := make([]Lab, 1)

	for lab := range labs {	
		one_lab := build(labs[lab].(map[string]interface{}))
		all_labs = append(all_labs, one_lab)
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

//
//
func updateStatuses(machines []*Machine) {
	var wg sync.WaitGroup
	for _, machine := range machines {
		wg.Add(1)

		go func(m *Machine) {
			defer wg.Done()
			//m.UpdateStatus()
		}(machine)
	}

	wg.Wait()
}
