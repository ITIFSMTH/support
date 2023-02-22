package operatorController

import (
	"support-back/ws"

	"github.com/gin-gonic/gin"
)

func (*OperatorController) Upgrade(c *gin.Context) {
	hub := ws.GetHub()
	id, _ := c.Get("id")
	ws.ServeWs(hub, c.Writer, c.Request, int(id.(float64)))
}
