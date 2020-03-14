package auth

import (
	"ecobook/repository/auth"
	"math/rand"
	"strconv"
)

func OtpServer(phoneNum string) error {
	otpCode := rand.Int() % 10000

	otpCodeStr := strconv.Itoa(otpCode)

	// отправка смс OTP юзеру
	if err := auth.OtpServer(phoneNum, otpCodeStr); err != nil {
		return err
	}

	return nil
}
