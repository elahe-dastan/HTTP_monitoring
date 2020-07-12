package authentication

import (
	"HTTP_monitoring/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user model.User) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfks") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(token string) (in bool, i int) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("jdnfksdmfks"), nil
	})

	if err != nil {
		return false, 0
	}

	auth := claims["authorized"].(bool)
	exp := claims["exp"].(float64)
	id := claims["user_id"].(float64)

	if auth && exp > float64(time.Now().Unix()){
		return true, int(id)
	}

	return false, int(id)
}
