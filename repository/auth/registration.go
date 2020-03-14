package auth

import (
	"ecobook/db"
	"ecobook/models"
	"errors"
	"log"
)

func Registration(request models.Registration, userId *int) error {
	var id []int

	sqlQuery := `insert into "user"(name, phone) values(?,?) returning id`

	if err := db.GetDBConn().Raw(sqlQuery, request.Name, request.PhoneNum).Pluck("id", &id).Error; err != nil {
		log.Println("Registrantion func query error:", err.Error())
		return errors.New("something went wrong")
	}

	*userId = id[0]

	return nil
}
