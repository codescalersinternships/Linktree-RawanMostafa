package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func signupAndLogin(t *testing.T) (token string) {
	t.Helper()
	r := gin.Default()
	r.POST("/public/register", Signup)
	randomStr := GenerateRandomString()
	body := creds{
		Username:   "test_user" + randomStr,
		Password:   "test_password",
		Bio:        "test bio",
		FirstName:  "test",
		SecondName: "user",
	}
	marshalled, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/public/register", bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	body = creds{
		Username: "test_user" + randomStr,
		Password: "test_password",
	}

	r.POST("/public/login", Login)
	marshalled, _ = json.Marshal(body)
	req, _ = http.NewRequest("POST", "/public/login", bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")

	res = httptest.NewRecorder()
	r.ServeHTTP(res, req)

	bodyData, _ := io.ReadAll(res.Body)

	if res.Code == http.StatusOK {
		part := strings.Split(string(bodyData), ":")[1]
		token = strings.Trim(part, `"} \n`)
	}
	return
}
func TestAddLink(t *testing.T) {
	token := signupAndLogin(t)
	r := gin.Default()
	r.POST("/links/add", AddLink)
	body := struct {
		Url      string `json:"url"`
		Platform string `json:"platform"`
	}{
		Url:      "test_url_" + GenerateRandomString(),
		Platform: "test_platform",
	}
	marshalled, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/links/add", bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}
