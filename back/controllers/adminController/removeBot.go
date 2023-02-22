package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

        "os/exec"
	"github.com/gin-gonic/gin"
)

type RemoveBotRequestBody struct {
	ID int `json:"id" binding:"required,min=0"`
}

func (*AdminController) RemoveBot(c *gin.Context) {
	// Get User Data
	var requestBody RemoveBotRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Get bot
	var bot models.Bot
	db.Model(&bot).Where("id = ?", requestBody.ID).Scan(&bot)

	// Delete webhook
	// http.Get("https://api.telegram.org/bot" + bot.TGAPI + "/setWebhook?remove")

	// Restart TG Bots
        exec.Command("/root/.nvm/versions/node/v18.2.0/bin/pm2", "restart", "index").Output()


	// Remove bot
	db.Unscoped().Delete(&models.Bot{}, requestBody.ID)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
