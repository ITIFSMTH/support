package adminController

import (
	"net/http"
	"support-back/bot"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"os/exec"

	"github.com/gin-gonic/gin"
)

type AddBotRequestBody struct {
	TGAPI string `json:"tg_api"`
}

func (*AdminController) AddBot(c *gin.Context) {
	// Get User Data
	var requestBody AddBotRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Check Token
	name, err := bot.CheckToken(requestBody.TGAPI)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorTokenInvalid,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Create new bot
	botDb := models.Bot{
		TGAPI:  requestBody.TGAPI,
		TGName: name,
	}

	if err := db.Model(&models.Bot{}).Create(&botDb).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBotExist,
		})
		return
	}

	// Init new bot
	//bot.InitBot(botDb)

	// Restart TG Bots
	exec.Command("/root/.nvm/versions/node/v18.2.0/bin/pm2", "restart", "index").Output()

	// Send Answer
	c.JSON(
		http.StatusOK,
		&responses.Response{
			Data: gin.H{"bot": botDb},
		},
	)
}
