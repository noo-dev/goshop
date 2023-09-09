package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Printf("Failed to generate password: %v", err)
		return ""
	}

	return string(hashed)
}
