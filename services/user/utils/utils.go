package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword is a func to generate password
func GeneratePassword(password string) (string, error) {
	pwdByte := []byte(password)
	bcryptPassword, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword[:]), nil
}

// CheckPassword check the password
// return a nil on success
func CheckPassword(plainPassword string, password string) error {
	hashedPassword := []byte(password)
	pwd := []byte(plainPassword)
	return bcrypt.CompareHashAndPassword(hashedPassword, pwd)
}
