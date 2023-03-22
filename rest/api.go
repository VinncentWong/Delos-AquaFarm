package rest

import (
	"net/http"

	"github.com/VinncentWong/Delos-AquaFarm/app/farm/handler"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	router *gin.Engine
}

func NewRouting(router *gin.Engine) *Routing {
	return &Routing{
		router: router,
	}
}

func (r *Routing) InitializeCheckHealthRouting() {
	r.router.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

func (r *Routing) InitializeFarmRouting(handler *handler.FarmHandler) {
	farmGroup := r.router.Group("/farm")
	farmGroup.POST("/create", handler.CreateFarm)
}
