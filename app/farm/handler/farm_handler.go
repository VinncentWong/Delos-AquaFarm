package handler

import (
	"net/http"
	"strconv"

	"github.com/VinncentWong/Delos-AquaFarm/app/farm/usecase"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	err = validator.New().Struct(&container)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, errs.Error(), false, nil)
		return
	}
	result, err := h.usecase.CreateFarm(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create farm", true, result)
}

func (h *FarmHandler) UpdateFarm(c *gin.Context) {
	id := c.Param("farmId")
	var container domain.Farm
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	err = validator.New().Struct(&container)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, errs.Error(), false, nil)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	container.ID = uint(idInt)
	err = h.usecase.UpdateFarm(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success update farm", true, nil)
}

func (h *FarmHandler) DeleteFarm(c *gin.Context) {
	farmId := c.Param("farmId")
	err := h.usecase.DeleteFarm(farmId)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success delete farm", true, nil)
}

func (h *FarmHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll()
	if err == nil {
		util.SendResponse(c, http.StatusOK, "success get farms", true, result)
		return
	}
	switch err.Error() {
	case "farms doesn't exist":
		util.SendResponse(c, http.StatusNotFound, err.Error(), false, nil)
		return
	default:
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
}

func (h *FarmHandler) GetFarmById(c *gin.Context) {
	farmId := c.Param("farmId")
	result, err := h.usecase.GetFarmById(farmId)
	if err != nil {
		util.SendResponse(c, http.StatusNotFound, "farm not found", false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get farm", true, result)
}
