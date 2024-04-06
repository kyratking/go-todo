package routes

import (
	todo_controllers "example/todo/controllers/todo"
	"example/todo/middlewares"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todos", todo_controllers.Show).Methods("GET")
	r.Use(middlewares.ContentTypeJSONMiddleware)
	return r
}
