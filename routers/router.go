package routers

import (
	"github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	r := mux.NewRouter()

	ProductHandler(r)

	return r
}