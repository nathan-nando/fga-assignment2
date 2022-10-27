package main

import (
	_ "assignment2/docs"
	"assignment2/models"
	"assignment2/routers"
)

//@title Assignment 2 API
//@version 1.0
//@description Simple service
//@termsOfService http://swagger.io/terms/
//@contact.name API Support
//@contact.email nathan@gmail.com
//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html
//@host localhost:8080
//@BasePath /

func main() {
	models.StartDB()

	PORT := ":8080"
	gin := routers.StartServer()
	err := gin.Run(PORT)
	if err != nil {
		panic(err)
	}
}
