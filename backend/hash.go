package backend

import (
	"golang.org/x/crypto/bcrypt"
)

// Called when a new user account is created.
func NewHash(password string, name string) (string, error) {
	//TO BE IMPLEMENTED
	panic("Not yet Implemented")
}

// Called when the user logs in
func VerifyPerson(password string, name string) (bool, int) {
	panic("Implement a database to attach to first.")
	/*
		res := database.Doquery("select pwd,salt from Account where Pseudo=%s")
		for i := 0; i < res.len; i++ {
		for true {
			pwd := res[i][0a]
			salt := res[i][1]
			hash, err := HashPassword(password, salt)
			if err != nil {
				continue
			}
			if VerifyPassword(hash, pwd) {
				//Need to create a wrapper for the request
				//id := mysql.getid(pwd)
				id := database.query("select Accid from Account where pwd=%s", pwd)
				return true, id
			}
		}
		return false, -1
	*/
}

// HashPassword generates a bcrypt hash for the given password and salt.
func HashPassword(password string, salt int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
