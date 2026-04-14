package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var  page_path string = "./websites/main/"

var port_webpage string = "8080"

func pageHandler(w http.ResponseWriter, r *http.Request) {

	path := page_path + "pages" + r.URL.Path

	// if root, serve index
	if r.URL.Path == "/" {
		path = page_path + "pages/index.html"
	} else {
		path = filepath.Join("./websites/main/pages", r.URL.Path, "index.html")
	}

	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if r.URL.Path == "/" {
			path = r.URL.Path
		} else {
			path = filepath.Join("./websites/main/assets", r.URL.Path)
		}
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, path)
}

func main() {

	// page routing
	http.HandleFunc("/", pageHandler)

	fmt.Println("Website running on http://localhost:" + port_webpage)

	http.ListenAndServe(":" + port_webpage, nil)
}