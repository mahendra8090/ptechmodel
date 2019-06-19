package server

import (
	"net/http"
	"newgo/test/routes"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	routes.Addroutes(r)
	http.ListenAndServe(":8080", r)

}
