package services

import (
	"curd_gin/database"
	"curd_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

func CreatePost(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article)

	// Validation
	// validate := validator.New()
	// err := validate.Struct(article)
	// validationErrors := err.(validator.ValidationErrors)

	// if validationErrors != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
	// 	return
	// }
	// Saving in the DB
	result := database.DB.Create(&article)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"article": article})
}

func GetPost(c *gin.Context) {
	var article []models.Article
	result := database.DB.Find(&article)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
}
