package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello Get")
	})

	server := http.Server{
		Handler: router,
		Addr:    ":9090",
	}

	server.ListenAndServe()
}
