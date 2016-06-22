package main

import (
        "fmt"
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
func update(labs map[string][]map[string]int) (map[string][]map[string]int) {
  all_labs :=  make(map[string][]map[string]int)

  for lab_name, machine_list := range labs {
    machines_in_lab := make([]map[string]int, 1)
    for _, machine := range machine_list {
      hostname := fmt.Sprintf("%s-%02d.***REMOVED***", lab_name, machine["machine"])
      old_status := machine["status"]
      new_status := systemStatus(hostname)
      if new_status != old_status {
        new_machine := make(map[string]int)
        new_machine["machine"], new_machine["status"] = machine["machine"], new_status
        machines_in_lab = append(machines_in_lab, new_machine)
      }
    }
  all_labs[lab_name] = machines_in_lab
  } 
  return all_labs
}
