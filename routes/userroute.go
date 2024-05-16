package routes

import (
	"fmt"
	"go-mvc-server/controllers"
	"go-mvc-server/middlewares"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func saveToFile(filename string, data []byte) error {
	// Open file with create flag and write permissions
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func SetupRouter() *gin.Engine {

	r := gin.New() // NEW

	r.Use(gin.Recovery()) // NEW
	r.Use(middlewares.LoggingMiddleware())

	test := r.Group("/user")
	{
		test.GET("validate", controllers.UserValidation)
	}
	test2 := r.Group("/user2")
	{
		test2.GET("validate", controllers.UserValidation)
	}
	master := r.Group("/test")
	{
		master.GET("/run", func(c *gin.Context) {
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

				err = saveToFile("response_"+endpoint+".json", body)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		})
	}
	return r
}
