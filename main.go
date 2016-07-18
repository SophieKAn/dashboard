package main

/////////////
// Main.go //
/////////////

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader websocket.Upgrader

type Machine struct {
	hostname string
	status   int
}

//
//
func main() {
	/* > Get lab configuration from config file. */
	labConfig := GetConfig("./static/config.json")

	/* > Create a struct for each machine. */
	var allMachines []*Machine
	allMachines = GetMachines(labConfig, allMachines)

	/* > Establish handlers. */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", func(w http.ResponseWriter, r *http.Request) {
		ServeUpdates(w, r, allMachines)
	})

	/* Start the server. */
	go http.ListenAndServe("localhost:8080", nil)

	/* > Update statuses forever at the given pace. */
	go func() {
		for {
			UpdateStatuses(allMachines)
			time.Sleep(5 * time.Second) }
	}()
}

//
//
func ServeUpdates(w http.ResponseWriter, r *http.Request, allMachines []*Machine) {
	/* > Open the websocket connection. */
	ws, err := upgrader.Upgrade(w, r, nil); Check(err); defer ws.Close()

	/* > Marshal allMachines into a JSON. */
	jsn, err := json.Marshal(allMachines); Check(err)

	/* > Send message to client */
	err = ws.WriteMessage(websocket.TextMessage, jsn); Check(err)
}
