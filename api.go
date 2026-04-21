package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var port_api string = "5553"

func main() {
	http.HandleFunc("/ping", ping)

	fmt.Println("Website running on http://localhost:" + port_api)
	http.ListenAndServe(":" + port_api, nil)
}

func ping(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "GET only", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Server")

	response := map[string]string{
		"status": "active", 
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}