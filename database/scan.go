package database

import (
	"database/sql"
	"errors"
)

func Onerowint(row *sql.Rows) (int, error) {
	if row.Next() {
		var id int
		err := row.Scan(&id)
		if err != nil {
			return -1, err
		}
		if row.Next() {
			return -1, errors.New("There should be only 1 row returned. (2+ found)")
		}
		return id, nil
	}
	return -1, errors.New("There should be 1 row returned. (0 found)")
}

func Getidfrompseudoandhash(pseudo string, hash string) (int, error) {
	rows, err := Db.Query("Select Id_Account from Account where Pseudo=? and HashPassword=?", pseudo, hash)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	id, err := Onerowint(rows)
	if err != nil {
		return -1, err
	}
	return id, nil
}

type Category struct {
	Id   int
	Name string
}

type SubCategory struct {
	Id         int
	Name       string
	CategoryId int
}

type Comment struct {
	Id        int
	Content   string
	SheetId   int
	AccountId int
	Pseudo    string // To include username
}

func GetCategories() ([]Category, error) {
	rows, err := Db.Query("SELECT Id_Category, Name FROM Category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cats []Category
	for rows.Next() {
		var c Category
		err = rows.Scan(&c.Id, &c.Name)
		if err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, nil
}

func GetSubCategories() ([]SubCategory, error) {
	rows, err := Db.Query("SELECT Id_SubCategory, Name, Id_Category FROM SubCategory")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subs []SubCategory
	for rows.Next() {
		var s SubCategory
		err = rows.Scan(&s.Id, &s.Name, &s.CategoryId)
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func GetNextSheetId() (int, error) {
	rows, err := Db.Query("SELECT MAX(Id_Sheet) FROM StudySheet")
	if err != nil {
		return 1, err
	}
	defer rows.Close()
	if rows.Next() {
		var maxId sql.NullInt64
		err = rows.Scan(&maxId)
		if err != nil {
			return 1, err
		}
		if maxId.Valid {
			return int(maxId.Int64) + 1, nil
		}
	}
	return 1, nil
}

// IsAdmin checks if a user is an admin (Role = 0)
func IsAdmin(accountId int) (bool, error) {
	rows, err := Db.Query("SELECT Role FROM Account WHERE Id_Account = ?", accountId)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		var role int
		err := rows.Scan(&role)
		if err != nil {
			return false, err
		}
		return role == 0, nil
	}
	return false, errors.New("Account not found")
}

// GetMaxCategoryId returns the next available category ID
func GetMaxCategoryId() (int, error) {
	rows, err := Db.Query("SELECT MAX(Id_Category) FROM Category")
	if err != nil {
		return 1, err
	}
	defer rows.Close()
	if rows.Next() {
		var maxId sql.NullInt64
		err = rows.Scan(&maxId)
		if err != nil {
			return 1, err
		}
		if maxId.Valid {
			return int(maxId.Int64) + 1, nil
		}
	}
	return 1, nil
}

// DeleteCategory deletes a category by ID
func DeleteCategory(categoryId int) error {
	result, err := Db.Exec("DELETE FROM Category WHERE Id_Category = ?", categoryId)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}

// UpdateCategory updates a category name
func UpdateCategory(categoryId int, name string) error {
	_, err := Db.Exec("UPDATE Category SET Name = ? WHERE Id_Category = ?", name, categoryId)
	return err
}

// GetMaxSubCategoryId returns the next available subcategory ID
func GetMaxSubCategoryId() (int, error) {
	rows, err := Db.Query("SELECT MAX(Id_SubCategory) FROM SubCategory")
	if err != nil {
		return 1, err
	}
	defer rows.Close()
	if rows.Next() {
		var maxId sql.NullInt64
		err = rows.Scan(&maxId)
		if err != nil {
			return 1, err
		}
		if maxId.Valid {
			return int(maxId.Int64) + 1, nil
		}
	}
	return 1, nil
}

// DeleteSubCategory deletes a subcategory by ID
func DeleteSubCategory(subcategoryId int) error {
	result, err := Db.Exec("DELETE FROM SubCategory WHERE Id_SubCategory = ?", subcategoryId)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}

// DeleteSubCategoriesByCategoryId deletes all subcategories for a given category ID
func DeleteSubCategoriesByCategoryId(categoryId int) error {
	_, err := Db.Exec("DELETE FROM SubCategory WHERE Id_Category = ?", categoryId)
	return err
}

// GetSheetIdByHash returns the Id_Sheet for a given hash
func GetSheetIdByHash(hash string) (int, error) {
	rows, err := Db.Query("SELECT Id_Sheet FROM StudySheet WHERE Hash = ?", hash)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	return Onerowint(rows)
}

// GetCommentsBySheetId returns all comments for a sheet, including pseudo
func GetCommentsBySheetId(sheetId int) ([]Comment, error) {
	rows, err := Db.Query(`
		SELECT Comment.Id_Comment, Comment.Content, Comment.Id_Sheet, Comment.Id_Account, Account.Pseudo
		FROM Comment
		JOIN Account ON Comment.Id_Account = Account.Id_Account
		WHERE Comment.Id_Sheet = ?
		ORDER BY Comment.Id_Comment ASC
	`, sheetId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.Id, &c.Content, &c.SheetId, &c.AccountId, &c.Pseudo)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

// InsertComment adds a new comment
func InsertComment(content string, sheetId int, accountId int) error {
	_, err := Db.Exec("INSERT INTO Comment (Content, Id_Sheet, Id_Account) VALUES (?, ?, ?)", content, sheetId, accountId)
	return err
}

// GetMaxCommentId returns the next available comment ID
func GetMaxCommentId() (int, error) {
	rows, err := Db.Query("SELECT MAX(Id_Comment) FROM Comment")
	if err != nil {
		return 1, err
	}
	defer rows.Close()
	if rows.Next() {
		var maxId sql.NullInt64
		err = rows.Scan(&maxId)
		if err != nil {
			return 1, err
		}
		if maxId.Valid {
			return int(maxId.Int64) + 1, nil
		}
	}
	return 1, nil
}
