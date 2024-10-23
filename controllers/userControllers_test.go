package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type creds struct {
	Username string
	Password string
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateRandomString() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

var random string = GenerateRandomString()

func TestSignup(t *testing.T) {
	testcases := []struct {
		testcaseName       string
		body               creds
		expectedStatusCode int
		expectedBody       string
	}{
		{
			testcaseName: "test correct signup",
			body: creds{
				Username: "test_user" + fmt.Sprint(random),
				Password: "test_password",
			},
			expectedStatusCode: http.StatusCreated,
			expectedBody:       `{"message":"User registered successfully"}`,
		},
		{
			testcaseName: "test incorrect signup : username already exists",
			body: creds{
				Username: "test_user" + fmt.Sprint(random),
				Password: "test_password",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"error":"Username already exists"}`,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.testcaseName, func(t *testing.T) {
			r := gin.Default()
			r.POST("/public/register", Signup)
			marshalled, err := json.Marshal(testcase.body)
			if err != nil {
				log.Fatalf("failed to marshall credentials: %s", err)
			}
			req, err := http.NewRequest("POST", "/public/register", bytes.NewReader(marshalled))
			if err != nil {
				log.Fatalf("impossible to build request: %s", err)
			}
			req.Header.Add("content-type", "application/json")

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			assert.Equal(t, testcase.expectedStatusCode, res.Code)

			bodyData, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error reading response body %v", err)
			}

			assert.Equal(t, testcase.expectedBody, string(bodyData))
		})
	}
}

func TestLogin(t *testing.T) {
	testcases := []struct {
		testcaseName       string
		body               creds
		expectedStatusCode int
		expectedBody       string
	}{
		{
			testcaseName: "test correct login",
			body: creds{
				Username: "test_user" + fmt.Sprint(random),
				Password: "test_password",
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"token":`,
		},
		{
			testcaseName: "test incorrect login : username doesn't exist",
			body: creds{
				Username: "test_userz",
				Password: "testpassword",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       `{"error":"username doesn't exist"}`,
		},
		{
			testcaseName: "test incorrect login : invalid password",
			body: creds{
				Username: "test_user" + fmt.Sprint(random),
				Password: "test_password2",
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody:       `{"error":"invalid password"}`,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.testcaseName, func(t *testing.T) {
			r := gin.Default()
			r.POST("/public/login", Login)
			marshalled, err := json.Marshal(testcase.body)
			if err != nil {
				log.Fatalf("failed to marshall credentials: %s", err)
			}
			req, err := http.NewRequest("POST", "/public/login", bytes.NewReader(marshalled))
			if err != nil {
				log.Fatalf("impossible to build request: %s", err)
			}
			req.Header.Add("content-type", "application/json")

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			assert.Equal(t, testcase.expectedStatusCode, res.Code)
			bodyData, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error reading response body %v", err)
			}

			if testcase.expectedStatusCode == http.StatusOK {
				assert.Contains(t, string(bodyData), testcase.expectedBody)
			} else {
				assert.Equal(t, testcase.expectedBody, string(bodyData))
			}

		})
	}
}
