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
func GetStatus(hostname string) int {
	if Accessible(hostname, "***REMOVED***") {
		return LINUX
	} else if Accessible(hostname, "***REMOVED***") {
		return WINDOWS
	} else {
		return INACCESSIBLE
	}
}

// accessible takes a hostname and a port number and tries to establish a
// connection using those parameters.
func Accessible(hostn string, port string) bool {
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
func UpdateStatuses(machines []*Machine, updatesChannel chan *Machine) {
	fmt.Println("updating")
	var wg sync.WaitGroup
	for _, machine := range machines {
		wg.Add(1)

		go func(m *Machine) {
			defer wg.Done()
			m.Update(updatesChannel)
		}(machine)
	}
	wg.Wait()
}

// Update takes the updates channel. For the Machine it was called on, it
// checks whether the status has changed, and sends any changes on the updates
// channel, and changes the status.
func (m *Machine) Update(updatesChannel chan *Machine) {
	newStatus := GetStatus(m.hostname)

	if newStatus != m.status {
		m.status = newStatus
		updatesChannel <- m
	}
}
