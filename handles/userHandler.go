package handles

import (
	"CaptureRideBackendGoLang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users)
}

func AddUsers(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.Users = append(models.Users, newUser)
	c.JSON(http.StatusOK, newUser)
}
