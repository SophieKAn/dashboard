package main

import (
	"net"
	"time")

// operatingSystem takes a hostname(str) and returns what operating system that
// machine is running based on which port is successfully used to connect. 
// It returns 'linux', 'windows', or 'inaccessible'.
func systemStatus(hostname string) (int) {
  //try to connect on various ports
  if accessible(hostname, "***REMOVED***") {
    return 0 
  } else if accessible(hostname, "***REMOVED***") {
    return 1 
  } else {
    return 2
  }
}


// accessible takes a hostname and a port number and tries to establish a
// connection using those parameters.
// It returns an error, if one occurrs.
func accessible(hostn string, port string) (bool) {
  var ok bool
  conn, err := net.DialTimeout("tcp", hostn + ":" + port, time.Millisecond*50)
  if err == nil {
    conn.Close()
    ok = true
  } else {
    ok = false
  }
  return ok
}

// update will be responsible for updating the status of the lab machines.
func update([]byte) ([]byte) {
}
