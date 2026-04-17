package routers

import (
	"fmt"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page_path := "templates/home" //r.URL.Path[len("/templates/"):]
	p, err := os.ReadFile(page_path)
	fmt.Print(p)
	if err != nil {
		fmt.Print("Error reading home")
		return 
	}
	fmt.Fprintf(w, "%s", p)
}
