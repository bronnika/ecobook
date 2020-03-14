package service

import (
	"ecobook/models"
	"ecobook/repository"
)

func GetProductsByCategoryID(categoryID int64) []models.Product {
	return repository.GetProductsByCategoryID(categoryID)
}
