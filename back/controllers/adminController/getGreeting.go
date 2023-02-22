package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

func (*AdminController) GetGreeting(c *gin.Context) {
	// Get DB
	db := db.GetDb()

	// Get greeting
	var setting models.Setting
	db.Model(&models.Setting{}).Find(&setting)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{
		Data: gin.H{"greeting": setting.Greeting},
	})
}
