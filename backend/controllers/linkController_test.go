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

	"github.com/codescalersinternships/Linktree-RawanMostafa/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func signupAndLogin(t *testing.T) (token string, username string) {
	t.Helper()
	r := gin.Default()
	r.POST("/user/register", Signup)
	randomStr := GenerateRandomString()
	body := creds{
		Username:   "test_user" + randomStr,
		Password:   "test_password",
		Bio:        "test bio",
		FirstName:  "test",
		SecondName: "user",
	}
	marshalled, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	body = creds{
		Username: "test_user" + randomStr,
		Password: "test_password",
	}

	r.POST("/user/login", Login)
	marshalled, _ = json.Marshal(body)
	req, _ = http.NewRequest("POST", "/user/login", bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")

	res = httptest.NewRecorder()
	r.ServeHTTP(res, req)

	bodyData, _ := io.ReadAll(res.Body)

	if res.Code == http.StatusOK {
		part := strings.Split(string(bodyData), ":")[1]
		token = strings.Trim(part, `"} \n`)
	}
	return token, body.Username
}

type AddLinkRes struct {
	Message string `json:"message"`
	LinkID  string `json:"linkid"`
}

var addLinkResBody AddLinkRes

func TestAddLink(t *testing.T) {
	token, _ := signupAndLogin(t)
	r := gin.Default()
	r.POST("/api/v1/link/", AddLink)
	body := struct {
		Url      string `json:"url"`
		Platform string `json:"platform"`
	}{
		Url:      "test_url_" + GenerateRandomString(),
		Platform: "test_platform",
	}
	marshalled, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/v1/link/", bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
	bodyData, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(bodyData, &addLinkResBody)
	if err != nil {
		t.Error("unmarshal error")
	}

}

func TestEditLink(t *testing.T) {
	token, _ := signupAndLogin(t)
	r := gin.Default()
	r.PUT("/api/v1/link/:link_id", EditLink)
	body := models.LinkRequest {
		Url: "test_url_new" + GenerateRandomString(),
		Platform: "test_new_platform",
	}
	marshalled, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", "/api/v1/link/"+addLinkResBody.LinkID, bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteLink(t *testing.T) {
	token, _ := signupAndLogin(t)
	r := gin.Default()

	r.DELETE("/api/v1/link/:link_id", DeleteLink)
	req, _ := http.NewRequest("DELETE", "/api/v1/link/"+addLinkResBody.LinkID, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetUserLinks(t *testing.T) {
	token, username := signupAndLogin(t)
	r := gin.Default()

	r.POST("/api/v1/link/", AddLink)
	body := struct {
		Url      string `json:"url"`
		Platform string `json:"platform"`
	}{
		Url:      "test_url_" + GenerateRandomString(),
		Platform: "test_platform",
	}
	marshalled, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/v1/link/", bytes.NewReader(marshalled))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
	bodyData, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(bodyData, &addLinkResBody)
	if err != nil {
		t.Error("unmarshal error")
	}

	r.GET("/api/v1/link/:username", GetUserLinks)
	req, _ = http.NewRequest("GET", "/api/v1/link/"+username, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res = httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	bodyData, _ = io.ReadAll(res.Body)
	var links []models.Link
	err = json.Unmarshal(bodyData, &links)
	if err != nil {
		t.Error("unmarshal error")
	}

	assert.Equal(t, len(links), 1)
}
