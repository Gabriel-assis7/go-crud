package main

import (
	"log"
	"net/http"

	"github.com/gabriel_assis7/simple-go-mod/configs"
	"github.com/gabriel_assis7/simple-go-mod/models"
	"github.com/gabriel_assis7/simple-go-mod/routes"
	"github.com/gorilla/mux"
)

func main() {
	dbConnection := configs.ConnectDB()

	_, err := dbConnection.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	defer dbConnection.Close()

	router := mux.NewRouter()
	routes.RegisterRouters(router)

	log.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
