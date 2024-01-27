package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

func EncryptPassword(data *string) (*string, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(*data), bcryptCost)
	if err != nil {
		return nil, err
	}
	ret := string(encpw)
	return &ret, nil
}
func ComparePassword(encryptedPassword *string, password *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*encryptedPassword), []byte(*password))
	return err == nil
}
