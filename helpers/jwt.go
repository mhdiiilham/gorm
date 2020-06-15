package helpers

import (
	"fmt"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// CreateJWTToken ...
func CreateJWTToken(id uint, email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authenticated"] = true
	atClaims["user_id"] = id
	atClaims["user_email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte(os.Getenv("SECRET")))
}


func verifyToken(ht string) (*jwt.Token, error) {
	token, err := jwt.Parse(ht, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", ok)
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil{
		return nil, err
	}

	return token, nil
}

// ExtractedJWT ...
func ExtractedJWT(ht string) (string, error) {
	token, err := verifyToken(ht)
	if err != nil {
		log.Info(err)
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessID, ok := claims["user_email"].(string)
		if !ok {
			log.Info(ok)
			return "", err
		 }
		return accessID, nil
	}
	log.Info(err)
	return "", err
}