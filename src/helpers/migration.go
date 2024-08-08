package helpers

import (
	"log"
	"transaction-be/src/configs"
	"transaction-be/src/models"
)

func Migration() {
	err := configs.DB.AutoMigrate(&models.Product{}, &models.Customer{}, &models.Transaction{}, &models.TransactionDetail{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
