package helpers

import (
	"golang.org/x/crypto/bcrypt"
	log "github.com/sirupsen/logrus"
)

// HashPassword ...
func HashPassword(p []byte) string {
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.MinCost)
	if err != nil {log.Fatal(err)}

	stringPwd := string(hash)
	return stringPwd
}

// Compare ...
func Compare(hashed string, plain[]byte) bool {
	byteHashed := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHashed, plain)
	if err != nil {
		log.Info(err)
		return false
	}

	return true
}