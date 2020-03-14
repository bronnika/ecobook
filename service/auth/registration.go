package auth

import (
	"ecobook/models"
	"ecobook/repository/auth"
	"ecobook/utils"
	"errors"
)

func Registration(request models.Registration, userId *int) error {
	if request.Token != utils.AppSettings.Token {
		return errors.New("wrong token")
	}

	if err := auth.Registration(request, userId); err != nil {
		return err
	}

	return nil
}
