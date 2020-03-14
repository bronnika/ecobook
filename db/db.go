package db

import (
	"fmt"
	"log"

	"ecobook/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

func initDB() *gorm.DB {
	settingParams := utils.AppSettings.PostgresParams

	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		settingParams.Server,
		settingParams.Port,
		settingParams.Database,
		settingParams.User,
		settingParams.Password)

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		log.Fatal("Connection to database failed!", err.Error())
	}

	db.LogMode(true)

	db.SingularTable(true)

	return db
}

func StartDbConnection() {
	database = initDB()
}

func GetDBConn() *gorm.DB {
	return database
}
