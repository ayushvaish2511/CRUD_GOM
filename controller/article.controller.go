package controller

import (
	"curd_gin/services"

	"github.com/gin-gonic/gin"
)

func ArticleController(r *gin.Engine) {
	r.POST("/create-post", services.CreatePost)
	r.GET("/get-posts", services.GetPost)
}
