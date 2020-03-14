package models

type OtpServer struct {
	PhoneNum string `json:"phone_num,required"`
}

type OtpCheck struct {
	PhoneNum string `json:"phone_num"`
	OtpCode  string `json:"otp_code"`
}

type OtpCheckQuery struct {
	IsRegistered bool   `gorm:"column:p_is_registered"`
	UserId       int    `gorm:"column:p_user_id"`
	ErrCode      int    `gorm:"column:p_err_code"`
	ErrDesc      string `gorm:"column:p_err_desc"`

	/*Id      int       `gorm:"column:id"`
	OtpCode string    `gorm:"column:otp_code"`
	ExpDate time.Time `gorm:"column:exp_date"`*/
}

type Registration struct {
	PhoneNum string `json:"phone_num"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
