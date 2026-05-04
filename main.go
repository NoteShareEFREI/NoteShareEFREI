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
	//Database Setup
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

	//Initialize the global values.
	backend.Setup()

	//Initialize template cache
	err = routers.InitializeTemplates()
	if err != nil {
		println("Failed to initialize templates:", err.Error())
		return
	}

	//Global http requests
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("files"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/pdfs/", http.StripPrefix("/pdfs/", http.FileServer(http.Dir("templates/pdfs"))))

	//Router end points
	http.HandleFunc("/", routers.Handler)
	http.HandleFunc("/home", routers.HomeHandler)
	http.HandleFunc("/createsheet", routers.CreateSheetHandler)
	http.HandleFunc("/login", routers.LoginHandler)
	http.HandleFunc("/signup", routers.CreateHandler)
	http.HandleFunc("/logout", routers.LogoutHandler)
	http.Handle("/account", backend.Accountmiddleware(http.HandlerFunc(routers.AccountHandler)))
	http.Handle("/admin", backend.Accountmiddleware(http.HandlerFunc(routers.AdminHandler)))

	//API end points (http responses with no html)
	http.HandleFunc("/api/login", api.LoginHandler)
	http.HandleFunc("/api/signup", api.CreateHandler)
	http.HandleFunc("/api/check-admin", api.CheckAdminHandler)
	http.HandleFunc("/api/comments", api.GetCommentsHandler)
	http.HandleFunc("/api/comments/add", api.PostCommentHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	//http.ListenAndServeTLS() //Serve over https (need to create certificate and key with openssl)(Requires 'personal' information and the domain name on which the website is hsoted (FQDN))
}
