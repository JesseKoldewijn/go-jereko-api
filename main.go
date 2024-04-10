package main

import (
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"

	v1 "github.com/JesseKoldewijn/go-jereko-api/routes/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get all environment variables
	env_vars := os.Environ()

	// Default values
	PORT := "8080"

	// Get PORT env var
	for _, env_var := range env_vars {
		if env_var[:4] == "PORT" {
			PORT = env_var[5:]
		}
	}

	// Create router
	router := gin.Default()

	// Default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Jereko API!",
		})
	})

	// Routes
	v1.RouteGroupV1(router)

	// Create server
	s := &http.Server{
		Addr:           ":" + string(PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Server address
	port := ":" + string(PORT)
	url := "http://localhost" + port

	println("Server running on: " + url)

	// Listen and serve
	err := endless.ListenAndServe(port, s.Handler)

	// Check for errors
	if err != nil {
		println("Error starting server: ", err)
	}
}
