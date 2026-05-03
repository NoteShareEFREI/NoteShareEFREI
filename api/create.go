package api

import (
	"NoteShareEFREI/backend"
	"NoteShareEFREI/database"
	"fmt"
	"net/http"
	"unicode/utf8"
)

func isAllowedSpecial(c rune) bool {
	switch c {
	case '@', '$', '!', '%', '*', '?', '&', '|', ':',
		'{', '}', '£', '¬', '_', '+', '#', '[', ']', '^', '(',
		')', '-', '~':
		return true
	}
	return false
}

func ValidatePassword(pw string) bool {
	if utf8.RuneCountInString(pw) < 8 || utf8.RuneCountInString(pw) > 31 {
		return false
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, c := range pw {
		switch {
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= '0' && c <= '9':
			hasDigit = true
		case isAllowedSpecial(c):
			hasSpecial = true
		case c == '\'' || c == '.' || c == ',' || c == ';' || c == '\\':
			return false
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	err := r.ParseForm()
	if err != nil {
		err := Addredirect(w, "signup", http.StatusBadRequest)
		if err != nil {
			fmt.Print(err.Error())
			return
		} //400
		return
	}

	//Verify if the user is sending the right type of request.
	switch r.Method {
	case "POST":
		hash := r.Form.Get("pwd")
		name := r.Form.Get("name")
		email := r.Form.Get("mail")
		Phone := r.Form.Get("phone")

		isMatch := ValidatePassword(hash)

		if isMatch == false {
			//If an hacker tried to bypass the regex on the page and send an unsecure password.
			err := Addredirect(w, "signup", http.StatusBadRequest)
			if err != nil {
				return
			} //400
			return
		}

		//Hash the password to not store it plain.
		hash, salt, err := backend.NewHash(hash)
		if err != nil {
			err := Addredirect(w, "signup", http.StatusInternalServerError)
			if err != nil {
				return
			}
			return
		}

		//Try the query 5 times
		i := 0 //Counter for the loop
		for while := true; while; while = i < 5 {
			_, err := database.Newaccount(name, email, hash, Phone, salt)
			if err == nil {
				break
			}

			i++
			if i == 5 { // && err != nil (implicit, already verified)
				//Even after retrying multiple times, we couldn't store the info in the database.
				err := Addredirect(w, "signup", http.StatusInternalServerError)
				if err != nil {
					return
				}
				return
			}
		}

		//Get the Account ID of the created account
		id, err := database.Getidfrompseudoandhash(name, hash)
		if err != nil {
			err := Addredirect(w, "signup", http.StatusInternalServerError)
			if err != nil {
				return
			}
			return
		}

		//Add validation and authentification cookies.
		jwt := backend.GenerateCookieWithJWT(backend.GenerateJWT(id))
		http.SetCookie(w, &jwt)

		err = Addredirect(w, "home", http.StatusSeeOther)
		if err != nil {
			return
		}

	default:
		err := Addredirect(w, "create", http.StatusNotFound)
		if err != nil {
			return
		}
		_, err = w.Write([]byte(`{"message": "not found"}`))
		if err != nil {
			return
		}
	}

}
