package routers

import (
	"html/template"
	"net/http"
	"strings"
)

type studySheet struct {
	Title       string
	Category    string
	Description string
}

var homeStudySheets = []studySheet{
	{
		Title:"Linear Algebra",
		Category:"maths",
		Description:"Matrix, Euler",
	},
	{
		Title: "Calculus",
		Category:    "maths",
		Description: "Integrals",
	},
	{
		Title:"Data Structures in C",
		Category: "programming",
		Description: "Arrays, linked lists...",
	},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pagePath := "templates/home"
	p, err := template.ParseFiles(pagePath)
	if err != nil {
		http.Error(w, "Failed to load home page", http.StatusInternalServerError)
		return
	}

	query := strings.TrimSpace(r.URL.Query().Get("q"))
	selectedCategory := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("category")))
	normalizedQuery := strings.ToLower(query)

	results := make([]studySheet, 0, len(homeStudySheets))
	for _, sheet := range homeStudySheets {
		if selectedCategory != "" && selectedCategory != strings.ToLower(sheet.Category) {
			continue
		}
		if normalizedQuery != "" {
			if !strings.Contains(strings.ToLower(sheet.Title), normalizedQuery) &&
				!strings.Contains(strings.ToLower(sheet.Description), normalizedQuery) &&
				!strings.Contains(strings.ToLower(sheet.Category), normalizedQuery) {
				continue}
		}
		results = append(results, sheet)
	}

	data := struct {
		Query string
		SelectedCategory string
		Results []studySheet
		HasFilters bool
	}{
		Query: query,
		SelectedCategory: selectedCategory,
		Results:  results,
		HasFilters: query != "" || selectedCategory != "",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = p.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
		return
	}
}
