package main

import (
	"simple-payment-api/routes"
)

func main() {
	r := routes.InitRoutes()
	r.Run(":8000")
}