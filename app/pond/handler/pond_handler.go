package handler

import (
	"net/http"
	"strconv"

	"github.com/VinncentWong/Delos-AquaFarm/app/pond/usecase"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	err = validator.New().Struct(&container)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, errs.Error(), false, nil)
		return
	}
	result, err := h.usecase.CreatePond(farmId, &container)
	if err != nil {
		switch err.Error() {
		case "farm id doesn't exist":
			util.SendResponse(c, http.StatusNotFound, err.Error(), false, nil)
		default:
			util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		}
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create pond", true, result)
}

func (h *PondHandler) UpdatePond(c *gin.Context) {
	pondId := c.Param("pondId")
	var container domain.Pond
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
	pondIdInt, err := strconv.Atoi(pondId)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	container.ID = uint(pondIdInt)
	err = h.usecase.UpdatePond(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success update pond", true, nil)
}

func (h *PondHandler) DeletePond(c *gin.Context) {
	pondId := c.Param("pondId")
	err := h.usecase.DeletePond(pondId)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success delete pond", true, nil)
}

func (h *PondHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll()
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get ponds", true, result)
}
