package main

///////////////
// Status.go //
///////////////

import (
	"log"
	"net"
	"sync"
	"time"
)

// updateStatuses takes the list of Machine pointers and iterates through them
// using nested goroutines to call Update for each one. It waits until all
// goroutines are finished before returning.
func updateStatuses(machines []*Machine, config *Config) chan *Machine {
	if config.Debug {
		log.Println("updating")
	}
	out := make(chan *Machine)
	go func(chan *Machine) {
		var wg sync.WaitGroup
		for _, machine := range machines {
			wg.Add(1)

			go func(m *Machine) {
				defer wg.Done()
				m.Update(out, config)
			}(machine)
		}
		wg.Wait()
		close(out)
	}(out)
	return out
}

// Update takes an output channel and the config settings. For machine m, it
// checks for a change in status, and if it has changed sends itself along the
// 'out' channel.
func (m *Machine) Update(out chan *Machine, config *Config) {
	newStatus := getStatus(m.Hostname, config)

	if newStatus != m.Status {
		m.Status = newStatus
		out <- m
	}
}

// getStatus takes a hostname and the config, and checks that hostname
// according to all the machine identifiers in the config.
func getStatus(hostname string, config *Config) string {

	for _, identifier := range config.MachineIdentifiers {
		if accessible(hostname, identifier["port"].(string)) {
			return identifier["name"].(string)
		}
	}

	return "inaccessible"
}

// accessible takes a hostname and a port number and tries to establish a
// connection using those parameters.
func accessible(hostn string, port string) bool {
	conn, err := net.DialTimeout("tcp", hostn+":"+port, 1*time.Second)

	if err == nil {
		conn.Close()
		return true
	} else {
		return false
	}
}
