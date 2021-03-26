package Routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers(routes *mux.Router) {
	routes.HandleFunc("/Notification", GetNotifications).Methods("GET")

	fmt.Println("Server Running on port :9090")
	http.ListenAndServe(":9090", routes)
}
