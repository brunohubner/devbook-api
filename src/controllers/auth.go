package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/brunohubner/devbook-api/src/auth"
	"github.com/brunohubner/devbook-api/src/database"
	"github.com/brunohubner/devbook-api/src/models"
	"github.com/brunohubner/devbook-api/src/repositories"
	"github.com/brunohubner/devbook-api/src/responses"
	"github.com/brunohubner/devbook-api/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	savedUser, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ComparePasswords(savedUser.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateJWT(savedUser.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
}
