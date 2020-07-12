package service

import (
	"HTTP_monitoring/authentication"
	"HTTP_monitoring/config"
	"HTTP_monitoring/model"
	"HTTP_monitoring/request"
	"HTTP_monitoring/store"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrLoggedOut = errors.New("you are not logged in")

type API struct {
	User store.SQLUser
	URL  store.SQLURL
	Config config.JWT
}

func (a API) Run() {
	e := echo.New()

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
		return c.JSON(http.StatusNotFound, err.Error())
	}

	token, err := authentication.CreateToken(us.ID, a.Config)
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

	in, id := authentication.ValidateToken(newUrl.Token)

	if !in {
		return c.JSON(http.StatusForbidden, ErrLoggedOut)
	}

	var url model.URL

	url.UserId = id
	url.Url = newUrl.Url

	if err := a.URL.Insert(url); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, url)
}