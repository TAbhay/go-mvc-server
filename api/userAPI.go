package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallUserAPI(c *gin.Context) (map[string]interface{}, error) {

	client := openshiftHTTPClient()
	req, err := http.NewRequest(http.MethodGet, "https://reqres.in/api/users?page=2", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("foo", "bar1")

	response, err := client.Do(req)

	if err != nil {

		return nil, err
	}

	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
