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
				"api":         "Operators",
				"description": "Are required operators a present",
				"response":    "a is not present",
				"result":      "Failed",
				"type":        "Test",
				"value":       0,
			},
			{
				"api":         "Operators",
				"description": "Are required operators b present",
				"response":    "b is present",
				"result":      "Passed",
				"type":        "Test",
				"value":       1,
			},
		},
	}
	c.JSON(http.StatusOK, responseData)
}
