package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"NoteShareEFREI/backend"
	"NoteShareEFREI/database"

	"github.com/lestrrat-go/jwx/v4/jwt"
)

type studySheet struct {
	Title       string
	Category    string
	Description string
	Hash        string
	SubCategory string
}

type CategoryWithSubs struct {
	database.Category
	SubCategories []database.SubCategory
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p, err := GetTemplate("home")
	if err != nil {
		http.Error(w, "Failed to load home page", http.StatusInternalServerError)
		return
	}

	// Check authentication
	token, err := jwt.ParseRequest(r, jwt.WithCookieKey("Http-Jwt"), jwt.WithVerify(false))
	isLoggedIn := false
	isAdmin := false
	var accountId int
	if err == nil {
		accId, err := backend.ValidateJWT(token)
		if err == nil {
			isLoggedIn = true
			accountId = accId
			// Check if user is admin
			admin, err := database.IsAdmin(accountId)
			if err == nil {
				isAdmin = admin
			}
		}
	}

	// Fetch categories and subcategories
	cats, err := database.GetCategories()
	if err != nil {
		http.Error(w, "Failed to load categories", http.StatusInternalServerError)
		return
	}
	subs, err := database.GetSubCategories()
	if err != nil {
		http.Error(w, "Failed to load subcategories", http.StatusInternalServerError)
		return
	}
	// If no subcategories, insert initial data
	if len(cats) == 0 || len(subs) == 0 {
		lastIdInsert, err := database.InsertCategory("TestCategory")
		if err != nil {
			fmt.Println("Error inserting Testcategory category:", err.Error())
		}
		_, err = database.InsertSubCategory("TestSubCategory", lastIdInsert)
		if err != nil {
			fmt.Println("Error inserting TestSubcategory category:", err.Error())
		}
		// Refetch after insert
		cats, err = database.GetCategories()
		if err != nil {
			http.Error(w, "Failed to load categories", http.StatusInternalServerError)
			return
		}
		subs, err = database.GetSubCategories()
		if err != nil {
			http.Error(w, "Failed to load subcategories", http.StatusInternalServerError)
			return
		}
	}
	subCatsJSON, _ := json.Marshal(subs)

	query := strings.TrimSpace(r.URL.Query().Get("q"))
	selectedCategory := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("category")))
	normalizedQuery := strings.ToLower(query)

	// Query actual study sheets from database
	sqlQuery := `
		SELECT StudySheet.Name, Category.Name AS CatName, SubCategory.Name AS SubName, StudySheet.Hash
		FROM StudySheet
		INNER JOIN SubCategory ON StudySheet.Id_SubCategory = SubCategory.Id_SubCategory
		INNER JOIN Category ON SubCategory.Id_Category = Category.Id_Category
		WHERE (? = '' OR LOWER(Category.Name) LIKE ? OR LOWER(SubCategory.Name) LIKE ? OR LOWER(StudySheet.Name) LIKE ?)
		AND (? = '' OR LOWER(Category.Name) = ?)
	`
	likeQuery := "%" + normalizedQuery + "%"
	rows, err := database.Db.Query(sqlQuery, query, likeQuery, likeQuery, likeQuery, selectedCategory, selectedCategory)
	if err != nil {
		http.Error(w, "Failed to query study sheets", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	results := make([]studySheet, 0)
	for rows.Next() {
		var name, catName, subName, hash string
		err = rows.Scan(&name, &catName, &subName, &hash)
		if err != nil {
			http.Error(w, "Failed to scan study sheet", http.StatusInternalServerError)
			return
		}
		results = append(results, studySheet{
			Title:       name,
			Category:    catName,
			Description: subName,
			Hash:        hash,
			SubCategory: subName,
		})
	}

	// Prepare categories with subcategories for template
	categoriesWithSubs := make([]CategoryWithSubs, len(cats))
	for i, cat := range cats {
		categoriesWithSubs[i] = CategoryWithSubs{
			Category:      cat,
			SubCategories: make([]database.SubCategory, 0),
		}
	}
	for _, sub := range subs {
		for i := range categoriesWithSubs {
			if categoriesWithSubs[i].Id == sub.CategoryId {
				categoriesWithSubs[i].SubCategories = append(categoriesWithSubs[i].SubCategories, sub)
				break
			}
		}
	}

	data := struct {
		Query              string
		SelectedCategory   string
		Results            []studySheet
		HasFilters         bool
		IsLoggedIn         bool
		IsAdmin            bool
		Categories         []database.Category
		SubCategories      []database.SubCategory
		CategoriesWithSubs []CategoryWithSubs
		SubCategoriesJSON  string
	}{
		Query:              query,
		SelectedCategory:   selectedCategory,
		Results:            results,
		HasFilters:         query != "" || selectedCategory != "",
		IsLoggedIn:         isLoggedIn,
		IsAdmin:            isAdmin,
		Categories:         cats,
		SubCategories:      subs,
		CategoriesWithSubs: categoriesWithSubs,
		SubCategoriesJSON:  string(subCatsJSON),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = p.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "Http-Jwt",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
