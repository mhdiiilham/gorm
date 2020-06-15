package routers

import (
	"github.com/gorilla/mux"
	c "github.com/mhdiiilham/gorm/controllers"
	a "github.com/mhdiiilham/gorm/middlewares"
)

// ProductHandler ...
func ProductHandler(r *mux.Router) {
	r.Use(a.IsAuthenticated)
	r.HandleFunc("/products", c.GetProducts).Methods("GET")
	r.HandleFunc("/products", c.PostProduct).Methods("POST")
}