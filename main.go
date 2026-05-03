package main

import (
	"NoteShareEFREI/database"
	"NoteShareEFREI/routers"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
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

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("files"))))
    http.HandleFunc("/", routers.Handler)
	http.HandleFunc("/home", routers.HomeHandler)
	http.HandleFunc("/createsheet", routers.CreateSheetHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
