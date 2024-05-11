package routes

import (
	"api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/properties/", controllers.GetProperties).Methods("GET")
	router.HandleFunc("/properties/{PropertyId}", controllers.GetPropertyById).Methods("GET")
	router.HandleFunc("/properties/", controllers.CreateProperty).Methods("POST")
	router.HandleFunc("/properties/{PropertyId}", controllers.DeleteProperty).Methods("DELETE")
	router.HandleFunc("/properties/{PropertyId}", controllers.UpdateProperty).Methods("PUT")

}
