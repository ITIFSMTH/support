package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

func (*AdminController) GetBots(c *gin.Context) {
	// Get DB
	db := db.GetDb()

	// Get bots
	var bots []models.Bot
	db.Model(&models.Bot{}).Find(&bots)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{
		Data: gin.H{"bots": bots},
	})
}
