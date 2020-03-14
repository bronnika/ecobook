package models

type NewsResponse struct {
	ID           int    `json:"id" gorm:"column:p_id"`
	Title        string `json:"title" gorm:"column:p_title"`
	Text         string `json:"text" gorm:"column:p_text"`
	Photo        string `json:"photo" gorm:"column:p_photo"`
	EventDate    string `json:"event_date" gorm:"column:p_event_date"`
	NewsTypeCode string `json:"news_type_code" gorm:"p_news_type_code"`
	Increment    int    `json:"increment" gorm:"column:p_increment"`
	Participate  bool   `json:"participate" gorm:"column:participate"`
	CreateDate   string `json:"create_date" gorm:"column:p_create_date"`
}
