// Package main is responsible for obtaining the statuses of the machines/labs
// indicated in config.json. It periodically checks for changes and updates the
// client with new information.
package main

import(
  "os"
  "net"
  "fmt"
  "time"
  "color"
  "io/ioutil"
  "encoding/json"
)

// main opens and reads config.json into a MyJson type, and then proceeds to
// find the machine statuses for all the labs outlined in the config file.
func main() {
  config_file, err := ioutil.ReadFile("config.json")
  check(err)

  var labs []interface{}
  err = json.Unmarshal(config_file, &labs)
  check(err)

  for lab,_ := range labs {
    status_list  := processLab(labs[lab].(map[string]interface{}))
    fmt.Println(status_list[0])
  }
}

// operatingSystem takes a hostname(str) and returns what operating system that
// machine is running based on which port is successfully used to connect. 
// It returns 'linux', 'windows', or 'inaccessible'.
func operatingSystem(hostname string) (string) {
  //try to connect on various ports
  if tryToConnect(hostname, "***REMOVED***") == nil {
    fmt.Println("linux")
    return color.GreenString("linux")
  } else if tryToConnect(hostname, "***REMOVED***") == nil {
    fmt.Println("windows")
    return color.BlueString("windows")
  } else {
    fmt.Println("inaccessible")
    return color.RedString("inaccessible")
  }
}

// tryToConnect takes a hostname and a port number and tries to establish a
// connection using those parameters.
// It returns an error, if one occurrs.
func tryToConnect(hostname string, port string) (error) {
  fmt.Println("Trying to connect to " + hostname + ":" + port)
  conn, err := net.DialTimeout("tcp", hostname + ":" + port, time.Millisecond*50)
  if err == nil {
    conn.Close()
  }
  return err
}

// processLab takes the information for one lab and returns an ordered list
// containing the operating statuses for each machine.
func processLab(lab map[string]interface{}) ([]string) {
  prefix := lab["prefix"]
  start := int(lab["start"].(float64))
  end := int(lab["end"].(float64))
  var status_list []string

  for i := start; i <= end; i++ {
    hostname := fmt.Sprintf("%s-%02d.***REMOVED***\n", prefix, i)
    status_list = append(status_list, operatingSystem(hostname))
  }
  return status_list
}

// check
func check(e error) {
  if e != nil {
    fmt.Printf("Error: %v", e)
    os.Exit(1)
  }
}
