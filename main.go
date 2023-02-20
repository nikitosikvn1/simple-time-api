package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	timeResponse := TimeResponse{Time: currentTime}
	timeResponseJSON, err := json.Marshal(timeResponse)
}

func main() {}
