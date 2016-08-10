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
	Status   string    `json:"status"`
}

func RunServer(config Config) {

	if config.Debug {
		fmt.Printf("interface: %s\n", config.Interface)
		fmt.Printf("port:      %s\n", config.Port)
		fmt.Printf("interval:  %s\n", config.Interval)
		fmt.Printf("debug:     %t\n", config.Debug)
	}

	/* > Get lab configuration */
	allMachines := GetMachines(config.MachineRanges)

	/* > Run the Hub */
	hub := newHub()
	go hub.run()

	/* > Start the server */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/config.json", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(config)
		fmt.Fprintf(w, string(data))
	})
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		ServeUpdates(hub, allMachines, w, r)
	})

	go func() {
		err := http.ListenAndServe(config.Interface+":"+config.Port, nil)
		Check(err)
	}()

	/* > Update forever */
	var updates []*Machine
	for {
		for machine := range UpdateStatuses(allMachines, config) {
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

		time.Sleep(config.Interval)
	}
}

func ServeUpdates(hub *Hub, allMachines []*Machine, w http.ResponseWriter, r *http.Request) {
	/* > Open the websocket connections. */
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
