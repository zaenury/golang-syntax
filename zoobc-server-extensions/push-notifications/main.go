package main

import (
	"github.com/zoobc-server-extensions/push-notifications/Routers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	Routers.Routers(router)
}
