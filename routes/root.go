package routes

import (
	"github.com/JesseKoldewijn/go-jereko-api/routes/utils"
	v1 "github.com/JesseKoldewijn/go-jereko-api/routes/v1"
	"github.com/gin-gonic/gin"
)

func RouteGroupRoot(router *gin.Engine) {
	// path: /v1
	v1.RouteGroupV1(router)
	// path: /utils
	utils.RouteGroupUtils(router)
}
