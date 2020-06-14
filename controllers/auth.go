package controllers

import (
	// "encoding/json"
	"encoding/json"
	"net/http"

	"github.com/mhdiiilham/gorm/db"
	"github.com/mhdiiilham/gorm/models"
	h "github.com/mhdiiilham/gorm/helpers"
	log "github.com/sirupsen/logrus"
)

// SignUp to handle user signup
func SignUp(w http.ResponseWriter, r *http.Request) {
	var u model.UserInput

	w.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{
		Fullname: u.Fullname,
		Email: u.Email,
		PasswordHash: h.HashPassword([]byte(u.Password)),
	}
	saveUser := db.Connection().Save(&user)

	if saveUser.Error != nil {
		model.RespondError(w, http.StatusInternalServerError, saveUser.Error.Error())
		return
	}

	defer db.Connection().Close()
	msg := "User created!"
	model.RespondJSON(w, http.StatusOK, msg)
	
}