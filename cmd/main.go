package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/thalesmonteiro/measurementApi/pkg/routes"
	"api/pkg/routes"
	"log"
	"net/http"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}