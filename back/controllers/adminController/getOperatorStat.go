package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

type GetOperatorStatRequestBody struct {
	ID int `json:"id" binding:"required"`
}

func (*AdminController) GetOperatorStat(c *gin.Context) {
	// Get User Data
	var requestBody GetOperatorStatRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Get operator stats
	var operatorStat models.WorkerStat
	db.Model(&operatorStat).Where("worker = ?", requestBody.ID).First(&operatorStat)

	// Check is stat fresh
	operatorStat.Update(db)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{
		Data: gin.H{"statistic": operatorStat},
	})
}
