package service

import (
	"HTTP_monitoring/model"
	"HTTP_monitoring/request"
	"HTTP_monitoring/store"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var ErrLoggedOut = errors.New("you are not logged in")

type API struct {
	User   store.SQLUser
	URl    store.SQLURL
}

func (a API) Run() {
	e := echo.New()

	//Users register
	e.POST("/register", a.Register)
	e.POST("/login", a.Login)
	e.POST("/url", a.Add)
	e.Logger.Fatal(e.Start(":8080"))
}

func (a API) Register(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == ""{
		return echo.NewHTTPError(http.StatusBadRequest, "Email and password cannot be empty")
	}

	if err := a.User.Insert(user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (a API) Login(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == ""{
		return echo.NewHTTPError(http.StatusBadRequest, "Email and password cannot be empty")
	}

	us, err := a.User.Retrieve(user)
	if err != nil {
		return err
	}

	token, err := CreateToken(us)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}

func (a API) Add(c echo.Context) error {
	var newUrl request.URL

	if err := c.Bind(&newUrl); err != nil {
		return err
	}

	in, id := ValidateToken(newUrl.Token)

	if !in {
		return c.JSON(http.StatusForbidden, ErrLoggedOut)
	}

	var url model.URL

	url.UserId = id
	url.Url = newUrl.Url

	if err := a.URl.Insert(url); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, url)
}

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