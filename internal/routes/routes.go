package routes

import (
	"api/internal/controllers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)
//Routes
var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/user/{userInfo}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{userName}", controllers.GetUserByUsername).Methods("GET")
	router.HandleFunc("/usermeasure/", controllers.GetUsersHasMeasure).Methods("GET")

	router.HandleFunc("/type/", controllers.CreateValueTypes).Methods("POST")
	router.HandleFunc("/type/{username}", controllers.GetTypesByUser).Methods("GET")
	router.HandleFunc("/alltypes/{description}", controllers.GetTypeForAllUsersByDescription).Methods("GET")

	router.HandleFunc("/measure/", controllers.L3DecoderPayload).Methods("POST")
	router.HandleFunc("/measure/{username}", controllers.GetAllMeasureFromUsername).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
