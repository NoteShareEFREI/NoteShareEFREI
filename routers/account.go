package routers

import (
	"NoteShareEFREI/database"
	"database/sql"
	"fmt"
	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	acc_ID := r.Context().Value("Account ID").(int)
	var templateName string
	if acc_ID == -1 {
		//The user is not logged in so we use a redirection page proposing that he creates an account.
		templateName = "account_redirect"
		w.WriteHeader(http.StatusOK)
	} else {
		templateName = "account"
	}
	p, err := GetTemplate(templateName)
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

	row, err := database.Db.Query("Select Pseudo, Email, Phonenumber from Account where Id_Account=?", acc_ID)
	if err != nil {
		return
	}
	if row.Next() {
		var mail sql.NullString
		var phone sql.NullString
		err := row.Scan(&data.Username, &mail, &phone)
		if err != nil {
			return
		}
		if mail.Valid {
			data.Mail = mail.String
		}
		if phone.Valid {
			data.Phone = phone.String
		}
	} else {
		fmt.Println("Error fetching information for user account id:", acc_ID, "No rows were returned.")
	}

	// Need To put the page inside another template

	err = p.Execute(w, data)

	if err != nil {
		fmt.Println("Struct data is bad")
		fmt.Println(err.Error())
		return
	}
}
