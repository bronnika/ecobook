package auth

import (
	"ecobook/repository/auth"
	"math/rand"
	"strconv"
)

func OtpServer(phoneNum string, otp *string) error {
	otpCode := rand.Intn(9999-1000) + 1000

	otpCodeStr := strconv.Itoa(otpCode)

	*otp = otpCodeStr

	// отправка смс OTP юзеру
	if err := auth.OtpServer(phoneNum, otpCodeStr); err != nil {
		return err
	}

	return nil
}
