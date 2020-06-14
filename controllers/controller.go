package controllers

import (
	"encoding/json"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/mhdiiilham/gorm/db"
	"github.com/mhdiiilham/gorm/models"
)

// GetProducts return list of products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	db.Connection().Find(&products)
	model.RespondJSON(w, http.StatusOK, products)
}

// PostProduct ...
func PostProduct(w http.ResponseWriter, r *http.Request) {
	product := model.Product{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		model.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := db.Connection().Save(&product).Error; err != nil {
		model.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	model.RespondJSON(w, http.StatusCreated, product)
}