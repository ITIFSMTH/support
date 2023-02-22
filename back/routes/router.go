package routes

import (
	"support-back/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	InitMiddleware(engine)

	ApiGroup := engine.Group("/api")

	RegisterBotRoutes(ApiGroup)
	RegisterAdminRoutes(ApiGroup)
	RegisterOperatorRoutes(ApiGroup)
	RegisterAuthRoutes(ApiGroup)
}

func InitMiddleware(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}
