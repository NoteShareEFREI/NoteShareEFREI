package routers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	page_path := "templates/account"
	acc_ID := r.Context().Value("Account ID").(int)
	if acc_ID == -1 {
		//The user is not logged in so we use a redirection page proposing that he creates an account.
		page_path = "templates/account_redirect"
		w.WriteHeader(http.StatusOK)
	}
	p, err := template.ParseFiles(page_path)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	data := struct {
		Username string
		Mail     string
		Phone    string
	}{
		Username: `%Error%`, //It should be impossible to not have a username.
		Mail:     "None",
		Phone:    "None",
	}

	//database.Doquery("hello world") // And put the results in data

	// Need To put the page inside another template

	err = p.Execute(w, data)

	if err != nil {
		fmt.Println("Struct data is bad")
		fmt.Println(err.Error())
		return
	}
}
