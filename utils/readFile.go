package utils

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
	"../models"
)

var (
	AppSettings models.Settings
)

func ReadSettings(){
	fmt.Println("starting reading settings file")
	configFile, err := os.Open("./settings.json")

	defer configFile.Close()

	if err != nil {
		log.Fatal("Couldn't open config file", err.Error())
	}

	fmt.Println("starting decoding settings file")
	err = json.NewDecoder(configFile).Decode(&AppSettings)

	if err != nil {
		log.Fatal("Couldn't decode settings json file", err.Error())
	}

	log.Println(AppSettings)
	return
}