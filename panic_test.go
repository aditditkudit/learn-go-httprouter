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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(writer, "Panic : ", err)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "https://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))
	//fmt.Println(string(body))
}
