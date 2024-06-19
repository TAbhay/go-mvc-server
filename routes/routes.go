package routes

import (
	"go-mvc-server/controllers"
	"go-mvc-server/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggingMiddleware())

	test := r.Group("/user")
	{
		test.GET("validate", controllers.UserValidation)
	}
	fake := r.Group("/fake")
	{
		fake.GET("data", controllers.FakeController)
	}

	master := r.Group("/report")
	{
		master.POST("v1", controllers.ReportController)
	}

	cluster := r.Group("/openshift")
	{
		cluster.POST("/cluster", controllers.CreateCluster())

		cluster.GET("/cluster", controllers.GetClusters())

		cluster.GET("/cluster/:id", controllers.GetCluster())

		cluster.PUT("/cluster/:id", controllers.UpdateCluster())

		cluster.GET("/cluster/filter", controllers.FilterClusters())

		cluster.DELETE("/cluster/:id", controllers.DeleteCluster())
	}

	return r
}
