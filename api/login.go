package api

import (
	crypto "NoteShareEFREI/backend"
	"fmt"
	"net/http"

	regexp2 "github.com/dlclark/regexp2/v2"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := r.ParseForm()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	//Verify if the user is sending the right type of request.
	switch r.Method {
	case "POST":
		hash := r.Form.Get("pwd")
		name := r.Form.Get("name")

		//Here the original go package (regexp) does not support positive lookahead so i need to use another package.
		re := regexp2.MustCompile(`(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&\|:\{\}£¬_+#\[\]^\(~\)\-])(?=.*[^'.;,\\])[A-Za-z\d@#$!\[\]%*^?&\|\(~\):\{\}£¬_+\-]{8,31}`)
		isMatch, err := re.MatchString(hash)
		if err != nil {
			panic(err.Error() + "\n Error in api.login, POST request.")
		}
		if isMatch == false {
			//If an hacker tried to bypass the regex on the page and send an unsecure password.
			w.WriteHeader(http.StatusBadRequest) //400
			return
		}

		fmt.Println(name, hash) //To verify informations sent (Debug)

		/*correct, id := crypto.VerifyPerson(hash, name)
		if !correct {
			w.WriteHeader(http.StatusForbidden) //403
			return
		}*/

		//Store the results in the database. For Create Account.
		//database.Doquery("A query")

		id := 4 //For testing purposes, to be deleted.
		//Add validation and authentification cookies.
		jwt := crypto.GenerateCookieWithJWT(crypto.GenerateJWT(id))
		fmt.Print(jwt.Valid())
		http.SetCookie(w, &jwt)

		//Redirect to home page
		w.Header().Set("Location", "/home")
		w.Header().Set("Refresh", " 0; url=/home")

		w.WriteHeader(http.StatusSeeOther) //303

	default:
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte(`{"message": "not found"}`))
	}

	//To send everyone that makes a request to this page back to home page.
	meta_redirection := "<meta http-equiv='refresh' content='0;url=/home'>"
	fmt.Fprintf(w, "%s", meta_redirection)
}
