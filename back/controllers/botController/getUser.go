package botController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

type GetUserRequestBody struct {
	TGID int `json:"tg_id" binding:"required"`
}

func (*BotController) GetUser(c *gin.Context) {
	// Get User Data
	var requestBody GetUserRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Get User
	var user models.User
	db.Model(&models.User{}).Where("tg_id = ?", requestBody.TGID).Find(&user)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{
		Data: gin.H{"user": user},
	})
}
