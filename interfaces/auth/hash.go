package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Hashed is encrypting my password
func Hashed(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// CheckHash is checking my password & hashed data
func CheckHash(hash []byte, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
