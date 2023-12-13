package router

import (
	// "net/http"

	"curd_gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	controller.ArticleController(r)
	return r
}