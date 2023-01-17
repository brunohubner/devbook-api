package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/brunohubner/devbook-api/src/database"
	"github.com/brunohubner/devbook-api/src/models"
	"github.com/brunohubner/devbook-api/src/repositories"
	"github.com/brunohubner/devbook-api/src/responses"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	user.ID, err = userRepository.Create(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func FindMany(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	users, err := userRepository.FindMany(nameOrNick)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("findOne"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("remove"))
}
