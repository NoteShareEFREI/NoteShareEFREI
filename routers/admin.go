package routers

import (
	"NoteShareEFREI/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	acc_ID, ok := r.Context().Value("Account ID").(int)
	if !ok || acc_ID == -1 {
		http.Error(w, "Unauthorized: You must be logged in", http.StatusUnauthorized)
		return
	}

	// Check if user is admin
	isAdmin, err := database.IsAdmin(acc_ID)
	if err != nil || !isAdmin {
		http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
		return
	}

	// Handle form submission
	if r.Method == http.MethodPost {
		action := r.FormValue("action")

		switch action {
		case "add_category":
			categoryName := r.FormValue("category_name")
			categoryName = strings.TrimSpace(categoryName)
			if categoryName == "" {
				// Handle error, perhaps redirect with error message
				http.Redirect(w, r, "/admin?error=Category name cannot be empty", http.StatusSeeOther)
				return
			}
			_, err = database.InsertCategory(categoryName)
			if err != nil {
				fmt.Println("Error adding category:", err.Error())
			}
		case "add_subcategory":
			categoryIdStr := r.FormValue("category_id")
			subcategoryName := r.FormValue("subcategory_name")
			subcategoryName = strings.TrimSpace(subcategoryName)
			if categoryIdStr == "" || subcategoryName == "" {
				http.Redirect(w, r, "/admin?error=Category ID and subcategory name are required", http.StatusSeeOther)
				return
			}
			categoryId, err := strconv.Atoi(categoryIdStr)
			if err != nil {
				http.Redirect(w, r, "/admin?error=Invalid category ID", http.StatusSeeOther)
				return
			}
			_, err = database.InsertSubCategory(subcategoryName, categoryId)
			if err != nil {
				fmt.Println("Error adding subcategory:", err.Error())
			}
		case "delete_category":
			categoryIdStr := r.FormValue("category_id")
			categoryId, err := strconv.Atoi(categoryIdStr)
			if err != nil {
				http.Redirect(w, r, "/admin?error=Invalid category ID", http.StatusSeeOther)
				return
			}
			err = database.DeleteCategoryWithSubcategories(categoryId)
			if err != nil {
				fmt.Println("Error deleting category and subcategories:", err.Error())
			}
		case "delete_subcategory":
			subcategoryIdStr := r.FormValue("subcategory_id")
			subcategoryId, err := strconv.Atoi(subcategoryIdStr)
			if err != nil {
				http.Redirect(w, r, "/admin?error=Invalid subcategory ID", http.StatusSeeOther)
				return
			}
			err = database.DeleteSubCategory(subcategoryId)
			if err != nil {
				fmt.Println("Error deleting subcategory:", err.Error())
			}
		}

		// Redirect to avoid duplicate submission
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	// Fetch all categories and subcategories
	categories, err := database.GetCategories()
	if err != nil {
		http.Error(w, "Failed to load categories", http.StatusInternalServerError)
		return
	}

	subcategories, err := database.GetSubCategories()
	if err != nil {
		http.Error(w, "Failed to load subcategories", http.StatusInternalServerError)
		return
	}

	// If no categories exist, insert initial data
	if len(categories) == 0 || len(subcategories) == 0 {
		// Insert categories
		lastIdInsert, err := database.InsertCategory("TestCategory")
		if err != nil {
			fmt.Println("Error inserting Testcategory category:", err.Error())
		}
		_, err = database.InsertSubCategory("TestSubCategory", lastIdInsert)
		if err != nil {
			fmt.Println("Error inserting TestSubcategory category:", err.Error())
		}

		// Refetch categories after insert
		categories, err = database.GetCategories()
		if err != nil {
			http.Error(w, "Failed to load categories after insert", http.StatusInternalServerError)
			return
		}
	}

	data := struct {
		Categories    []database.Category
		SubCategories []database.SubCategory
	}{
		Categories:    categories,
		SubCategories: subcategories,
	}

	p, err := GetTemplate("admin")
	if err != nil {
		fmt.Println("Error getting template:", err.Error())
		http.Error(w, "Failed to load admin page", http.StatusInternalServerError)
		return
	}

	err = p.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
		http.Error(w, "Failed to render admin page", http.StatusInternalServerError)
		return
	}
}
