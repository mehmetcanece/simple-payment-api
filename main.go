package main

import (
	"simple-payment-api/database"
	"simple-payment-api/models"
	"simple-payment-api/routes"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Payment{}) //tablo oluşturmak için



	r := routes.InitRoutes()
	r.Run(":8000")
}