package main

import(
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json")

// main opens and reads config.json into a MyJson type, and then proceeds to
// find the machine statuses for all the labs outlined in the config file.
func main() {
  labs := getConfig("config.json") //information directly from the config file.
  lab_status := processAllLabs(labs) //return the statuses in format to send.

  serve(lab_status)



}








/* FUNCTIONS */

// processAllLabs obtains the beginning statuses of all the labs.
func processAllLabs(labs []interface{}) (map[string][]map[string]int) {
  all_labs := make(map[string][]map[string]int)

  for lab := range labs {
    prefix,one_lab := processLab(labs[lab].(map[string]interface{}))
    all_labs[prefix] = one_lab
  }
  return all_labs
}

// processLab takes the information for one lab and returns an ordered list
// containing the operating statuses for each machine.
func processLab(lab map[string]interface{}) (string, []map[string]int) {
  prefix := lab["prefix"]
  start := int(lab["start"].(float64))
  end := int(lab["end"].(float64))
  machines_in_lab := make([]map[string]int, 1)

  for i := start; i <= end; i++ {
    hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)

    machine := make(map[string]int)
    machine["machine"], machine["status"] = i, operatingSystem(hostname)
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
