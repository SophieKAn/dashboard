package main

//////////////
// Utils.go //
//////////////

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// check takes an error, and prints the error if it isn't nil, then exits.
func check(e error) {
	if e != nil {
		log.Printf("Error: %v\n", e)
		os.Exit(1)
	}
}

// getConfig takes the name of the configuration file, attempts to open the
// file and unmarshal it into a map of strings to interfaces.
func getConfig(fileName string) map[string]interface{} {

	configFile, err := ioutil.ReadFile(fileName)
	check(err)

	var configuration map[string]interface{}
	err = json.Unmarshal(configFile, &configuration)
	check(err)

	return configuration
}

// getMachines takes the unmarshalled "machineRanges" list and constructs a
// slice of pointers to Machine structs representing all the machines from the
// config.
func getMachines(labs []map[string]interface{}) []*Machine {
	var allMachines []*Machine
	for _, lab := range labs {
		prefix := lab["prefix"].(string)
		start, end := int(lab["start"].(float64)), int(lab["end"].(float64))

		for i := start; i <= end; i++ {
			hostname := fmt.Sprintf("%s-%02d.generic-domain", prefix, i)
			allMachines = append(allMachines, &Machine{hostname, "inaccessible"})
		}
	}
	return allMachines
}
