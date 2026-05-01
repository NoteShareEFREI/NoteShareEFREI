package main

import (
	"NoteShareEFREI/api"
	"NoteShareEFREI/backend"
	"NoteShareEFREI/routers"
	"log"
	"net/http"
)

func main() {
	//Initialize the JWT private keys.
	backend.Setup()

	//Global http requests
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/css/styles.css", http.NotFoundHandler()) //Handler not present.

	//Router end points
	http.HandleFunc("/", routers.Handler)
	http.HandleFunc("/home", routers.HomeHandler)
	http.HandleFunc("/login", routers.LoginHandler)
	http.Handle("/account", backend.Accountmiddleware(http.HandlerFunc(routers.AccountHandler)))

	//API end points (http responses with no html)
	http.HandleFunc("/api/login", api.LoginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
