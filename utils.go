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

// Check takes an error, and prints the error if it isn't nil.
func Check(e error) {
	if e != nil {
		log.Printf("Error: %v\n", e)
		os.Exit(1)
	}
}

// GetConfig takes the name of the configuration file (currently "config.json")
// and attempts to open/read file then unmarshal it into a list of interfaces.
func GetConfig(fileName string) map[string]interface{} {

	configFile, err := ioutil.ReadFile(fileName)
	Check(err)

	var configuration map[string]interface{}
	err = json.Unmarshal(configFile, &configuration)
	Check(err)

	return configuration
}

// GetMachines takes the unmarshalled config.json and constructs a slice of
// pointers to Machine structs representing all the machines indicated in the
// config.
func GetMachines(labs []map[string]interface{}) []*Machine {
	var allMachines []*Machine
	for _, lab := range labs {
		prefix := lab["prefix"].(string)
		start, end := int(lab["start"].(float64)), int(lab["end"].(float64))

		for i := start; i <= end; i++ {
			hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)
			allMachines = append(allMachines, &Machine{hostname, 2})
		}
	}
	return allMachines
}
