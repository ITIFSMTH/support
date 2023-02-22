package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

type EditOperatorLoginRequestBody struct {
	ID    int    `json:"id" binding:"required,min=0"`
	Login string `json:"login" binding:"required,min=3,max=20"`
}

func (*AdminController) EditOperatorLogin(c *gin.Context) {
	// Get User Data
	var requestBody EditOperatorLoginRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Edit operator login
	db.Model(&models.Worker{}).Where("id = ?", requestBody.ID).Update("login", requestBody.Login)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
