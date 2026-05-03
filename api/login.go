package api

import (
	"NoteShareEFREI/backend"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := r.ParseForm()
	if err != nil {
		Addredirect(w, "login", http.StatusBadRequest) //400
		fmt.Print(err.Error())
		return
	}

	//Verify if the user is sending the right type of request.
	switch r.Method {
	case "POST":
		hash := r.Form.Get("pwd")
		name := r.Form.Get("name")

		//The regex checking was removed since it is too costly and can take up to a minute.

		correct, id := backend.VerifyPerson(hash, name)
		if !correct {
			Addredirect(w, "login", http.StatusForbidden) //403
			if id != -1 {
				fmt.Println("The `backend.VerifyPerson` function returned with an error of code, code =", id)
			}
			return
		}

		//Add validation and authentification cookies.
		jwt := backend.GenerateCookieWithJWT(backend.GenerateJWT(id))
		http.SetCookie(w, &jwt)

		Addredirect(w, "home", http.StatusSeeOther) //303

	default:
		Addredirect(w, "login", http.StatusNotFound) //404
		w.Write([]byte(`{"message": "not found"}`))
	}

}

func Addredirect(w http.ResponseWriter, page string, code int) error {
	w.Header().Set("Location", "/"+page)
	w.Header().Set("Refresh", " 0; url=/"+page)
	w.WriteHeader(code)
	meta_redirection := fmt.Sprintf("<meta http-equiv='refresh' content='0;url=/%s'>", page)
	_, err := fmt.Fprintf(w, "%s", meta_redirection)
	return err
}
