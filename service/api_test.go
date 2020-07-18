//nolint: testpackage
package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/db"
	"github.com/elahe-dastan/HTTP_monitoring/mock"
	"github.com/elahe-dastan/HTTP_monitoring/store"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEmailEmpty(t *testing.T) {
	cfg := config.Read()

	api := API{
		User:   mock.User{Info: map[string]string{}},
		URL:    mock.URL{Urls: map[string]int{}},
		Config: cfg.JWT,
	}

	e := echo.New()
	e.POST("/register", api.Register)

	registerationJSON := `{"Password":"1378"}`

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(registerationJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err, "Cannot read body")

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	fmt.Println(string(body))
}

func Register(t *testing.T, api API) {
	e := echo.New()
	e.POST("/register", api.Register)

	registerationJSON := `{"Email":"parham.alvani@gmail.com",
							"Password":"1378"}`

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

func Login(t *testing.T, api API) string {
	e := echo.New()
	e.POST("/login", api.Login)

	loginJSON := `{"Email":"parham.alvani@gmail.com",
							"Password":"1378"}`

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

func Add(t *testing.T, token string, api API) {
	e := echo.New()
	e.POST("/url", api.Add)

	addJSON := `{"URL": "https://www.google.com", "Period": 2}`

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
	cfg := config.Read()
	d := db.New(cfg.Database)

	api := API{
		User:   store.NewUser(d),
		URL:    store.NewURL(d),
		Config: cfg.JWT,
	}

	Register(t, api)
	token := Login(t, api)
	Add(t, token, api)
}

func checkClose(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
