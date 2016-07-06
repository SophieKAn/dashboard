package main

///////////////
// Status.go //
///////////////

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	LINUX        = 0
	WINDOWS      = 1
	INACCESSIBLE = 2
)

// getStatus takes a hostname and checks whether it is available on port ***REMOVED***
// or ***REMOVED*** (linux and windows respectively). Otherwise it is labeled
// inaccessible.
func getStatus(hostname string) int {
	if accessible(hostname, "***REMOVED***") {
		return LINUX
	} else if accessible(hostname, "***REMOVED***") {
		return WINDOWS
	} else {
		return INACCESSIBLE
	}
}

// accessible takes a hostname and a port number and tries to establish a
// connection using those parameters.
func accessible(hostn string, port string) bool {
	conn, err := net.DialTimeout("tcp", hostn+":"+port, 50 * time.Millisecond)

	if err == nil {
		conn.Close()
		return true
	} else {
		return false
	}
}

// updateStatuses takes the list of Machine pointers and iterates through them
// using goroutines to call Update for each one. It waits until all goroutines
// are finished before returning.
func updateStatuses(machines []*Machine, updates chan *Machine) {
	fmt.Println("updating")
	var wg sync.WaitGroup
	for _, machine := range machines {
		wg.Add(1)

		go func(m *Machine) {
			defer wg.Done()
			m.Update(updates)
		}(machine)
	}
	wg.Wait()
}

// Update takes the updates channel. For the Machine it was called on, it
// checks whether the status has changed, and sends any changes on the updates
// channel, and changes the status.
func (m *Machine) Update(updates chan *Machine) {
	new_status := getStatus(m.hostname)

	if new_status != m.status {
		m.status = new_status
		updates <- m
	}
}
