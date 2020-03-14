package participate

import (
	"ecobook/db"
	"ecobook/models"
)

func Participate(eventID int, userID int) (int, error) {

	sqlQuery := `INSERT INTO public.participant(news_id, user_id) VALUES (?, ?)`

	if err := db.GetDBConn().Exec(sqlQuery, eventID, userID).Error; err != nil {
		return 0, err
	}
	sqlCount := `SELECT COUNT(ID) FROM public.participant WHERE news_id = ?`
	count := models.Count{}
	if err := db.GetDBConn().Raw(sqlCount, eventID).Scan(&count).Error; err != nil {
		return 0, err
	}

	return count.Count, nil
}
