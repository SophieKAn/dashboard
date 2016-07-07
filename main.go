package main

/////////////
// Main.go //
/////////////

import (
	//"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{}

type Machine struct {
	Hostname string
	Status   int
}

// main starts the server, gets all the lab info from config.json, sets up a
// channel for receiving updates, and then updates system statuses every 5 min.
func main() {

	/* > Define handlers. */
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", ServeUpdates)

	/* > Start the server. */
	http.ListenAndServe("localhost:8080", nil)
}




func ServeUpdates(w http.ResponseWriter, r *http.Request) {

	/* > Open the websocket connection. */
	ws, err := upgrader.Upgrade(w, r, nil)
	Check(err); defer ws.Close()


	/* > Get lab configuration from config file. */
	labConfig := GetConfig("./static/config.json")

	/* > Create a struct for each machine. */
	allMachines := GetMachines(labConfig)


	/* > Create channel to receive status updates. */
	updatesChannel := make(chan *Machine)

	go func(updatesChannel chan *Machine ) {
		for {
			sendUpdate(<-updatesChannel, ws)
		}
	}(updatesChannel)

	for {
		UpdateStatuses(allMachines, updatesChannel)
		time.Sleep(10 * time.Second)
	}
}





//
//
func sendUpdate(machine *Machine, ws *websocket.Conn) {
	machine.Hostname = strings.TrimSuffix(machine.Hostname, ".***REMOVED***")
	jmsg, err := json.Marshal(machine)
	Check(err)

	err = ws.WriteMessage(websocket.TextMessage, jmsg)
	Check(err)
}
