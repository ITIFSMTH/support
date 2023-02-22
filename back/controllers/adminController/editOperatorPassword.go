package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

type EditOperatorPasswordRequestBody struct {
	ID       int    `json:"id" binding:"required,min=0"`
	Password string `json:"password" binding:"required,min=8,max=80"`
}

func (*AdminController) EditOperatorPassword(c *gin.Context) {
	// Get User Data
	var requestBody EditOperatorPasswordRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get Password Hash
	hash, err := getPasswordHash(requestBody.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &responses.Response{
			Error: responses.ErrorServer,
		})
		return
	}
	// Get DB
	db := db.GetDb()

	// Edit operator login
	db.Model(&models.Worker{}).Where("id = ?", requestBody.ID).Update("password", hash)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
