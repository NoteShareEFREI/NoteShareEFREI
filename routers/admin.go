package routers

import (
	"NoteShareEFREI/database"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
			if categoryName != "" {
				nextId, err := database.GetMaxCategoryId()
				if err != nil {
					nextId = 1
				}
				_, err = database.InsertCategory(nextId, categoryName)
				if err != nil {
					fmt.Println("Error adding category:", err.Error())
				}
			}
		case "add_subcategory":
			categoryIdStr := r.FormValue("category_id")
			subcategoryName := r.FormValue("subcategory_name")
			if categoryIdStr != "" && subcategoryName != "" {
				categoryId, err := strconv.Atoi(categoryIdStr)
				if err == nil {
					nextId, err := database.GetMaxSubCategoryId()
					if err != nil {
						nextId = 1
					}
					_, err = database.InsertSubCategory(nextId, subcategoryName, categoryId)
					if err != nil {
						fmt.Println("Error adding subcategory:", err.Error())
					}
				}
			}
		case "delete_category":
			categoryIdStr := r.FormValue("category_id")
			categoryId, err := strconv.Atoi(categoryIdStr)
			if err == nil {
				// First delete all subcategories of this category
				err = database.DeleteSubCategoriesByCategoryId(categoryId)
				if err != nil {
					fmt.Println("Error deleting subcategories:", err.Error())
				}
				// Then delete the category
				err = database.DeleteCategory(categoryId)
				if err != nil {
					fmt.Println("Error deleting category:", err.Error())
				}
			}
		case "delete_subcategory":
			subcategoryIdStr := r.FormValue("subcategory_id")
			subcategoryId, err := strconv.Atoi(subcategoryIdStr)
			if err == nil {
				err = database.DeleteSubCategory(subcategoryId)
				if err != nil {
					fmt.Println("Error deleting subcategory:", err.Error())
				}
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
	if len(categories) == 0 {
		// Insert categories
		_, err = database.InsertCategory(1, "Maths")
		if err != nil {
			fmt.Println("Error inserting Maths category:", err.Error())
		}
		_, err = database.InsertCategory(2, "Physics")
		if err != nil {
			fmt.Println("Error inserting Physics category:", err.Error())
		}
		_, err = database.InsertCategory(3, "Programming")
		if err != nil {
			fmt.Println("Error inserting Programming category:", err.Error())
		}
		_, err = database.InsertCategory(4, "Formation_Generale")
		if err != nil {
			fmt.Println("Error inserting Formation_Generale category:", err.Error())
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

	p, err := template.ParseFiles("templates/admin")
	if err != nil {
		fmt.Println("Error parsing template:", err.Error())
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
