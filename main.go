package main

/////////////
// Main.go //
/////////////

import (
	"encoding/json"
	//"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Machine struct {
	hostname string
	status   int
}


func main() {

	var updates []*Machine

	labConfig := GetConfig("./static/config.json")
	allMachines := GetMachines(labConfig)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", ServeUpdates)

	go hub.run()

	go http.ListenAndServe("localhost:8080", nil)

	for {
		c := UpdateStatuses(allMachines)
		for machine := range c {
			updates = append(updates, machine)
		}

		message, err := json.Marshal(updates); Check(err)
		hub.broadcast <- message
		updates = nil
		time.Sleep(5 * time.Second)
	}
}


func ServeUpdates(w http.ResponseWriter, r *http.Request) {
	/* > Open the websocket connection. */
	ws, err := upgrader.Upgrade(w, r, nil); Check(err); defer ws.Close()

	hub.register <- &Conn{ws, make(chan []byte)}
}
