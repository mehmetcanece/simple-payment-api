package controllers

import (
	"fmt"
	"net/http"

	"simple-payment-api/database"
	"simple-payment-api/models"

	"github.com/gin-gonic/gin"
)

func PostPayment(c *gin.Context) { //jsondan veriyi alıyoruz, model objesine dönüştürüyoruz payment diye. en son da dbye kaydediyoruz
	var req models.Payment

	if err := c.ShouldBindJSON(&req); err != nil { // jsondaki veriyle req structını eşleştirmeye çalışıyoruz eğer yanlış struct gelirse hata dönüyoruz doğruysa decam ediyoruz.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	payment:= models.Payment{
		Amount: req.Amount,
		Method: req.Method,
		Status: "pending",
	}
	database.DB.Create(&payment)

	fmt.Printf("Received payment: %.2f via %s\n", req.Amount, req.Method)

	//2.f iki basamak s string anlamında

	c.JSON(http.StatusOK, gin.H{ //gin.H json döndürüyor genelde
		"status": "Payment received",
		"amount": req.Amount,
		"method": req.Method,
	})
}
 
func GetPayments(c *gin.Context)  { //dbden tüm ödemeleri çekiyoruz ve json olarak return ediyoruz
		var payments []models.Payment

		// tüm ödeme kayıtlarını çekjcez dbden

		result:= database.DB.Find(&payments)
		if result.Error !=nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		
		
		c.JSON(http.StatusOK, gin.H{
			"status":  "Payment received",
			"payment": payments,
		})

	}