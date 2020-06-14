package routers

import (
	"github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	r := mux.NewRouter()

	HandleAuth(r)
	ProductHandler(r)

	return r
}