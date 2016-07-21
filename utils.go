package main

//////////////
// Utils.go //
//////////////

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Check takes an error, and prints the error if it isn't nil.
func Check(e error) {
	if e != nil {
		fmt.Printf("Error: %v", e)
		os.Exit(1)
	}
}

// GetConfig takes the name of the configuration file (currently "config.json")
// and attempts to open/read file then unmarshal it into a list of interfaces.
func GetConfig(fileName string) []interface{} {

	configFile, err := ioutil.ReadFile(fileName)
	Check(err)

	var labs []interface{}
	err = json.Unmarshal(configFile, &labs)
	Check(err)

	return labs
}

// GetMachines takes the unmarshalled config.json and constructs a slice of
// pointers to Machine structs representing all the machines indicated in the
// config.
func GetMachines(labs []interface{}) []*Machine {
	var allMachines []*Machine
	for lab := range labs {
		aLab := labs[lab].(map[string]interface{})
		prefix := aLab["prefix"].(string)
		start, end := int(aLab["start"].(float64)), int(aLab["end"].(float64))

		for i := start; i <= end; i++ {
			hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)
			allMachines = append(allMachines, &Machine{hostname, 2})
		}
	}
	return allMachines
}
