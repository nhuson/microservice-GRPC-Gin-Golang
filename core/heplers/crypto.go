package heplers

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GenerateToken(object *jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, object)
	string, err := token.SignedString([]byte(Getenv("SECRET_KEY")))

	if err != nil {
		log.Println(err)
		return ""
	}

	return string
}
