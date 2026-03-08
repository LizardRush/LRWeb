package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {

	path := "./public/pages" + r.URL.Path

	// if root, serve index
	if r.URL.Path == "/" {
		path = "./public/pages/index.html"
	} else {
		path = filepath.Join("./public/pages", r.URL.Path, "index.html")
	}

	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, path)
}

func main() {

	// serve assets
	fs := http.FileServer(http.Dir("./public/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// page routing
	http.HandleFunc("/", pageHandler)

	fmt.Println("Website running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}