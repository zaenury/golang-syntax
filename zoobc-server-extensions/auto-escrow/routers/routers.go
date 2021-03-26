package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routers(routes *mux.Router) {
	routes.HandleFunc("/Block", CreateItem).Methods("POST")
	// routes.HandleFunc("/Item", UpdateItem).Methods("PUT")

	http.ListenAndServe(":9090", routes)
}
