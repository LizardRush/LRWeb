package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

type api struct{
	websiteModule
}

func main() {
	http.HandleFunc("/", ping)

	// website files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":5000", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
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