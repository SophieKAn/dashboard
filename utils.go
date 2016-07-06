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

// check takes an error. If the error exists it gets printed, and then the
// program exits. If it is nil, nothing happens.
func Check(e error) bool {
	if e != nil {
		fmt.Printf("Error: %v", e)
		os.Exit(1)
	}
	return true
}

// getConfig takes the name of the configuration file (currently "config.json")
// and attempts to open/read file then unmarshal it into a list of interfaces.
func GetConfig(fileName string) []interface{} {
	configFile, err := ioutil.ReadFile(fileName)
	Check(err)

	var labs []interface{}
	err = json.Unmarshal(configFile, &labs)
	Check(err)
	return labs
}
