package adminController

import (
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"
	"support-back/shared"
	"time"

	"github.com/gin-gonic/gin"
)

type AddOperatorRequestBody struct {
	Login    string `json:"login" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=80"`
}

func (*AdminController) AddOperator(c *gin.Context) {
	// Get User Data
	var requestBody AddOperatorRequestBody
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

	// Create operator
	operator := &models.Worker{
		Login:    requestBody.Login,
		Password: hash,
		Role:     shared.RoleOperator,
	}

	// Add new operator
	if err := db.Model(&models.Worker{}).Create(operator).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorOperatorLoginExist,
		})
		return
	}

	// Add new operator stat
	db.Model(&models.WorkerStat{}).Create(&models.WorkerStat{
		Worker:   operator.ID,
		LastDate: time.Now(),
	})

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{
		Data: gin.H{"operator": operator},
	})
}
