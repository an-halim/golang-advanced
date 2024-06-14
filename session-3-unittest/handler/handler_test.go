package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/an-halim/golang-advanced/session-3-unittest/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHelloMessage(t *testing.T) {
	t.Run("Positive Case - Correct Message", func(t *testing.T) {
		expectedOutput := "Halo dari Gin!"
		actualOutput := handler.GetHelloMessage()
		require.Equal(t, expectedOutput, actualOutput, "The message should be '%s'", expectedOutput)
	})
}

func TestRootHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/", handler.RootHandler)

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code) // Check if the status code is 200

	expectedBody := `{"message":"Halo dari Gin!"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

type JsonRequest struct {
	Message string `json:"message"`
}

func TestPostHandler(t *testing.T) {
	// Setup router
	r := gin.Default()
	r.POST("/", handler.PostHandler)

	t.Run("Positive Case", func(t *testing.T) {
		// Persiapan data JSON
		requestBody := JsonRequest{Message: "Hello from test!"}
		requestBodyBytes, _ := json.Marshal(requestBody)

		// Buat permintaan HTTP POST
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		expectedBody := `{"message":"Hello from test!"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Negative Case - EOF Error", func(t *testing.T) {
		requestBody := ""
		requestBodyBytes := []byte(requestBody)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "{\"error\":\"EOF\"}")
	})
}

func TestGetHandler(t *testing.T) {
	// Setup router
	r := gin.Default()
	r.GET("/", handler.GetHandler)

	t.Run("Positive Case + With query name", func(t *testing.T) {

		name := "Praisindo"

		// Buat permintaan HTTP GET
		req, _ := http.NewRequest("GET", "/?name="+name, nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		expectedBody := `{"message":"Halo dari ` + name + `!"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Positive Case + Without query name", func(t *testing.T) {

		// Buat permintaan HTTP GET
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		expectedBody := `{"message":"Halo dari Gin!"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Negative Case - Route Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/NotFound", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
