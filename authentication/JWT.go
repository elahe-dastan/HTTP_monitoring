package authentication

import (
	"HTTP_monitoring/config"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id int, cfg config.JWT) (string, error) {
	var err error
	//Creating Access Token
	err = os.Setenv("ACCESS_SECRET", cfg.SECRET) //this should be in an env file
	if err != nil {
		return "", err
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(cfg.Expiration)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

//nolint: gofumpt
func ValidateToken(token string, cfg config.JWT) (in bool, i int) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRET), nil
	})

	if err != nil {
		return false, 0
	}

	auth := claims["authorized"].(bool)
	exp := claims["exp"].(float64)
	id := claims["user_id"].(float64)

	if auth && exp > float64(time.Now().Unix()) {
		return true, int(id)
	}

	return false, int(id)
}
