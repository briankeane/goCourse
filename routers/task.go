package routers

import (
	"github.com/gorilla/mux"

	"github.com/briankeane/goCourse/controllers"
)

func SetTaskRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/tasks", controllers.Tasks.Create).Methods("POST")
	router.HandleFunc("/tasks", controllers.Tasks.ShowAll).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Show).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Update).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Delete).Methods("DELETE")
	return router
}
