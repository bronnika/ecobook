package auth

import (
	"ecobook/db"
	"errors"
	"log"
	"time"
)

func OtpServer(phoneNum string, otpCode string) error {
	log.Println("phone number:", phoneNum, " otp code:", otpCode)

	sqlQuery := `insert into otp(phone,otp_code, exp_date) values (?,?,?)`

	now := time.Now()

	expDate := now.Add(20 * time.Minute)

	if err := db.GetDBConn().Exec(sqlQuery, phoneNum, otpCode, expDate).Error; err != nil {
		log.Println("OtpServer func query error:", err.Error())
		return errors.New("something went wrong")
	}

	return nil
}
