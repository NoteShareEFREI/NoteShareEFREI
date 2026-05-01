package main

import (
	"NoteShareEFREI/api"
	"NoteShareEFREI/backend"
	"NoteShareEFREI/database"
	"NoteShareEFREI/routers"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database Setup
	var err error

	database.Db, err = database.ConnectToDBTcp()
	if err != nil {
		println(err.Error())
		return
	}
	defer database.Db.Close()

	err = database.Db.Ping()
	if err != nil {
		println(err.Error())
		return
	}

	

	backend.Setup()
	//Global http requests
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("files"))))

	//Router end points
    http.HandleFunc("/", routers.Handler)
	http.HandleFunc("/home", routers.HomeHandler)
	http.HandleFunc("/createsheet", routers.CreateSheetHandler)
	http.HandleFunc("/login", routers.LoginHandler)

	http.Handle("/account", backend.Accountmiddleware(http.HandlerFunc(routers.AccountHandler)))
	//API end points (http responses with no html)
	http.HandleFunc("/api/login", api.LoginHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
