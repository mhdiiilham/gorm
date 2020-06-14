package routers

import (
	"github.com/gorilla/mux"
	c "github.com/mhdiiilham/gorm/controllers"
)

// ProductHandler ...
func ProductHandler(r *mux.Router) {
	r.HandleFunc("/products", c.GetProducts).Methods("GET")
	r.HandleFunc("/products", c.PostProduct).Methods("POST")
}