package main

import (
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default() //default burada otomatik logger ve recovery middleware ile gelen bir router objesi

	// "/" endpoint'ine GET isteği gelirse bu çalışacak
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Backend API is running.",
		})
	})

	r.Run(":8000")  // 8080 portundan server'ı başlat
}