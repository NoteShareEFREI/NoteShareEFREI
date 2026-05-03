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
			`Insert into Account (Pseudo, HashPassword, Phone, salt, Role) values 
		(?, ?, ?, ?, 1)`
		return Db.Exec(query, name, hash, Phone, salt)
	}

	query =
		`Insert into Account (Pseudo, Email, HashPassword, Phonenumber, salt, Role) values 
	(?, ?, ?, ?, ?, 1)`
	return Db.Exec(query, name, email, hash, Phone, salt)
}
