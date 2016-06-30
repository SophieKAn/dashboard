package main

/////////////
// Main.go //
/////////////

import (
	"net/http"
	"sync"
	"time"
)

// Machine represents a single machine in a lab.
type Machine struct {
	hostname string
	status   int
}

func main() {

	/* > Start the server and set up framework of labs from config.json */
	go http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))

	/* > Get lab configuration from config file. */
	labs := getConfig("./static/config.json")

	/* > Create a struct for each machines */
	all_machines := getMachines(labs)

	/* > Update the statuses continually */
	for {
		updateStatuses(all_machines)
		time.Sleep(5 * time.Minute)
	}
}

/* FUNCTIONS */

//
//
func getMachines(labs []interface{}) []*Machine {
	/* Make a list of Machine structs. */
	all_machines := make([]*Machine, 1)
	
	for lab := range labs {
		this_lab := labs[lab].(map[string]interface{})
		start := int(this_lab["start"].(float64))
		end := int(this_lab["end"].(float64))

		for i := start; i <= end; i++ {
			all_machines = append(all_machines, &Machine{"hostname", 1})
		}
	}

	return all_machines
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
