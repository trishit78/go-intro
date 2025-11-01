package utils 

import (
	"time"
	"github.com/dgrijalva/jwt-go"

)

var secretKey = []byte("secretpassword")

func GenerateToken (userId int )  (string,error){
	claims:= jwt.MapClaims{}
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims,error){
	token,err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error){
		if _,ok := token.Method.(*jwt.SigningMethodHS256); !ok{
			return nil, fmt.Errorf("Invalid signing method")
		}
		return secretKey,nil
	})

	if err != nil {
		return nil,err
	}
	    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
}
 
  return nil, fmt.Errorf("Invalid token")
}