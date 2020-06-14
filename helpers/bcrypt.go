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