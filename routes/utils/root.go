package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
)

func RouteGroupUtils(router *gin.Engine) {
	utils := router.Group("/utils")
	{
		utils.GET("/ip", func(c *gin.Context) {
			ip := c.ClientIP()

			uaString := c.Request.Header.Get("User-Agent")
			uaParsed := useragent.Parse(uaString)

			c.JSON(200, gin.H{
				"ip":       ip,
				"os":       uaParsed.OS,
				"browser":  uaParsed.Name,
				"isMobile": uaParsed.Mobile,
			})
		})

	}
}
