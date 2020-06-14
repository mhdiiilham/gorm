package main

import (
	"github.com/mhdiiilham/gorm/db"
	"github.com/mhdiiilham/gorm/routers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	port := ":8000"
	router := routers.Router()
	db.Connection()
	log.Info("server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}