package routers

import (
	"fmt"
	"html/template"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	page_path := "templates/create_account"
	backend_url := "api/signup"
	p, err := template.ParseFiles(page_path)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	data := struct {
		Actionurl              string
		Eventual_error_display string
	}{
		Actionurl:              backend_url,
		Eventual_error_display: "",
	}

	// To put the page inside another template

	err = p.Execute(w, data)

	if err != nil {
		fmt.Print("Struct data is bad")
		fmt.Print(err.Error())
		return
	}
}
