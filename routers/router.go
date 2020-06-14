package routers
 import (
	"github.com/gorilla/mux"
	"github.com/mhdiiilham/gorm/controllers"
 )

 // Router ...
 func Router() *mux.Router {
	 router := mux.NewRouter()

	 router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	 router.HandleFunc("/products", controllers.PostProduct).Methods("POST")

	 return router
 }