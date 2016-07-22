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
	Status   int    `json:"status"`
}

func Server() {

	/* > Get lab configuration */
	labConfig := GetConfig(defaultConfig)
	allMachines := GetMachines(labConfig)

	/* > Run the Hub */
	hub := newHub()
	hub.run()

	/* > Start the server */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		ServeUpdates(hub, allMachines, w, r)
	})

	go http.ListenAndServe(defaultInterface+":"+defaultPort, nil)

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
