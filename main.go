package main

import (
	"curd_gin/database"
	"curd_gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.Connect()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	fmt.Println("Successfully Connected")
	err = models.Migrate()
	if err != nil {
		fmt.Println("Error during migration:", err)
		return
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
