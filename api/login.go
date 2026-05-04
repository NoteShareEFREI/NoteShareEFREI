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
		err := Addredirect(w, "login", http.StatusBadRequest)
		if err != nil {
			return
		} //400
		fmt.Print(err.Error())
		return
	}

	//Verify if the user is sending the right type of request.
	switch r.Method {
	case "POST":
		hash := r.Form.Get("pwd")
		name := r.Form.Get("name")

		correct, id := backend.VerifyPerson(hash, name)
		if !correct {
			if id != -1 {
				fmt.Println("The `backend.VerifyPerson` function returned with an error of code, code =", id)
			}
			err := Addredirect(w, "login", http.StatusForbidden)
			if err != nil {
				return
			} //403
			return
		}

		//Add validation and authentification cookies.
		jwt := backend.GenerateCookieWithJWT(backend.GenerateJWT(id))
		http.SetCookie(w, &jwt)

		err := Addredirect(w, "home", http.StatusSeeOther)
		if err != nil {
			return
		} //303

	default:
		err := Addredirect(w, "login", http.StatusNotFound)
		if err != nil {
			return
		} //404
		_, err = w.Write([]byte(`{"message": "not found"}`))
		if err != nil {
			return
		}
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
