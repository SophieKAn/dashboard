package main

/////////////
// Main.go //
/////////////

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	args, err := docopt.Parse(usage, nil, true, Version, false)
	if err != nil {
		fmt.Println(err)
	}
  fmt.Println(len(args))

	/* > Get lab configuration */
	labConfig := GetConfig("./static/config.json")
	allMachines := GetMachines(labConfig)

	/* > Run the Hub */
	hub := newHub()
	go hub.run()

	/* > Start the server */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		ServeUpdates(hub, allMachines, w, r)
	})

	go http.ListenAndServe("localhost:8080", nil)

	/* > Update forever */
	var updates []*Machine
	for {
		for machine := range UpdateStatuses(allMachines) {
			updates = append(updates, machine)
		}

		if updates != nil {
			message, err := json.Marshal(updates)
			Check(err)
			hub.broadcast <- message
			updates = nil
		} else {
			fmt.Println("no changes")
		}

		time.Sleep(5 * time.Second)
	}
}

func ServeUpdates(hub *Hub, allMachines []*Machine, w http.ResponseWriter, r *http.Request) {
	/* > Open the websocket connection. */
	ws, err := upgrader.Upgrade(w, r, nil)
	Check(err)
	defer ws.Close()

	client := &Client{hub, ws, make(chan []byte)}
	hub.register <- client
	go func() {
		data, _ := json.Marshal(allMachines)
		client.send <- data
	}()
	client.writePump()
}
