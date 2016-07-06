package main

/////////////
// Main.go //
/////////////

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const updateWait = 1 * time.Second

var upgrader = websocket.Upgrader{}

// A Machine represents one system.
type Machine struct {
	hostname string
	status   int
}

// main starts the server, gets all the lab info from config.json, sets up a
// channel for receiving updates, and then updates system statuses every 5 min.
func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/upd", ServeUpdates)

	/* > Start the server in its own Goroutine. */
	go http.ListenAndServe("localhost:8080", nil)

	time.Sleep(5 * time.Minute)
	/* > Get lab configuration from config file. */
	labConfig := GetConfig("./static/config.json")

	/* > Create a struct for each machine. */
	allMachines := GetMachines(labConfig)

	/* > Create channel to receive status updates. */
	updatesChannel := make(chan *Machine)

	go func(updatesChannel chan *Machine) {
		for {
			fmt.Println(<-updatesChannel)
		}
	}(updatesChannel)

	/* > Update the statuses every second. */
	for {
		UpdateStatuses(allMachines, updatesChannel)
		time.Sleep(updateWait)
	}
}

func ServeUpdates(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	Check(err)
	defer ws.Close()

	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, world."))
	Check(err)

}

// getMachines takes the unmarshalled config.json and construct a slice of
// pointers to Machine structs representing all the machines in all the labs.
func GetMachines(labs []interface{}) []*Machine {
	allMachines := make([]*Machine, 0)

	for lab := range labs {
		aLab := labs[lab].(map[string]interface{})
		prefix := aLab["prefix"].(string)
		start, end := int(aLab["start"].(float64)), int(aLab["end"].(float64))

		for i := start; i <= end; i++ {
			hostname := fmt.Sprintf("%s-%02d.***REMOVED***", prefix, i)
			allMachines = append(allMachines, &Machine{hostname, 2})
		}
	}
	return allMachines
}
