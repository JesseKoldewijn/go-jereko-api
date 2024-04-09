package v1

import "github.com/gin-gonic/gin"

func RouteGroupV1 (router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the v1 root!",
			})
		})
	}
}
