/* Sophia Anderson | ander569@wwu.edu */

package main

import(
  "os"
  "net"
  "fmt"
  "time"
  "color"
  "strconv"
  "io/ioutil"
  "encoding/json"
)

type MyJson []map[string]string

func main() {

  config_file, err := ioutil.ReadFile("config.json")
  if err != nil {
    fmt.Printf("File error: %v", err)
    os.Exit(1)
  }

  var data MyJson
  err = json.Unmarshal(config_file, &data)
  if err != nil {
    fmt.Printf("File error: %v", err)
    os.Exit(1)
  }

  type mainMessage struct {
    Title string
    Access map[string]string
  }
  for index,_ := range data {
    processLab(data[index])
  }
}

/* OPERATING_SYSTEM */
// Inputs:The hostname in String form
// Outputs:A string indicating the status of the machine
// Function:Given the machine's hostname, this function will
//          return a string containing the current OS of that particular
//          machine, or if it is inaccessible at the time.
func operatingSystem(hostname string) (string) {
  fmt.Print(hostname)
  //try to connect on various ports
  if tryToConnect(hostname, "***REMOVED***") == nil {
    return color.GreenString(" linux")
  } else if tryToConnect(hostname, "***REMOVED***") == nil {
    return color.BlueString(" windows")
  } else {
    return color.RedString(" inaccessible")
  }
}

/* TRY_TO_CONNECT */
// Inputs:Two strings: an hostname and a port number
// Outputs:An error, either nil or not
// Function:Tries to connect to the given machine using the given port.
//          returns a new error upon failure, and a nil error upon success.
func tryToConnect(hostname string, port string) (error) {
  conn, err := net.DialTimeout("tcp", hostname + ":" + port, time.Millisecond*50)
  if err == nil {
    conn.Close()
  }
  return err
}

/* PROCESS_LAB */
// Inputs: A list of maps of strings that represents the config.json file.
// Outputs: Nothing yet
// Function: Prints the accessibility of the lab machines present in
//           config.json
func processLab(lab map[string]string) {

  title, prefix := lab["title"], lab["prefix"]
  start, err1 := strconv.Atoi(lab["start"])
  end, err2 := strconv.Atoi(lab["end"])
  if (err1 != nil || err2 != nil) {
    fmt.Printf("Error with JSON file")
    os.Exit(1)
  }
  /* Print Everything */
  fmt.Println("::::::::::::::::::::::: " + title + " :::::::::::::::::::::::")
  for i := start; i<=end; i++ {
    if i < 10 {
      fmt.Println(operatingSystem(prefix + "-0" +
                                  strconv.Itoa(i) + ".***REMOVED***"))
    } else {
      fmt.Println(operatingSystem(prefix + "-" +
                                  strconv.Itoa(i) + ".***REMOVED***"))
    }
  }
  fmt.Println();
}
