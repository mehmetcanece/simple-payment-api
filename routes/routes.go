package routes

import (
	"simple-payment-api/controllers"

	"github.com/gin-gonic/gin"
)
 
func InitRoutes() *gin.Engine{ // Geriye bir *gin.Engine döndürüyor !! Bu, Gin’in HTTP server’ını temsil eder.
	r := gin.Default() // default logger ve error recovery middleware

	r.GET("/", func(c *gin.Context) { //get endpointi
		c.JSON(200, gin.H{
			"message": "Backend API is running"})
	})

	r.POST("/payments",controllers.PostPayment) //post isteklerini paymentsa yönlendiriyoruz
	r.GET("/payments", controllers.GetPayments)

	return r
}