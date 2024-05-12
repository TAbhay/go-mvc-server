package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userValidation(c *gin.Context) {

	res, err := api.callUserAPI(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call user API"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"response": res,
	})
}
