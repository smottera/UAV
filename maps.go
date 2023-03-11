package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// corsMiddleware sets up CORS headers for requests to the server
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type")
		c.Header("Access-Control-Expose-Headers", "Authorization")
		c.Header("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	// Get the Google Maps API key from an environment variable
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing GOOGLE_MAPS_API_KEY environment variable")
	}

	// Create a new Gin router
	router := gin.Default()

	// Set up CORS middleware
	router.Use(corsMiddleware())

	// Set up a route to act as a proxy to the Google Maps API
	router.GET("/maps/api/*path", func(c *gin.Context) {
		path := c.Param("path")

		// Make a request to the Google Maps API
		resp, err := http.Get("https://maps.googleapis.com/" + path + "&key=" + apiKey)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()

		// Copy the response headers and body to the Gin context
		for key, values := range resp.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}
		c.Status(resp.StatusCode)
		_, err = io.Copy(c.Writer, resp.Body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	})

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
