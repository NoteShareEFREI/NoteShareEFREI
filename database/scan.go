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
