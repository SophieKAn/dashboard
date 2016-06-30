package main

/////////////
// Main.go //
/////////////

import (
	"fmt"
	"net/http"
	"time"
)

// Machine represents a single machine in a lab.
type Machine struct {
	hostname string
	status   int
}

func main() {
	/* > Start the server in its own Goroutine. */
	go http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))

	/* > Get lab configuration from config file. */
	labs := getConfig("./static/config.json")

	/* > Create a struct for each machine. */
	all_machines := getMachines(labs)


	updates := make(chan *Machine)
	go func (updates chan *Machine) {
		for {
			select {
				case n := <- updates:
					fmt.Printf("status changed to %d\n", n) //---------------------------------------------* What happens when status change
			}


		}
	}(updates)

	/* > Update the statuses every 5 minutes. */
	for {
		updateStatuses(all_machines, updates)
		time.Sleep(5 * time.Minute)
	}

}

// getMachines takes the unmarshalled config.json and construct a slice of
// pointers to Machine structs representing all the machines in all the labs.
func getMachines(labs []interface{}) []*Machine {
	all_machines := make([]*Machine, 0)

	for lab := range labs {
		this_lab := labs[lab].(map[string]interface{})
		prefix := this_lab["prefix"].(string)
		start := int(this_lab["start"].(float64))
		end := int(this_lab["end"].(float64))

		for i := start; i <= end; i++ {
			hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)
			all_machines = append(all_machines, &Machine{hostname, 2})
		}
	}
	return all_machines
}
