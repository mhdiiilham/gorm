package routers

import (
	"github.com/gorilla/mux"
	c "github.com/mhdiiilham/gorm/controllers"
)

// HandleAuth ...
func HandleAuth(r *mux.Router) {
	r.HandleFunc("/auth/signup", c.SignUp).Methods("POST")
}