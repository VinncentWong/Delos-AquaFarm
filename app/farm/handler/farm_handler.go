package handler

import (
	"net/http"

	"github.com/VinncentWong/Delos-AquaFarm/app/farm/usecase"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/util"
	"github.com/gin-gonic/gin"
)

type FarmHandler struct {
	usecase usecase.IFarmUsecase
}

func NewFarmHandler(usecase usecase.IFarmUsecase) *FarmHandler {
	return &FarmHandler{
		usecase: usecase,
	}
}

func (h *FarmHandler) CreateFarm(c *gin.Context) {
	var container domain.Farm
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, "fail to bind request body", false, nil)
		return
	}
	result, err := h.usecase.CreateFarm(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create farm", true, result)
}
