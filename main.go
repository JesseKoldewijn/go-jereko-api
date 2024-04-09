package main

import (
	"net/http"
	"time"

	v1 "github.com/JesseKoldewijn/go-jereko-api/routes/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Jereko API!",
		})
	})

	v1.RouteGroupV1(router)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	println("Server running on: http://localhost:8080")

	s.ListenAndServe()
}
