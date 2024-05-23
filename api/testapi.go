package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchTestData(endpoints []string) (map[string]interface{}, error) {
	combinedData := map[string]interface{}{
		"tests": []map[string]interface{}{},
	}

	for _, endpoint := range endpoints {
		resp, err := http.Get("http://localhost:8080/" + endpoint)
		if err != nil {
			fmt.Println("Failed to fetch endpoint:", err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			continue
		}

		var testGroup map[string]interface{}
		if err := json.Unmarshal(body, &testGroup); err != nil {
			fmt.Println("Failed to parse JSON response:", err)
			continue
		}

		combinedData["tests"] = append(combinedData["tests"].([]map[string]interface{}), testGroup)
	}

	return combinedData, nil
}
