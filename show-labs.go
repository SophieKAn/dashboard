/* Sophia Anderson | ander569@wwu.edu */
// The goal of this program is to access all the computers in the labs and
// indicate whether each one is running Windows or Linux, or is inaccessible.

package main

import(
  "os"
  "net"
  "fmt"
  "time"
  "color"
  "io/ioutil"
  "encoding/json"
  //"reflect"
)

type MyJson []map[string]string

func main() {
  var data MyJson

  config_file, err := ioutil.ReadFile("config.json")
  if err != nil {
    fmt.Printf("File error: %v", err)
    os.Exit(1)
  }

  err = json.Unmarshal(config_file, &data)
  if err != nil {
    fmt.Printf("File error: %v", err)
    os.Exit(1)
  }


  /*for index,_ := range data {
    processLab(data[index])
  }*/








    fmt.Println(operatingSystem("linux-10.***REMOVED***"))

}



/* OPERATING_SYSTEM */
// Inputs:An IP address in String form
// Outputs:A string indicating the status of the machine
// Function:Given the machine's IP address, this function will
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




func processLab(lab []map[string]string) (string) {

  //fmt.Println(lab["title"])
  return "yes"
}
