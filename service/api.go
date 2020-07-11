package service

import (
	"HTTP_monitoring/model"
	"HTTP_monitoring/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

	if err := a.User.Retrieve(user); err != nil {
		return err
	}



	//create token

	return c.JSON(http.StatusCreated, user)
}

func (a API) Add(c echo.Context) error {
	var url model.URL

	if err := c.Bind(&url); err != nil {
		return err
	}

	if err := a.URl.Insert(url); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, url)
}