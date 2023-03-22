package handler

import (
	"net/http"

	"github.com/VinncentWong/Delos-AquaFarm/app/pond/usecase"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/util"
	"github.com/gin-gonic/gin"
)

type PondHandler struct {
	usecase usecase.IPondUsecase
}

func NewPondHandler(usecase usecase.IPondUsecase) *PondHandler {
	return &PondHandler{
		usecase: usecase,
	}
}

func (h *PondHandler) CreatePond(c *gin.Context) {
	farmId := c.Param("farmId")
	var container domain.Pond
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	result, err := h.usecase.CreatePond(farmId, &container)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create pond", true, result)
}
