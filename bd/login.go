package bd

import (
	"backendgo_aws/models"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	user, found, _ := CheckAlreadyExist(email)
	if !found {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
