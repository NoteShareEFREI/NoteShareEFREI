package backend

import (
	"NoteShareEFREI/database"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

// Called when a new user account is created.
func NewHash(password string) (string, int, error) {
	salt_rounds := rand.Intn(15) + 10 //Random number between 11 and 25
	//The salt rounds represents the number of time the string will be hashed.
	//It being random makes it harder to calculate the original password but we need to be carefull for it to not to be to high as this is ressource intensive.
	hash, err := HashPassword(password, salt_rounds)
	if err != nil {
		return "", 0, err
	}

	return hash, salt_rounds, nil
}

// Called when the user logs in
func VerifyPerson(password string, name string) (bool, int) {
	//Query for the infos.
	//Trying to match the given name with Email and Pseudo fields.
	query := `
	select HashPassword,salt from Account where Pseudo=? OR Email=?;`
	rows, err := database.Db.Query(query, name, name)
	if err != nil {
		fmt.Println(err.Error())
		return false, -3 //The sql query did not work.
	}
	for rows.Next() { //Cycling through everyone with the same username
		var (
			pwd  string
			salt int
		)
		//for every if statement, we ignore errors and continue if there aren't errors.
		//(We need to check for everyone and obviously there will be time it doens't work.)
		//We only need the one output when it's valid.
		//It is done this way so that multiple people can have the same name and that we do not give any information to the client through errors.
		err := rows.Scan(&pwd, &salt)
		if err == nil { // If there are no errors, continue.
			if VerifyPassword(password, pwd) {
				id, err := database.Getidfrompseudoandhash(name, pwd)
				if err == nil {
					return true, id
				}
			}
		}
	}
	//Could not find any match.
	return false, -1
}

// HashPassword generates a bcrypt hash for the given password and salt.
func HashPassword(password string, salt int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	fmt.Print(hash)
	fmt.Print(password)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
