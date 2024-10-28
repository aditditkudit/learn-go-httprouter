package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		Text := "Products " + params.ByName("id")
		fmt.Fprint(writer, Text)
	})

	server := http.Server{
		Handler: router,
		Addr:    ":9090",
	}

	server.ListenAndServe()
}
