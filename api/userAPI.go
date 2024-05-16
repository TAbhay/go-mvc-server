package api

import (
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CallUserAPI(c *gin.Context) (map[string]interface{}, error) {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 15 * time.Second,
			}).Dial,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			TLSHandshakeTimeout:   15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			ExpectContinueTimeout: 15 * time.Second,
		},
	}
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
