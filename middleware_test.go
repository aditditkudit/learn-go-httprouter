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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHttp(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieve request")
	middleware.Handler.ServeHTTP(writer, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	middleware := LogMiddleware{router}
	request := httptest.NewRequest("GET", "https://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHttp(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
	//fmt.Println(string(body))
}
