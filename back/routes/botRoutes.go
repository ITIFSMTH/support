package routes

import (
	"support-back/controllers/adminController"
	"support-back/controllers/botController"
	"support-back/middlewares"
	"support-back/shared"

	"github.com/gin-gonic/gin"
)

func RegisterBotRoutes(r *gin.RouterGroup) {
	// Create new group "/bot"
	botGroup := r.Group("/bot")

	// Use middleware Auth (Only Role Bot)
	botGroup.Use(middlewares.AuthHandler(shared.RoleBot))

	// Get controller object
	controller := new(botController.BotController)

	// Bind routes
	botGroup.GET("/bots", new(adminController.AdminController).GetBots)
	botGroup.GET("/greeting", new(adminController.AdminController).GetGreeting)
	botGroup.GET("/getuser", controller.GetUser)
	botGroup.POST("/adduser", controller.AddUser)
	botGroup.POST("/addmessage", controller.AddMessage)
}
