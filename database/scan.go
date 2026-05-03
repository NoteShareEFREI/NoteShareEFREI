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
