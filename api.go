package main

import (
	"encoding/json"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "running",
		"name":   "LRWeb API",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/status", status)
	http.ListenAndServe(":5000", nil)
}
