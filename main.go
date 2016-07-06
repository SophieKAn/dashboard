package main

/////////////
// Main.go //
/////////////

import (
	"fmt"
	"net/http"
	"time"
)

// A Machine represents one system.
type Machine struct {
	hostname string
	status   int
}

// main starts the server, gets all the lab info from config.json, sets up a
// channel for receiving updates, and then updates system statuses every 5 min.
func main() {

	/* > Start the server in its own Goroutine. */
	go http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))

	/* > Get lab configuration from config file. */
	lab_config := getConfig("./static/config.json")

	/* > Create a struct for each machine. */
	all_machines := getMachines(lab_config)

	/* > Create channel to receive status updates. */
	updates := make(chan *Machine)

	go func(updates chan *Machine) {
		for {
			fmt.Println(<-updates)
		}
	}(updates)

	/* > Update the statuses every second. */
	for {
		updateStatuses(all_machines, updates)
		time.Sleep(1 * time.Second)
	}
}

// getMachines takes the unmarshalled config.json and construct a slice of
// pointers to Machine structs representing all the machines in all the labs.
func getMachines(labs []interface{}) []*Machine {
	all_machines := make([]*Machine, 0)

	for lab := range labs {
		a_lab := labs[lab].(map[string]interface{})
		prefix := a_lab["prefix"].(string)
		start, end := int(a_lab["start"].(float64)), int(a_lab["end"].(float64))

		for i := start; i <= end; i++ {
			hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)
			all_machines = append(all_machines, &Machine{hostname, 2})
		}
	}
	return all_machines
}
