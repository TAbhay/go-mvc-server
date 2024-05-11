package main

import (
	"fmt"
	Routes "go-mvc-server/routes"
)

func main() {

	fmt.Println("Let's begin Boom Boom !")
	r := Routes.SetupRouter()
	fmt.Printf("Listening to port %s", "8080")
	r.Run(":8080")

}
