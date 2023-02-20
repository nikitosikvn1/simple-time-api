package main

import (
	"encoding/json"
	"net/http"
	"time"
	"log"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	timeResponse := TimeResponse{Time: currentTime}

	timeResponseJSON, err := json.Marshal(timeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(timeResponseJSON)
}

func main() {
    http.HandleFunc("/time", currentTimeHandler)

    err := http.ListenAndServe(":8795", nil)
    if err != nil {
        log.Fatal(err)
    }
}
