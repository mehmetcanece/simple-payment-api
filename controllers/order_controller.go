package controllers

import (
	"fmt"
	"net/http"

	"simple-payment-api/database"
	"simple-payment-api/models"

	"github.com/gin-gonic/gin"
)


func PostOrder(c *gin.Context){
	var req models.Order


	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	order := models.Order{
		TableID: req.TableID,
		TotalAmount: req.TotalAmount,
	}
	database.DB.Create(&order)

	fmt.Printf("Received order: %d with total amount %.2f\n", req.TableID, req.TotalAmount)

	c.JSON(http.StatusOK, gin.H{
		"status": "Order received",
		"table_id": req.TableID,
		"total_amount": req.TotalAmount,
	})

}

func GetOrders(c *gin.Context)  {
	var orders []models.Order

	result:= database.DB.Find(&orders)
	if result.Error !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Orders received",
		"orders": orders,
	})

}
