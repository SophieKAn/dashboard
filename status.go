package main

///////////////
// Status.go //
///////////////

import (
	"net"
	"sync"
	"time"
	"fmt"
)

const (
	LINUX        = 0
	WINDOWS      = 1
	INACCESSIBLE = 2
)
//
//
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
	fmt.Println(hostn)
	conn, err := net.DialTimeout("tcp", hostn+":"+port, time.Millisecond*50)

	if err == nil {
		conn.Close()
		return true
	} else {
		return false
	}
}

//
//
func updateStatuses(machines []*Machine) {
	var wg sync.WaitGroup
	for _, machine := range machines {
		wg.Add(1)

		go func(m *Machine) {
			defer wg.Done()
			m.UpdateStatus()
		}(machine)
	}
	wg.Wait()
}

//
//
func (m *Machine) UpdateStatus() {
	old_status := m.status
	new_status := getStatus(m.hostname)

	if new_status != old_status {
		// Send out changes
		m.status = new_status
	}
}
