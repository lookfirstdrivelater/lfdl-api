package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestPing(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "pong",
	}
	// Grab our router
	router := setupRouter()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/ping")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["message"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)
}

func TestLogin(t *testing.T) {
	// Build our expected body
	// body := gin.H{
	// 	"code":   200,
	// 	"expire": "2019-03-26T20:55:11Z",
	// 	"token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTM2MzM3MTEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU1MzYzMDExMX0.6n1vrhIJq_Xqgf67NuGwWU3Hlx5zysBdE6lZrngqYBA",
	// }
	// Grab our router
	router := setupRouter()
	// Perform a POST request with that handler.
	w := performRequest(router, "GET", "/login?username=admin&password=admin")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	// var response map[string]string
	// err := json.Unmarshal([]byte(w.Body.String()), &response)
	// // Grab the value & whether or not it exists
	// value, exists := response["code"]
	// // Make some assertions on the correctness of the response.
	// assert.Nil(t, err)
	// assert.True(t, exists)
	// assert.Equal(t, body["code"], value)
}

func TestBadRequest(t *testing.T) {
	// Build our expected body
	// Grab our router
	router := setupRouter()
	// Perform a POST request with that handler.
	w := performRequest(router, "GET", "/fg")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	// Convert the JSON response to a map
	// var response map[string]string
	// err := json.Unmarshal([]byte(w.Body.String()), &response)
	// // Grab the value & whether or not it exists
	// value, exists := response["code"]
	// // Make some assertions on the correctness of the response.
	// assert.Nil(t, err)
	// assert.True(t, exists)
	// assert.Equal(t, body["code"], value)
}

func Login() string {
	router := setupRouter()
	w := performRequest(router, "GET", "/login?username=admin&password=admin")
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	if err != nil {
		fmt.Println("Failed to get token")
	}
	// Grab the value & whether or not it exists
	value, _ := response["token"]
	return value
}
