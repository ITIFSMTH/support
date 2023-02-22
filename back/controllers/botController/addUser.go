package botController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type AddUserRequestBody struct {
	TGID     int64  `json:"tg_id" binding:"required"`
	BotAPI   string `json:"bot_api" binding:"required"`
	Username string `json:"username" binding:"required"`
	TGLink   string `json:"tg_link"`
}

func (*BotController) AddUser(c *gin.Context) {
	// Get User Data
	var requestBody AddUserRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Get bot ID with API key
	var bot models.Bot
	db.Model(&models.Bot{}).Where("tg_api = ?", requestBody.BotAPI).Scan(&bot)

	// Create New User
	db.Model(&models.User{}).Create(&models.User{
		TGID:     requestBody.TGID,
		BotID:    int(bot.ID),
		Username: requestBody.Username,
		TGLink: requestBody.TGLink,
	})

	// Create New Chat
	db.Model(&models.Chat{}).Create(&models.Chat{
		UserID: requestBody.TGID,
		Users:  pq.Int64Array{int64(requestBody.TGID)},
	})

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
