package api

import (
	"NoteShareEFREI/backend"
	"fmt"
	"net/http"

	regexp2 "github.com/dlclark/regexp2/v2"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
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
		email := r.Form.Get("mail")
		Phone := r.Form.Get("phone")

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

		//Hash the password to not store it plain.
		hash, salt, err := backend.NewHash(hash, name)
		if err != nil {

		}

		//Debug print to get all the values to be put inside the database.
		fmt.Println(name, hash, email, Phone, salt)

		i := 0 //Counter for the loop
		for while := true; while; while = err != nil && i < 5 {
			//result, err = database.doquery(insert into account values(....))
			i++
		}
		if err != nil {
			//Even after retrying multiple times, we couldn't store the info in the database.
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//Databese.doquery("Get Id as a result of the creation")

		id := 501
		//Add validation and authentification cookies.
		jwt := backend.GenerateCookieWithJWT(backend.GenerateJWT(id))
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
