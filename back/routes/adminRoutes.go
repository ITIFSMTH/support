package routes

import (
	"support-back/controllers/adminController"
	"support-back/middlewares"
	"support-back/shared"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.RouterGroup) {
	// Create new group "/admin"
	adminGroup := r.Group("/admin")

	// Use middleware Auth (Only Role Admin)
	adminGroup.Use(middlewares.AuthHandler(shared.RoleAdmin))

	// Get controller object
	controller := new(adminController.AdminController)

	// Bind routes
	adminGroup.GET("/operators", controller.GetOperators)
	adminGroup.GET("/bots", controller.GetBots)
	adminGroup.GET("/greeting", controller.GetGreeting)
	adminGroup.POST("/edit/greeting", controller.EditGreeting)
	adminGroup.POST("/mailing", controller.Mailing)
	adminGroup.POST("/operator/add", controller.AddOperator)
	adminGroup.POST("/operator/remove", controller.RemoveOperator)
	adminGroup.POST("/operator/edit/login", controller.EditOperatorLogin)
	adminGroup.POST("/operator/edit/password", controller.EditOperatorPassword)
	adminGroup.POST("/operator/statistic", controller.GetOperatorStat)
	adminGroup.POST("/bot/add", controller.AddBot)
	adminGroup.POST("/bot/remove", controller.RemoveBot)
}
