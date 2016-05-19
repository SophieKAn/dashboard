/* Sophia Anderson | ander569@wwu.edu */
// The goal of this program is to access all the computers in the labs and
// indicate whether each one is running Windows or Linux, or is inaccessible.

package main

import(
  "net"
  "fmt"
  "time"
  "color"
  //"io/ioutil"
  //"encoding/json"
)

func main() {
  for i := 98; i < 110; i++ {
    ip := fmt.Sprintf("***REMOVED***.30.%v", i)
    fmt.Println(operatingSystem(ip))
  }
}

/* OPERATING_SYSTEM */
// Inputs:An IP address in String form
// Outputs:A string indicating the status of the machine
// Function:Given the machine's IP address, this function will
//          return a string containing the current OS of that particular
//          machine, or if it is inaccessible at the time.
func operatingSystem(IP string) (string) {
  //look up address
  addr, err := net.LookupAddr(IP)
  if err != nil {
    return color.YellowString(" issue with IP address")
  }
  //print address
  var adstring string = addr[0]
  fmt.Print(adstring[0:len(adstring)-1])
  //try to connect on various ports
  if tryToConnect(IP, "***REMOVED***") == nil {
    return color.GreenString(" linux")
  } else if tryToConnect(IP, "***REMOVED***") == nil {
    return color.BlueString(" windows")
  } else {
    return color.RedString(" inaccessible")
  }
}

/* TRY_TO_CONNECT */
// Inputs:Two strings: an IP address and a port number
// Outputs:An error, either nil or not
// Function:Tries to connect to the given machine using the given port.
//          returns a new error upon failure, and a nil error upon success.
func tryToConnect(IP string, port string) (error) {
  conn, err := net.DialTimeout("tcp", IP + ":" + port, time.Millisecond*50)
  if err == nil {
    conn.Close()
  }
  return err
}
