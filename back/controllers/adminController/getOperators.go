package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"
	"support-back/shared"

	"github.com/gin-gonic/gin"
)

func (*AdminController) GetOperators(c *gin.Context) {
	// Get DB
	db := db.GetDb()

	// Get operators
	var operators []models.Worker
	db.Model(&models.Worker{}).Where("role = ?", shared.RoleOperator).Find(&operators)

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{
		Data: gin.H{"operators": operators},
	})
}
