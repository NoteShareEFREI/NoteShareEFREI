package api

import (
	"NoteShareEFREI/backend"
	"NoteShareEFREI/database"
	"fmt"
	"net/http"

	regexp2 "github.com/dlclark/regexp2/v2"
)

var re *regexp2.Regexp

func Initialize() {
	//Here the original go package (regexp) does not support positive lookahead so i need to use another package.
	re = regexp2.MustCompile(`(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&\|:\{\}£¬_+#\[\]^\(~\)\-])(?=.*[^'.;,\\])[A-Za-z\d@#$!\[\]%*^?&\|\(~\):\{\}£¬_+\-]{8,31}`)
	//Stops from recompiling everytime (3-5 seconds)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	err := r.ParseForm()
	if err != nil {
		Addredirect(w, "signup", http.StatusBadRequest) //400
		fmt.Print(err.Error())
		return
	}

	//Verify if the user is sending the right type of request.
	switch r.Method {
	case "POST":
		hash := r.Form.Get("pwd")
		name := r.Form.Get("name")
		email := r.Form.Get("mail")
		Phone := r.Form.Get("phone")

		isMatch, err := re.MatchString(hash) //May take some time to run.
		if err != nil {
			fmt.Println(err.Error() + "\n Error in api.login, POST request.")
			isMatch = false
		}
		if isMatch == false {
			//If an hacker tried to bypass the regex on the page and send an unsecure password.
			Addredirect(w, "signup", http.StatusBadRequest) //400
			return
		}

		//Hash the password to not store it plain.
		hash, salt, err := backend.NewHash(hash, name)
		if err != nil {
			Addredirect(w, "signup", http.StatusInternalServerError)
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
				Addredirect(w, "signup", http.StatusInternalServerError)
				return
			}
		}

		//Get the Account ID of the created account
		id, err := database.Getidfrompseudoandhash(name, hash)
		if err != nil {
			Addredirect(w, "signup", http.StatusInternalServerError)
			return
		}

		//Add validation and authentification cookies.
		jwt := backend.GenerateCookieWithJWT(backend.GenerateJWT(id))
		http.SetCookie(w, &jwt)

		Addredirect(w, "home", http.StatusSeeOther)

	default:
		Addredirect(w, "create", http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

}
