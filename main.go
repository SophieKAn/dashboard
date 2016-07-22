package main

/////////////
// Main.go //
/////////////

import (
	"github.com/docopt/docopt-go"
)

const Version = "1.0.0"

var usage = `Start a server to display the current usage of all the labs.

Usage:
  dashboard -v | --version
  dashboard -h | --help

Options:
  -v, --version  Show version
  -h, --help     Show this message`

type Machine struct {
	Hostname string `json:"hostname"`
	Status   int    `json:"status"`
}

func main() {
	_, _ = docopt.Parse(usage, nil, true, Version, false)

	Server()
}
