package auth

import "ecobook/repository/auth"

func OtpCheck(phoneNum string, otpCode string, isRegistered *bool, userId *int) error {
	if err := auth.OtpCheck(phoneNum, otpCode, isRegistered, userId); err != nil {
		return err
	}

	return nil
}
