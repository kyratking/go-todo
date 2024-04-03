package routes

import (
	todo_controllers "example/todo/controllers/todo"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todos", todo_controllers.Show).Methods("GET")
	return r
}
