package controllers

import (
	"go-mvc-server/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserValidation(c *gin.Context) {

	res, err := api.CallUserAPI(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call user API"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"response": res,
	})
}
