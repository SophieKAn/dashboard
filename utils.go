package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

// check takes an error. If the error exists it gets printed, and then the
// program exits. If it is nil, nothing happens.
func check(e error) bool {
	if e != nil {
		fmt.Printf("Error: %v", e)
		os.Exit(1)
	}
	return true
}

// getConfig takes the name of the configuration file (currently "config.json")
// and attempts to open/read file then unmarshal it into a list of interfaces.
func getConfig(file_name string) []interface{} {
	config_file, err := ioutil.ReadFile(file_name)
	check(err)

	var labs []interface{}
	err = json.Unmarshal(config_file, &labs)
	check (err)
	return labs
}
