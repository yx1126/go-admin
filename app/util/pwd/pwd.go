package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func Encode(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Matches(password, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
	return err == nil
}
