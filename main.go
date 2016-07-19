package main

/////////////
// Main.go //
/////////////

import (
	"encoding/json"
	"net/http"
	"time"
)

type Machine struct {
	hostname string
	status   int
}


func main() {
	/* > Get lab configuration */
	labConfig := GetConfig("./static/config.json")
	allMachines := GetMachines(labConfig)

	/* > Run the Hub */
	hub := newHub()
	go hub.run()

	/* > Start the server */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		ServeUpdates(hub, w, r)
	})

	go http.ListenAndServe("localhost:8080", nil)

	/* > Update forever */
	var updates []*Machine
	for {
		for machine := range UpdateStatuses(allMachines) {
			updates = append(updates, machine)
		}

		message, err := json.Marshal(updates); Check(err)
		hub.broadcast <- message
		updates = nil

		time.Sleep(5 * time.Second)
	}
}


func ServeUpdates(hub *Hub, w http.ResponseWriter, r *http.Request) {
	/* > Open the websocket connection. */
	ws, err := upgrader.Upgrade(w, r, nil); Check(err); defer ws.Close()

	hub.register <- &Client{hub, ws, make(chan []byte)}
}
