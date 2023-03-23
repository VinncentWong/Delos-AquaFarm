package handler

import (
	"fmt"
	"net/http"

	"github.com/VinncentWong/Delos-AquaFarm/app/record/usecase"
	"github.com/VinncentWong/Delos-AquaFarm/util"
	"github.com/gin-gonic/gin"
)

type RecordHandler struct {
	usecase usecase.IRecordUsecase
}

func NewRecordHandler(usecase usecase.IRecordUsecase) *RecordHandler {
	return &RecordHandler{
		usecase: usecase,
	}
}

func (h *RecordHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetRecord()
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	if len(result) == 0 {
		util.SendResponse(c, http.StatusNotFound, "record doesn't exist", false, nil)
		return
	}
	var response []map[string]map[string]int
	for _, record := range result {
		result, err := h.usecase.GetUniqueAgent(record.Endpoint)
		if err != nil {
			continue
		}
		content := make(map[string]int)
		content["count"] = int(record.Count)
		content["unique_user_agent"] = result
		response = append(response, map[string]map[string]int{
			fmt.Sprintf("%s %s", record.MethodName, record.Endpoint): content,
		})
	}
	util.SendResponse(c, http.StatusOK, "success get records data", true, response)
}
