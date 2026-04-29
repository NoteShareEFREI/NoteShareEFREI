package main

import (
	"log"
	"net/http"
	"NoteShareEFREI/routers"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("files"))))
    http.HandleFunc("/", routers.Handler)
	http.HandleFunc("/home", routers.HomeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
