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

func RunServer(configs Config) {

	/* > Get lab configuration */
	settings := GetConfig(configs.Configfile)
	allMachines := GetMachines(settings["labsetup"].([]interface{}))


	/* > Run the Hub */
	hub := newHub()
	go hub.run()

	/* > Start the server */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		ServeUpdates(hub, allMachines, w, r)
	})

	go func() {
		err := http.ListenAndServe(configs.Interface+":"+configs.Port, nil)
		Check(err)
	}()

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

		time.Sleep(configs.Interval)
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
