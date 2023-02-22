package routes

import (
	"support-back/controllers/operatorController"
	"support-back/middlewares"
	"support-back/shared"

	"github.com/gin-gonic/gin"
)

func RegisterOperatorRoutes(r *gin.RouterGroup) {
	// Create new group "/operator"
	operatorGroup := r.Group("/operator")

	// Use middleware Auth (Role ADmin || Role Operator)
	operatorGroup.Use(middlewares.AuthHandler(shared.RoleOperator, shared.RoleAdmin))

	// Get controller object
	controller := new(operatorController.OperatorController)

	// Bind routes
	operatorGroup.GET("/ws", controller.Upgrade)
}
