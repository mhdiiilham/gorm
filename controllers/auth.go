package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/mhdiiilham/gorm/db"
	h "github.com/mhdiiilham/gorm/helpers"
	"github.com/mhdiiilham/gorm/models"
	log "github.com/sirupsen/logrus"
)

// SignUp to handle user signup
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.UserInput

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

// SignIn ...
func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	respond := map[string]string{}
	reqBody := map[string]string{}

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		log.Fatal(err)
		model.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	user, err := findEmail(reqBody["email"])

	if err != nil {
		model.RespondError(w, http.StatusBadRequest, "Email / Password is Wrong")
		return
	}

	// model.RespondJSON(w, 200, user.PasswordHash)
	plain := reqBody["password"]
	isValid := h.Compare(user.PasswordHash, []byte(plain))

	if !isValid {
		model.RespondError(w, http.StatusBadRequest, "Email / Password is Wrong")
		return
	}

	token, err := h.CreateJWTToken(string(user.ID), user.Email)
	if err != nil {
		log.Fatal(err)
		model.RespondError(w, http.StatusBadRequest, "Something went wrong")
		return
	}http://localhost:8000/auth/signin
	respond["token"] = token
	respond["fullname"] = user.Fullname
	model.RespondJSON(w, http.StatusAccepted, respond)
	
}

func findEmail(e string) (model.User, error) {
	user := model.User{}

	find := db.Connection().Where("email = ?", e).Find(&user)

	if find.Error != nil {
		return user, find.Error
	}

	return user, nil
}