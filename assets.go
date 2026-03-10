package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := "8100"
	directory := "./assets"
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}