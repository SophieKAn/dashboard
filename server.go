package main

///////////////
// Server.go //
///////////////

import (
	"./static"
	"encoding/json"
	"fmt"
	"log"
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
	allMachines := getMachines(config.MachineRanges, config.Domain)

	/* > Run the Hub */
	hub := newHub()
	go hub.run()

	/* > Define handlers */
	http.HandleFunc("/", serveString(static.Index, "text/html"))
	http.HandleFunc("/css/style.css", serveString(static.Style, "text/css"))
	http.HandleFunc("/js/script.js", serveString(static.Script, "application/javascript"))
	http.HandleFunc("/config.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(config)
		check(err)
		fmt.Fprintf(w, string(data))
	})
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		serveUpdates(hub, allMachines, w, r)
	})

	/* > Start the server */
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
		} else if config.Debug {
			log.Println("no changes")
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

// serveString takes a static file represented as a string and sets the correct
// content-type to be written to the ResponseWriter.
func serveString(s string, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		fmt.Fprintf(w, s)
	}
}

// debugMode checks the config to see if Debug is true, and if so prints
// the current settings.
func debugMode(config *Config) {
	if config.Debug {
		fmt.Printf("interface: %s\n", config.Interface)
		fmt.Printf("port:      %s\n", config.Port)
		fmt.Printf("interval:  %s\n", config.Interval)
		fmt.Printf("domain:    %s\n", config.Domain)
		fmt.Printf("debug:     %t\n", config.Debug)
	}
}
