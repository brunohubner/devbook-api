package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/brunohubner/devbook-api/src/security"
	"github.com/brunohubner/devbook-api/src/utils"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(steep string) error {
	if err := user.format(steep); err != nil {
		return err
	}

	if err := user.validate(steep); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(steep string) error {
	if user.Name == "" {
		return errors.New("The field 'name' is required")
	}

	if user.Nick == "" {
		return errors.New("The field 'nick' is required")
	}

	if user.Email == "" {
		return errors.New("The field 'email' is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New(fmt.Sprintf("Email: %s", err.Error()))
	}

	if steep == "signup" && user.Password == "" {
		return errors.New("The field 'password' is required")
	}

	return nil
}

func (user *User) format(steep string) error {
	user.Name = utils.StandardizeSpaces(strings.TrimSpace(user.Name))
	user.Nick = utils.StandardizeSpaces(strings.TrimSpace(user.Nick))
	user.Email = utils.StandardizeSpaces(strings.TrimSpace(user.Email))

	if steep == "signup" {
		hashedPassword, err := security.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
