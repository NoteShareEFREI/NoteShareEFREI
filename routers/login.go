package routers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	backend_url := "api/login"
	p, err := GetTemplate("login")
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
