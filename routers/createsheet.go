package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"crypto/sha256"
)

func CreateSheetHandler(w http.ResponseWriter, r *http.Request) {
	//TODO return the errors to the web page
	//TODO verify authentication
	
	if (r.Method == "POST"){
		file, h, err := r.FormFile("studysheet")
		if err != nil {
			fmt.Print("No file sent")
			problemwithsheet(w,r)
			return
		}
		fmt.Print(path.Ext(h.Filename))
		if path.Ext(h.Filename) != ".pdf" && path.Ext(h.Filename) != ".md"{
			fmt.Print("Wrong format, only pdf and markdown accepted")
			problemwithsheet(w,r)
			return
		}

		if h.Size >= 80000000{
			fmt.Print("File is more than 10Mb")
			problemwithsheet(w,r)
			return
		}
		hash := sha256.New()
		if _, err := io.Copy(hash, file); err != nil {
			fmt.Print("unable to compute name of file for database")
			problemwithsheet(w,r)
			return
		}
		filename := fmt.Sprintf("%x", hash.Sum(nil))
		newfile, err := os.Create("files/"+filename)
		if err != nil {
			fmt.Print(err)
			return
		}
		_, err = file.Seek(0,0)
		if err != nil {
			fmt.Print("Seek didn't function on the file to read it a second time")
			return
		}
		sizewritten, err := io.Copy(newfile,file)
		if err != nil {
			fmt.Print("Copy failed")
			return 
		}
		
		if sizewritten < h.Size {
			fmt.Print("uncomplete copy")
			return
		}

		page_path := "templates/createsheetsuccess"
		p, err := os.ReadFile(page_path)
		if err != nil {
			fmt.Print("Error reading createsheetsuccess")
			return 
		}
		fmt.Fprintf(w, "%s", p)
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func problemwithsheet(w http.ResponseWriter, r *http.Request) {
		page_path := "templates/createsheetfailed"
		p, err := os.ReadFile(page_path)
		if err != nil {
			fmt.Print("Error reading createsheetfailed")
			return 
		}
		fmt.Fprintf(w, "%s", p)
}
