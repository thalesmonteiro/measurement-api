package main

import (
	_ "api/cmd/docs"
	"api/internal/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

// @title Measurement API
// @version 1.0
// @description This is a server to easily access API features, including database and Sensor measurement.
// @BasePath
func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
