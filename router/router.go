package router

import (
	"github.com/gorilla/mux"
	"github.com/krutzika/backend_golang/handlers"
)

func SetupRouter() *mux.Router {
	request := mux.NewRouter()
	request.HandleFunc("/news", handlers.GetNews).Methods("GET")
	return request
}
