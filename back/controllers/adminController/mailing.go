package adminController

import (
	"net/http"
	"support-back/bot"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

type MailingRequestBody struct {
	Mail string `json:"mail" binding:"required,min=1,max=4096"`
}

func (*AdminController) Mailing(c *gin.Context) {
	// Get User Data
	var requestBody MailingRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	go bot.SendMailing(requestBody.Mail)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
