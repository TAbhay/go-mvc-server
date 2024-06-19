package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func FetchTestData(endpoints []string) (map[string]interface{}, error) {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	combinedData := map[string]interface{}{
		"tests": []map[string]interface{}{},
	}

	// Channel to collect results
	results := make(chan map[string]interface{}, len(endpoints))

	// Function to fetch data from an endpoint
	fetchData := func(endpoint string) {
		defer wg.Done()
		time.Sleep(5 * time.Second) // Simulating delay
		resp, err := http.Get("http://localhost:8080/" + endpoint)
		if err != nil {
			fmt.Println("Failed to fetch endpoint:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return
		}

		var testGroup map[string]interface{}
		if err := json.Unmarshal(body, &testGroup); err != nil {
			fmt.Println("Failed to parse JSON response:", err)
			return
		}

		// Send the result to the channel
		results <- testGroup
	}

	// Start goroutines for each endpoint
	for _, endpoint := range endpoints {
		wg.Add(1)
		go fetchData(endpoint)
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results from the channel
	for result := range results {
		mu.Lock()
		combinedData["tests"] = append(combinedData["tests"].([]map[string]interface{}), result)
		mu.Unlock()
	}

	return combinedData, nil
}
