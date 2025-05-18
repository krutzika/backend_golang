package router

import (
	"github.com/gorilla/mux"
	"github.com/krutzika/backend_golang/handlers"
)

func SetupRouter() *mux.Router {
	request := mux.NewRouter()
	request.HandleFunc("/news", handlers.GetNews).Methods("GET")
	request.HandleFunc("/news/category", handlers.GetNewsCategory).Methods("GET")
	return request
}
