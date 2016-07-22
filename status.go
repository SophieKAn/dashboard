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

// UpdateStatuses takes the list of Machine pointers and iterates through them
// using nested goroutines to call Update for each one. It waits until all
// goroutines are finished before returning.
func UpdateStatuses(machines []*Machine) chan *Machine {
	fmt.Println("updating")
	out := make(chan *Machine)
	go func(chan *Machine) {
		var wg sync.WaitGroup
		for _, machine := range machines {
			wg.Add(1)

			go func(m *Machine) {
				defer wg.Done()
				m.Update(out)
			}(machine)
		}
		wg.Wait()
		close(out)
	}(out)
	return out
}

// Update takes an output channel. For machine m, it checks for a change in
// status, and if it has changed sends itself along the 'out' channel.
func (m *Machine) Update(out chan *Machine) {
	newStatus := GetStatus(m.Hostname)

	if newStatus != m.Status {
		m.Status = newStatus
		out <- m
	}
}

// GetStatus takes a Hostname and checks whether it is available on port ***REMOVED***
// or ***REMOVED*** (linux and windows respectively). If not accessible on either port,
// that host is deemed inaccesssible.
func GetStatus(hostname string) int {
	if Accessible(hostname, "***REMOVED***") {
		return LINUX
	} else if Accessible(hostname, "***REMOVED***") {
		return WINDOWS
	} else {
		return INACCESSIBLE
	}
}

// Accessible takes a hostname and a port number and tries to establish a
// connection using those parameters.
func Accessible(hostn string, port string) bool {
	conn, err := net.DialTimeout("tcp", hostn+":"+port, 1*time.Second)

	if err == nil {
		conn.Close()
		return true
	} else {
		return false
	}
}

