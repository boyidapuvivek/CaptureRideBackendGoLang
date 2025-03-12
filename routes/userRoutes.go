package routes

import (
	"CaptureRideBackendGoLang/handles"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/users", handles.GetUsers)
	router.POST("/users", handles.AddUsers)
}
