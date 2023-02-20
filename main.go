package main

import "net/http"

type TimeResponse struct {
	Time string `json:"time"`
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {}

func main() {}
