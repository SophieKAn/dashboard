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


/* Main */

func main() {
  //Read configuration file
  config_file, err := ioutil.ReadFile("config.json")
  if err != nil {
    fmt.Printf("File error: %v", err)
    os.Exit(1)
  }

  //Unmarshal config.json
  var data MyJson
  err = json.Unmarshal(config_file, &data)
  if err != nil {
    fmt.Printf("File error: %v", err)
    os.Exit(1)
  }

  for index,_ := range data {
    status_list  := processLab(data[index])
  }


}






/* Functions */

/* OPERATING_SYSTEM */
// Inputs: The hostname in String form.
// Outputs: A string indicating the status of the hostname.
// Function: Given a machine's hostname, this function will try to connect to
//           the host and return 'linux,' 'windows,' or 'inaccessible.'
func operatingSystem(hostname string) (string) {
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
// Inputs: A hostname and a port number as strings.
// Outputs: An error or a nil error.
// Function: Tries to connect to the given machine using the given port.
//           Returns an error upon failure, and nil upon success.
func tryToConnect(hostname string, port string) (error) {
  conn, err := net.DialTimeout("tcp", hostname + ":" + port, time.Millisecond*50)
  if err == nil {
    conn.Close()
  }
  return err
}

/* PROCESS_LAB */
// Inputs: A map from the json that contains info about the lab.
// Outputs: A list of strings with the statuses of all the machines in the lab.
// Function: Finds the status of all the machines in the given lab.
func processLab(lab map[string]string) ([]string){

  prefix := lab["prefix"]
  start, err1 := strconv.Atoi(lab["start"])
  end, err2 := strconv.Atoi(lab["end"])
  if (err1 != nil || err2 != nil) {
    fmt.Printf("Error with JSON file")
    os.Exit(1)
  }
  //Get status and put in list
  var status_list []string //create a list for the lab

  for i := start; i<=end; i++ { //for each machine
    if i < 10 {
      status_list = append(status_list,
                           operatingSystem(prefix + "-0" + strconv.Itoa(i) + ".***REMOVED***"))
    } else {
      status_list = append(status_list,
                           operatingSystem(prefix + "-" + strconv.Itoa(i) + ".***REMOVED***"))
    }
  }
  return status_list
}
