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
	http.HandleFunc("/signup", routers.CreateHandler)
	http.Handle("/account", backend.Accountmiddleware(http.HandlerFunc(routers.AccountHandler)))

	//API end points (http responses with no html)
	http.HandleFunc("/api/login", api.LoginHandler)
	http.HandleFunc("/api/signup", api.CreateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	//http.ListenAndServeTLS() //Serve over https (need to create certificate and key with openssl)(Requires 'personal' information and the domain name on which the website is hsoted (FQDN))
}
