package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userValidation(c *gin.Context) {

	res, err := API.callUserAPI(c)
	fmt.Println(res)
	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call user API"})
		return
	}
	println(res)
	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"response": res,
	})
}
