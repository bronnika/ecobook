package models

import (
	"time"
)

type NewsResponse struct {
	ID           int       `json:"id" gorm:"column:id"`
	Title        string    `json:"title" gorm:"column:title"`
	Text         string    `json:"text" gorm:"column:text"`
	Photo        string    `json:"photo" gorm:"column:photo"`
	EventDate    time.Time `json:"event_date" gorm:"column:event_date"`
	NewsTypeCode string    `json:"news_type_code" gorm:"news_type_code"`
	CreateDate   time.Time `json:"create_date" gorm:"column:create_date"`
}
