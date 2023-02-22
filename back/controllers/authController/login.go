package authController

import (
	"net/http"
	"support-back/db"
	"support-back/middlewares"
	"support-back/models"
	"support-back/responses"
	"support-back/shared"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (*AuthController) Login(c *gin.Context) {
	// Get User Data
	var requestBody LoginRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Auth
	data := models.Worker{}
	db.Model(&models.Worker{}).Where("login = ?", requestBody.Login).Select("password, id, role").Scan(&data)
	expectedPassword := data.Password
	role := data.Role
	id := int(data.ID)

	// Check password
	if r := comparePasswordHash(requestBody.Password, expectedPassword); !r {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorWrongData,
		})
		return
	}

	// Issue JWT token
	token, _ := middlewares.GenerateToken([]byte(middlewares.SigningKey),
		requestBody.Login,
		role,
		id,
	)

	// Add to statistic
	if role == shared.RoleOperator {
		models.WorkerStat{}.NewLogin(db, uint(id))
	}

	// Send answer
	c.JSON(http.StatusOK, &responses.Response{
		Data: gin.H{"token": token},
	})
}
