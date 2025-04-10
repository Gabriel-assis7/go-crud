package routes

import (
	"github.com/gabriel_assis7/simple-go-mod/controllers"
	"github.com/gorilla/mux"
)

func RegisterRouters(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
}
