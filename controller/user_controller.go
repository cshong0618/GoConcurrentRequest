package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	. "gojsonplaceholder/service"
	"net/http"
	"strconv"
)

func UserController(r chi.Router) {
	r.Get("/{userId}", func (writer http.ResponseWriter, request *http.Request) {
		userId := chi.URLParam(request, "userId")

		id, err := strconv.Atoi(userId)

		if err != nil {
			writer.WriteHeader(500)
		} else {
			userData := GetUserData(id)
			writer.Header().Add("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(&userData)
		}
	})
}
