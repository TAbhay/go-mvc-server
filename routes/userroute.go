package routes

import (
	"go-mvc-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	test := r.Group("/user")
	{
		test.GET("validate", controllers.UserValidation)
	}
	return r
}
