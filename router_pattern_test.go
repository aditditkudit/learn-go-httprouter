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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		Id := params.ByName("id")
		ItemId := params.ByName("itemId")
		Text := "Products " + Id + " Item " + ItemId
		fmt.Fprint(writer, Text)
	})

	request := httptest.NewRequest("GET", "https://localhost:3000/products/1/items/23", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Products 1 Item 23", string(body))
}

func TestRouterPatternCatchParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		images := params.ByName("image")
		Text := "Image : " + images
		fmt.Fprint(writer, Text)
	})

	request := httptest.NewRequest("GET", "https://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}
