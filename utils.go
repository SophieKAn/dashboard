package main

import (
	"fmt"
	"os"
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
