package main

import(
  "os"
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json")

// main opens and reads config.json into a MyJson type, and then proceeds to
// find the machine statuses for all the labs outlined in the config file.
func main() {

  /* Start the server set up framework of labs from config.json */
  go http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))
  //check(err)

  /* Get initial status of every machine */
  labs := getConfig("./static/config.json")
  lab_statuses := processAll(labs)
  status, err := json.Marshal(lab_statuses) //right here 'status' is of type []uint8
  check(err)
  _ = status

  /* update the status once */
  update(lab_statuses)
}






/* FUNCTIONS */

// processAll calls process on each lab to obtain the initial status of all the
// machines.
func processAll(labs []interface{}) (map[string][]map[string]int) {
  all_labs := make(map[string][]map[string]int)

  for lab := range labs {
    prefix,one_lab := process(labs[lab].(map[string]interface{}))
    all_labs[prefix] = one_lab
  }
  return all_labs
}

// process gets called on one lab. It calls systemStatus on each machine in the
// lab to obtain the initial statuses of all the machines.
func process(lab map[string]interface{}) (string, []map[string]int) {
  prefix := lab["prefix"]
  start := int(lab["start"].(float64))
  end := int(lab["end"].(float64))
  machines_in_lab := make([]map[string]int, 1)

  for i := start; i <= end; i++ {
    hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)

    machine := make(map[string]int)
    machine["machine"], machine["status"] = i, systemStatus(hostname)
    machines_in_lab = append(machines_in_lab, machine)
  }
  return prefix.(string), machines_in_lab
}

// check takes an error. If the error exists it gets printed, and then the
// program exits. If it doesn't exist, nothing happens.
func check(e error) {
  if e != nil {
    fmt.Printf("Error: %v", e)
    os.Exit(1)
  }
}

// getConfig takes the name of the configuration file (currently "config.json")
// attempts to open/read file then unmarshal it into a list of interfaces.
func getConfig(file_name string) []interface{} {
  config_file, err := ioutil.ReadFile(file_name)
  check(err)

  var labs []interface{}
  err = json.Unmarshal(config_file, &labs)
  check(err)
  return labs
}
