package routers

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"NoteShareEFREI/backend"
	"NoteShareEFREI/database"

	"github.com/lestrrat-go/jwx/v4/jwt"
)

func CreateSheetHandler(w http.ResponseWriter, r *http.Request) {
	// Check authentication
	token, err := jwt.ParseRequest(r, jwt.WithCookieKey("Http-Jwt"), jwt.WithVerify(false))
	var accId int
	isLoggedIn := false
	if err == nil {
		accId, err = backend.ValidateJWT(token)
		if err == nil {
			isLoggedIn = true
		}
	}
	if !isLoggedIn {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		return
	}

	if r.Method == "POST" {
		// Get form values
		title := r.FormValue("title")
		catIdStr := r.FormValue("category")
		subName := r.FormValue("subcategory")

		if title == "" || catIdStr == "" || subName == "" {
			fmt.Print("Missing form fields")
			problemwithsheet(w, r)
			return
		}

		// Get category name
		catId := 0
		_, err2 := fmt.Sscanf(catIdStr, "%d", &catId)
		if err2 != nil {
			return
		}
		catRow := database.Db.QueryRow("SELECT Name FROM Category WHERE Id_Category = ?", catId)
		var catName string
		err = catRow.Scan(&catName)
		if err != nil {
			fmt.Print("Invalid category")
			problemwithsheet(w, r)
			return
		}

		// Get subcategory ID
		subRow := database.Db.QueryRow("SELECT Id_SubCategory FROM SubCategory WHERE Name = ? AND Id_Category = ?", subName, catId)
		var subId int
		err = subRow.Scan(&subId)
		if err != nil {
			fmt.Print("Invalid subcategory")
			problemwithsheet(w, r)
			return
		}

		file, h, err := r.FormFile("studysheet")
		if err != nil {
			fmt.Print("No file sent")
			problemwithsheet(w, r)
			return
		}
		fmt.Print(path.Ext(h.Filename))
		if path.Ext(h.Filename) != ".pdf" && path.Ext(h.Filename) != ".md" {
			fmt.Print("Wrong format, only pdf and markdown accepted")
			problemwithsheet(w, r)
			return
		}

		if h.Size >= 80000000 {
			fmt.Print("File is more than 10Mb")
			problemwithsheet(w, r)
			return
		}
		hash := sha256.New()
		if _, err := io.Copy(hash, file); err != nil {
			fmt.Print("unable to compute name of file for database")
			problemwithsheet(w, r)
			return
		}
		filename := fmt.Sprintf("%x", hash.Sum(nil))

		// Create directory
		dir := fmt.Sprintf("files/%s/%s", catName, subName)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Print("Failed to create directory")
			problemwithsheet(w, r)
			return
		}

		newfile, err := os.Create(dir + "/" + filename)
		if err != nil {
			fmt.Print(err)
			problemwithsheet(w, r)
			return
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Print("Seek didn't function on the file to read it a second time")
			return
		}
		sizewritten, err := io.Copy(newfile, file)
		if err != nil {
			fmt.Print("Copy failed")
			return
		}

		if sizewritten < h.Size {
			fmt.Print("uncomplete copy")
			problemwithsheet(w, r)
			return
		}
		err := newfile.Close()
		if err != nil {
			return
		}

		// Insert into database
		nextId, err := database.GetNextSheetId()
		if err != nil {
			fmt.Print("Failed to get next sheet ID")
			problemwithsheet(w, r)
			return
		}
		_, err = database.Newstudysheet(nextId, filename, title, subId, accId)
		if err != nil {
			fmt.Print("Failed to insert into database")
			problemwithsheet(w, r)
			return
		}

		page_path := "templates/createsheetsuccess"
		p, err := os.ReadFile(page_path)
		if err != nil {
			fmt.Print("Error reading createsheetsuccess")
			return
		}
		_, err = fmt.Fprintf(w, "%s", p)
		if err != nil {
			return
		}
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func problemwithsheet(w http.ResponseWriter) {
	pagePath := "templates/createsheetfailed"
	p, err := os.ReadFile(pagePath)
	if err != nil {
		fmt.Print("Error reading createsheetfailed")
		return
	}
	_, err = fmt.Fprintf(w, "%s", p)
	if err != nil {
		return
	}
}
