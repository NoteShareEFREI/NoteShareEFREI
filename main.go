package main

import (
	"log"
	"net/http"
	"NoteShareEFREI/routers"
)

func main() {
    http.HandleFunc("/", routers.Handler)
	http.HandleFunc("/home", routers.HomeHandler)
	http.HandleFunc("/createsheet", routers.CreateSheetHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
