package controllers

import (
	"fmt"
	"net/http"

	"simple-payment-api/models"

	"github.com/gin-gonic/gin"
)

func PostPayment(c *gin.Context) {
	var req models.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil { // jsondaki veriyle req structını eşleştirmeye çalışıyoruz eğer yanlış struct gelirse hata dönüyoruz doğruysa decam ediyoruz.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Printf("Received payment: %.2f via %s\n", req.Amount, req.Method)

	//2.f iki basamak s string anlamında

	c.JSON(http.StatusOK, gin.H{ //gin.H json döndürüyor genelde
		"status": "Payment received",
		"amount": req.Amount,
		"method": req.Method,
	})
}