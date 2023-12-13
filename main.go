package main

import (
	"curd_gin/database"
	"curd_gin/models"
	"curd_gin/router"
	"fmt"
	// "net/http"
	// "github.com/gin-gonic/gin"
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
	r := router.SetupRouter()
	r.Run()
	
}
