package handler

import (
	"net/http"

	"github.com/JesseKoldewijn/go-jereko-api/routes"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Create router
	router := gin.Default()

	// Routes
	routes.RouteGroupRoot(router)

	// running gin engine
	router.ServeHTTP(w, r)
}
