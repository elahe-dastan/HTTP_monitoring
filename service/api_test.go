//nolint: testpackage
package service

import (
	"HTTP_monitoring/config"
	"HTTP_monitoring/db"
	"HTTP_monitoring/store"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Register(t *testing.T) {
	cfg := config.Read()
	d := db.New(cfg.Database)

	api := API{
		User:   store.NewUser(d),
		URL:    store.NewURL(d),
		Config: cfg.JWT,
	}

	e := echo.New()
	e.POST("/register", api.Register)

	registerationJSON := `{"Email":"parham.alvani@gmail.com",
							"Password":"12345"}`

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(registerationJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err, "Cannot read body")

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	fmt.Println(string(body))
}

func Login(t *testing.T) string {
	cfg := config.Read()
	d := db.New(cfg.Database)

	api := API{
		User:   store.NewUser(d),
		URL:    store.NewURL(d),
		Config: cfg.JWT,
	}

	e := echo.New()
	e.POST("/login", api.Login)

	loginJSON := `{"Email":"parham.alvani@gmail.com",
							"Password":"12345"}`

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	//nolint: bodyclose
	resp := rec.Result()
	defer checkClose(resp)
	body, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err, "Cannot read body")

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	fmt.Println(string(body))

	return string(body)
}

func Add(t *testing.T, token string) {
	cfg := config.Read()
	d := db.New(cfg.Database)

	api := API{
		User:   store.NewUser(d),
		URL:    store.NewURL(d),
		Config: cfg.JWT,
	}

	e := echo.New()
	e.POST("/url", api.Add)

	addJSON := `{"URL": "https://www.google.com"}`

	req := httptest.NewRequest(http.MethodPost, "/url", strings.NewReader(addJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Token", token)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	//nolint: bodyclose
	resp := rec.Result()
	defer checkClose(resp)
	body, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err, "Cannot read body")

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	fmt.Println(string(body))
}

func TestAPI(t *testing.T) {
	Register(t)
	token := Login(t)
	Add(t, token)
}

func checkClose(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
