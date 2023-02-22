package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"

	"github.com/gin-gonic/gin"
)

type RemoveOperatorRequestBody struct {
	ID int `json:"id" binding:"required,min=0"`
}

func (*AdminController) RemoveOperator(c *gin.Context) {
	// Get User Data
	var requestBody RemoveOperatorRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Remove operator
	db.Unscoped().Delete(&models.Worker{}, requestBody.ID)

	// Remove operator stat
	db.Unscoped().Where("worker = ?", requestBody.ID).Delete(&models.WorkerStat{})

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
