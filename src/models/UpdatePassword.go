package models

import (
	"errors"
)

type UpdatePassword struct {
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
	CurrentPassword string `json:"currentPassword"`
}

func (password *UpdatePassword) Validate() error {
	if len(password.NewPassword) < 8 {
		return errors.New("New password must be longer than 8 characters")
	}

	if password.NewPassword != password.ConfirmPassword {
		return errors.New("Password and Password confirmation dont match")
	}

	return nil
}
