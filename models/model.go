package model

import (
	"net/http"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// Product model
type Product struct {
	gorm.Model
	Name string `gorm:"type:varchar(50)"`
	Price int `gorm:"type:int(30)"`
}

// Result ...
type Result struct {
	Status int `json:status`
	Message string `json:message`
	Data []Product
}

// RespondJSON ...
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}