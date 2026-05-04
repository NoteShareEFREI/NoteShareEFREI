package database

import (
	"database/sql"
)

func Newaccount(name string, email string, hash string, Phone string, salt int) (sql.Result, error) {
	//Store the query to be tried into a variable.
	var query string

	//We need to test all the possibilites since mysql doens't handle null strings well.
	if Phone == "" && email == "" {
		query =
			`Insert into Account (Pseudo, HashPassword, salt, Role) values 
		(?, ?, ?, 1)`
		return Db.Exec(query, name, hash, salt)
	}
	if Phone == "" {
		query =
			`Insert into Account (Pseudo, Email, HashPassword, salt, Role) values 
		(?, ?, ?, ?, 1)`
		return Db.Exec(query, name, email, hash, salt)
	}
	if email == "" {
		query =
			`Insert into Account (Pseudo, HashPassword, Phonenumber, salt, Role) values 
		(?, ?, ?, ?, 1)`
		return Db.Exec(query, name, hash, Phone, salt)
	}

	query =
		`Insert into Account (Pseudo, Email, HashPassword, Phonenumber, salt, Role) values 
	(?, ?, ?, ?, ?, 1)`
	return Db.Exec(query, name, email, hash, Phone, salt)
}

func Newstudysheet(hash string, name string, id_sub int, id_acc int) (int, error) {
	result, err := Db.Exec("INSERT INTO StudySheet (Hash, Name, Id_SubCategory, Id_Account) VALUES (?, ?, ?, ?)", hash, name, id_sub, id_acc)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func InsertCategory(name string) (int, error) {
	result, err := Db.Exec("INSERT INTO Category (Name) VALUES (?)", name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func InsertSubCategory(name string, catId int) (int, error) {
	result, err := Db.Exec("INSERT INTO SubCategory (Name, Id_Category) VALUES (?, ?)", name, catId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}
