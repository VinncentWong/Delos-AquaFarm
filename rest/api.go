package rest

import (
	"net/http"

	farmHandler "github.com/VinncentWong/Delos-AquaFarm/app/farm/handler"
	pondHandler "github.com/VinncentWong/Delos-AquaFarm/app/pond/handler"
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

func (r *Routing) InitializeFarmRouting(handler *farmHandler.FarmHandler) {
	farmGroup := r.router.Group("/farm")
	farmGroup.POST("/create", handler.CreateFarm)
	farmGroup.PUT("/update/:farmId", handler.UpdateFarm)
	farmGroup.DELETE("/delete/:farmId", handler.DeleteFarm)
	farmGroup.GET("/gets", handler.GetAll)
}

func (r *Routing) InitializePondRouting(handler *pondHandler.PondHandler) {
	pondGroup := r.router.Group("/pond")
	pondGroup.POST("/create/:farmId", handler.CreatePond)
	pondGroup.PUT("/update/:pondId", handler.UpdatePond)
	pondGroup.DELETE("/delete/:pondId", handler.DeletePond)
	pondGroup.GET("/gets", handler.GetAll)
}
