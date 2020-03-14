package auth

import (
	"ecobook/db"
	"ecobook/models"
	"errors"
	"log"
)

func OtpCheck(phoneNum string, otpCode string, isRegistered *bool, userId *int) error {
	var (
		queryResp models.OtpCheckQuery
	)

	sqlQuery := `select * from otp_check(?,?)`

	if err := db.GetDBConn().Raw(sqlQuery, phoneNum, otpCode).Scan(&queryResp).Error; err != nil {
		log.Println("OtpCheck func select query error:", err.Error())
		return errors.New("something went wrong")
	}

	log.Println("query resp:", queryResp)

	if queryResp.ErrCode != 0 {
		return errors.New(queryResp.ErrDesc)
	}

	*isRegistered = queryResp.IsRegistered
	*userId = queryResp.UserId

	return nil
}
