package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FakeController(c *gin.Context) {
	responseData := map[string]interface{}{
		"name": "faker",
		"tests": []map[string]interface{}{
			{
				"type":        "functional",
				"name":        "Login Test",
				"description": "Tests the login functionality",
				"passed":      true,
			},
			{
				"type":        "performance",
				"name":        "Load Test",
				"description": "Tests the load handling of the app",
				"passed":      false,
			},
		},
	}
	c.JSON(http.StatusOK, responseData)
}
