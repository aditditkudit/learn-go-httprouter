package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id")
		fmt.Fprintf(w, text)
	})

	request := httptest.NewRequest("GET", "https://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
