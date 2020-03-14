package service

import (
	"ecobook/models"
	"ecobook/repository"
)

func GetProductsByCategoryID(categoryID int64) []models.Product {
	return repository.GetProductsByCategoryID(categoryID)
}

func AddNewProduct(newProduct models.AddNewProductReq) error {
	return repository.AddNewProduct(newProduct)
}

func GetSaleTypes() []models.SalesTypes {
	return repository.GetSalesTypes()
}
