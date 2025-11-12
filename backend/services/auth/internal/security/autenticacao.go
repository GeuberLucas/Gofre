package security

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (string,error){
	perms := jwt.MapClaims{}
	perms["authorized"]=true
	perms["exp"]= time.Now().Add(time.Hour*1).Unix()
	perms["userId"]=userId
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,perms)

	return token.SignedString([]byte("HEJ7hRH9eWkJgkvghExLAAJvcDfZL9VlNYtAXuuCu9qQ3HCF7W1d7HGvrHvrnKbXcaM5+SPohHdHoUToeuAg6Q=="))


}