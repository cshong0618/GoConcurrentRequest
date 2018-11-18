package main

import (
	"github.com/go-chi/chi"
	"gojsonplaceholder/controller"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/user", controller.UserController)

	r.NotFound(func (writer http.ResponseWriter, request *http.Request){
		writer.WriteHeader(404)
	})

	http.ListenAndServe(":3333", r)
}


