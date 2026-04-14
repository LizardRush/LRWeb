package main

import (
	"flag"
	"log"
	"net/http"
)

var port_assets = "8100"

func main() {
	
	directory := "./assets"
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, port_assets)
	log.Fatal(http.ListenAndServe(":" + port_assets, nil))
}