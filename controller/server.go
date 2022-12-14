package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const route string = "/log"
const altRoute string = ""

type Data struct {
	NodeId    string    `json:"nodeid"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func ParseLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var d Data

		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Do something with the Person struct...
		fmt.Println(d)

	}

}

func logBackup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BACKUP REQUEST")
}

func initRoutes() {
	http.HandleFunc(route, ParseLogs)
	http.HandleFunc(altRoute, logBackup)
}

func StartServer() {
	initRoutes()
	http.ListenAndServe(":3001", nil)
}
