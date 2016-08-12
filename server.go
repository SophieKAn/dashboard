package main

///////////////
// Server.go //
///////////////

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Machine struct {
	Hostname string `json:"hostname"`
	Status   string `json:"status"`
}

// runServer takes the config struct. It runs a hub, starts the server, and
// continually updates the status of all the machines, then broadcasting those
// changes to all connected clients in the hub.
func runServer(config *Config) {
	/* > Check for debug mode */
	debugMode(config)

	/* > Get lab configuration */
	allMachines := getMachines(config.MachineRanges)

	/* > Run the Hub */
	hub := newHub()
	go hub.run()

	/* > Start the server */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/config.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(config)
		check(err)
		fmt.Fprintf(w, string(data))
	})
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		serveUpdates(hub, allMachines, w, r)
	})

	go func() {
		err := http.ListenAndServe(config.Interface+":"+config.Port, nil)
		check(err)
	}()

	/* > Update forever */
	var updates []*Machine
	for {
		for machine := range updateStatuses(allMachines, config) {
			updates = append(updates, machine)
		}

		if updates != nil {
			message, err := json.Marshal(updates)
			check(err)
			hub.broadcast <- message
			updates = nil
		} else {
			fmt.Println("no changes")
		}

		time.Sleep(config.Interval)
	}
}

// serveUpdates responds to a websocket connection by creating a 'client',
// sending said client to the hub, sending it the set of all machines, and
// finally calling writePump().
func serveUpdates(hub *Hub, allMachines []*Machine, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	check(err)
	defer ws.Close()

	client := &Client{hub, ws, make(chan []byte)}
	hub.register <- client
	go func() {
		data, _ := json.Marshal(allMachines)
		client.send <- data
	}()
	client.writePump()
}

// debugMode checks the config to see if Debug is true, and if so prints
// the current settings.
func debugMode(config *Config) {
	if config.Debug {
		fmt.Printf("interface: %s\n", config.Interface)
		fmt.Printf("port:      %s\n", config.Port)
		fmt.Printf("interval:  %s\n", config.Interval)
		fmt.Printf("debug:     %t\n", config.Debug)
	}
}
