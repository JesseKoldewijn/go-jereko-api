package routes

import (
	"net/http"

	utils "github.com/JesseKoldewijn/go-jereko-api/routes/api/utils"
	v1 "github.com/JesseKoldewijn/go-jereko-api/routes/api/v1"
	"github.com/gin-gonic/gin"
)

// The function `RouteGroupRoot` sets up route groups with different prefixes in a Gin router.
func RouteGroupRoot(router *gin.Engine) {
	// prefix /assets
	routeGroupStatic(router)
	// prefix /
	routeGroupCore(router)
	// prefix /api
	routeGroupApi(router)
}

// Route group for static file routes
func routeGroupStatic(router *gin.Engine) {
	staticGroup := router.Group("/assets")
	staticGroup.StaticFS("/", http.Dir("./public"))
}

// Route group for general routes
func routeGroupCore(router *gin.Engine) {
	coreGroup := router.Group("/")
	// path: /
	coreGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Jereko API!",
		})
	})
}

// Route group for API routes
func routeGroupApi(router *gin.Engine) {
	apiGroup := router.Group("/api")
	// path: /api/v1
	v1.RouteGroupV1(apiGroup)
	// path: /api/utils
	utils.RouteGroupUtils(apiGroup)
}
