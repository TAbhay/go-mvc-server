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
	test2 := r.Group("/user2")
	{
		test2.GET("validate", controllers.UserValidation)
	}
	master := r.Group("/report")
	{
		master.POST("v1", controllers.ReportController)
	}
	return r
}
