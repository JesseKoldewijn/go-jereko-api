package main

import (
	"net/http"
	"os"
	"time"

	v1 "github.com/JesseKoldewijn/go-jereko-api/routes/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	env_vars := os.Environ()

	var PORT = "8080"

	// get PORT env var
	for _, env_var := range env_vars {
		if env_var[:4] == "PORT" {
			PORT = env_var[5:]
		}
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Jereko API!",
		})
	})

	v1.RouteGroupV1(router)

	s := &http.Server{
		Addr:           ":" + string(PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	url := "http://localhost:" + string(PORT)

	println("Server running on: " + url)

	s.ListenAndServe()
}
