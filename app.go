package main

import (
	"foodapi/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	// Init gin router
	router := gin.Default()

	// Its great to version your API's
	v1 := router.Group("/api/v1")
	{
		// Define the hello controller
		hello := new(controllers.HelloWorldController)
		// Define a GET request to call the Default
		// method in controllers/hello.go
		v1.GET("/hello", hello.Default)
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Init our server
	router.Run(":"+os.Getenv("APP_PORT"))
}

