package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Printf("Serving %s on HTTP port: %s\n", "./public", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
