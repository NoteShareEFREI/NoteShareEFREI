package main

import (
	"log"
	"net/http"
	"NoteShareEFREI/routers"
)

func main() {
    http.HandleFunc("/", routers.Handler)
	http.HandleFunc("/home", routers.HomeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
