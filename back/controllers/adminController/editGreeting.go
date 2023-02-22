package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

type EditGreetingRequestBody struct {
	Greeting string `json:"greeting" binding:"required,min=1,max=4096"`
}

func (*AdminController) EditGreeting(c *gin.Context) {
	// Get User Data
	var requestBody EditGreetingRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Edit greeting
	db.Model(&models.Setting{}).Update("greeting", requestBody.Greeting)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
