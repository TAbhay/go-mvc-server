package controllers

import (
	"fmt"
	"go-mvc-server/utils"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ReportController(c *gin.Context) {
	var data []string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Received data:", data)
	files := make([]string, 0)
	endpoints := []string{"user/validate", "user2/validate"}
	for _, endpoint := range endpoints {
		fmt.Println(endpoint)
		resp, err := http.Get("http://localhost:8080/" + endpoint)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		filename := "response_" + endpoint + ".json"

		err = utils.SaveToFile(filename, body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		files = append(files, filepath.Join("reports", filename))
	}
	c.Header("Content-Disposition", "attachment; filename=responses.zip")
	c.Header("Content-Type", "application/zip")
	c.Writer.WriteHeader(http.StatusOK)
	for _, file := range files {
		c.FileAttachment(file, filepath.Base(file))
	}
}
