package controllers

import (
	"encoding/json"
	"fmt"
	"go-mvc-server/api"
	"go-mvc-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReportController(c *gin.Context) {

	endpoints := []string{"fake/data", "fake/data"}

	combinedData, err := api.FetchTestData(endpoints)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch test data"})
		return
	}

	err = generateAndSaveReport(combinedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate or save report"})
		return
	}

	c.JSON(http.StatusOK, combinedData)
}

func generateAndSaveReport(combinedData map[string]interface{}) error {

	combinedFilePath := "../reports/combined_responses.json"
	combinedDataBytes, err := json.Marshal(combinedData)
	if err != nil {
		fmt.Println("Failed to marshal combined data:", err)
		return err
	}

	err = utils.SaveToFile(combinedFilePath, combinedDataBytes)
	if err != nil {
		fmt.Println("Failed to save combined data to file:", err)
		return err
	}

	utils.GenerateReport()
	return nil
}