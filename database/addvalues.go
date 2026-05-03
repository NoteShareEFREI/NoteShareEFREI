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

func Newstudysheet(id int, hash string, name string, id_sub int, id_acc int) (sql.Result, error) {
	return Db.Exec("INSERT INTO StudySheet (Id_Sheet, Hash, Name, Id_SubCategory, Id_Account) VALUES (?, ?, ?, ?, ?)", id, hash, name, id_sub, id_acc)
}

func InsertCategory(id int, name string) (sql.Result, error) {
	return Db.Exec("INSERT IGNORE INTO Category (Id_Category, Name) VALUES (?, ?)", id, name)
}

func InsertSubCategory(id int, name string, catId int) (sql.Result, error) {
	return Db.Exec("INSERT IGNORE INTO SubCategory (Id_SubCategory, Name, Id_Category) VALUES (?, ?, ?)", id, name, catId)
}
