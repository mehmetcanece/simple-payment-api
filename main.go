package main

import (
	"simple-payment-api/database"
	"simple-payment-api/models"
	"simple-payment-api/routes"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Payment{}) //sunucu kapanıp açılsa da dbye kaydedilen her şey orada kalır. 
	database.DB.AutoMigrate(&models.Order{})



	r := routes.InitRoutes()
	r.Run(":8000")
}