package routes

import (
	"go-mvc-server/controllers"
	"go-mvc-server/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.New() // NEW

	r.Use(gin.Recovery()) // NEW
	r.Use(middlewares.LoggingMiddleware())

	test := r.Group("/user")
	{
		test.GET("validate", controllers.UserValidation)
	}
	return r
}
