package jwtToken

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 1).Unix()
	perms["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))

}



func ValidateToken(r *http.Request ) error{
	tokenString := extractToken(r)
	token,err:=jwt.Parse(tokenString,returnVerificationKey)
	if err != nil {
		return err
	}

	if _,ok:= token.Claims.(jwt.MapClaims);ok && token.Valid{
		return nil
	}

	return errors.New("invalid Token")
}


func extractToken( r *http.Request) string{
	headerAuthorization := r.Header.Get("Authorization")
	tokenValue := strings.Split(headerAuthorization," ")
	if len(tokenValue) == 2{
		return tokenValue[1]
	}
	return ""
}


func returnVerificationKey(token *jwt.Token)(interface{},error){
	if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok{
		return "", fmt.Errorf("verifyKeyToken: unexpected signing method ! %v",token.Header["alg"])
	}
	return []byte(os.Getenv("SECRET_KEY")),nil
}