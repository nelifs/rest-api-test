package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strconv"
	"time"
)

var key = "hxfijhsdlkufgdsiuofadsuylfgdsfglusgcf;oiudsvljadygdfhlsejbgidasuyhfldskbvgiyaddf"

func Generate(userId int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        strconv.Itoa(userId),
	})

	tokenString, _ := token.SignedString(key)

	return tokenString
}

func Validate(tokenString string) (bool, error) {
	t, e := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if e != nil {
		return false, fiber.NewError(403, "Invalid token")
	}

	_, ok := t.Claims.(jwt.MapClaims)

	return ok, nil
}

func GetData(tokenString string) jwt.MapClaims {
	t := parse(tokenString)

	claims, _ := t.Claims.(jwt.MapClaims)

	return claims
}

func parse(tokenString string) *jwt.Token {
	t, e := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if e != nil {
		log.Fatal(e)
	}

	return t
}
