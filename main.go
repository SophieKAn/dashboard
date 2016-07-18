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

type Machine struct {
	hostname string
	status   int
}


func main() {

	labConfig := GetConfig("./static/config.json")
	allMachines := GetMachines(labConfig)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", ServeUpdates)

	go hub.run()

	go http.ListenAndServe("localhost:8080", nil)

	for {
		c := UpdateStatuses(allMachines)
		time.Sleep(5 * time.Second)
	}
}


func ServeUpdates(w http.ResponseWriter, r *http.Request) {
	/* > Open the websocket connection. */
	ws, err := upgrader.Upgrade(w, r, nil); Check(err); defer ws.Close()

	alist := [2]string{"P","T"}
	jsn, err := json.Marshal(alist); Check(err)

	/* > Send message to client */
	err = ws.WriteMessage(websocket.TextMessage, jsn); Check(err)
}
