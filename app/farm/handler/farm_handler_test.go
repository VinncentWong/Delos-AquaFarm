package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/domain/mock/farm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var farmUsecaseMock farm.FarmUsecaseMock = farm.FarmUsecaseMock{
	Mock: mock.Mock{},
}

var farmHandler *FarmHandler = NewFarmHandler(&farmUsecaseMock)

func TestCreateFarm(t *testing.T) {
	data := []domain.Farm{
		{
			FarmName: "Farm1",
		},
		{
			FarmName: "Farm2",
		},
		{
			FarmName: "Farm3",
		},
		{
			FarmName: "Farm4",
		},
		{
			FarmName: "Farm5",
		},
	}
	for _, d := range data {
		t.Run("create farm handler should success", func(t *testing.T) {
			router := gin.Default()
			call1 := farmUsecaseMock.Mock.On("CreateFarm", &d).Return(d, nil)
			router.POST("/farm/create", farmHandler.CreateFarm)
			b, _ := json.Marshal(d)
			req, _ := http.NewRequest("POST", "/farm/create", bytes.NewBuffer(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured : ", err.Error())
			}
			dataContainer := container["data"].(map[string]any)
			assert.Equal(t, http.StatusCreated, w.Code, "should return http status code created")
			assert.Equal(t, d.FarmName, dataContainer["name"], "farm name should be equal")
			assert.Equal(t, "success create farm", container["message"], "message should be equal")
			assert.Equal(t, true, container["success"], "success should be true")
			call1.Unset()
		})
	}
}

func TestUpdateFarm(t *testing.T) {
	data := []struct {
		farmId string
		farm   domain.Farm
	}{
		{
			farmId: "1",
			farm: domain.Farm{
				FarmName: "Farm1",
				Model: gorm.Model{
					ID: 1,
				},
			},
		},
		{
			farmId: "2",
			farm: domain.Farm{
				FarmName: "Farm2",
				Model: gorm.Model{
					ID: 2,
				},
			},
		},
		{
			farmId: "3",
			farm: domain.Farm{
				FarmName: "Farm3",
				Model: gorm.Model{
					ID: 3,
				},
			},
		},
		{
			farmId: "4",
			farm: domain.Farm{
				FarmName: "Farm4",
				Model: gorm.Model{
					ID: 4,
				},
			},
		},
		{
			farmId: "5",
			farm: domain.Farm{
				FarmName: "Farm5",
				Model: gorm.Model{
					ID: 5,
				},
			},
		},
	}
	for _, d := range data {
		t.Run("update farm handler should success", func(t *testing.T) {
			router := gin.Default()
			router.PUT("/farm/update/:farmId", farmHandler.UpdateFarm)
			call1 := farmUsecaseMock.Mock.On("UpdateFarm", &(d.farm)).Return(nil)
			w := httptest.NewRecorder()
			bData, _ := json.Marshal(d.farm)
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/farm/update/%s", d.farmId), bytes.NewBuffer(bData))
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			assert.Equal(t, http.StatusOK, w.Code, "should return http status ok")
			assert.Equal(t, true, container["success"], "should return success true")
			assert.Equal(t, "success update farm", container["message"], "message should be equal")
			call1.Unset()
		})
	}
}

func TestDeleteFarm(t *testing.T) {
	data := []string{"1", "2", "3", "4", "5"}
	for _, d := range data {
		t.Run("delete farm should success", func(t *testing.T) {
			call1 := farmUsecaseMock.Mock.On("DeleteFarm", d).Return(nil)
			router := gin.Default()
			router.DELETE("/farm/delete/:farmId", farmHandler.DeleteFarm)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/farm/delete/%s", d), nil)
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			assert.Equal(t, http.StatusOK, w.Code, "should return http status ok")
			assert.Equal(t, "success delete farm", container["message"], "message should be equal")
			assert.Equal(t, true, container["success"], "success should be equal to true")
			call1.Unset()
		})
	}
}

func TestGetAll(t *testing.T) {
	data := [][]domain.Farm{
		[]domain.Farm{
			{
				FarmName: "Farm Array 1_1",
			},
			{
				FarmName: "Farm Array 1_2",
			},
			{
				FarmName: "Farm Array 1_3",
			},
		},
		[]domain.Farm{
			{
				FarmName: "Farm Array 2_1",
			},
			{
				FarmName: "Farm Array 2_2",
			},
			{
				FarmName: "Farm Array 2_3",
			},
		},
		[]domain.Farm{
			{
				FarmName: "Farm Array 3_1",
			},
			{
				FarmName: "Farm Array 3_2",
			},
			{
				FarmName: "Farm Array 3_3",
			},
		},
		[]domain.Farm{
			{
				FarmName: "Farm Array 4_1",
			},
			{
				FarmName: "Farm Array 4_2",
			},
			{
				FarmName: "Farm Array 4_3",
			},
		},
		[]domain.Farm{
			{
				FarmName: "Farm Array 5_1",
			},
			{
				FarmName: "Farm Array 5_2",
			},
			{
				FarmName: "Farm Array 5_3",
			},
		},
	}
	for i, d := range data {
		t.Run("get all farm handler should success", func(t *testing.T) {
			caller1 := farmUsecaseMock.Mock.On("GetAll").Return(d, nil)
			router := gin.Default()
			router.GET("/farm/gets", farmHandler.GetAll)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/farm/gets", nil)
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			bodyContainer := container["data"].([]interface{})
			j := 0
			for _, tempDataBody := range bodyContainer {
				dataBody := tempDataBody.(map[string]any)
				assert.Equal(t, data[i][j].FarmName, dataBody["name"], "farm name should be equal")
				j++
			}
			assert.Equal(t, http.StatusOK, w.Code, "http status code should be equal")
			assert.Equal(t, "success get farms", container["message"], "message should be equal")
			assert.Equal(t, true, container["success"], "success should be equal to true")
			caller1.Unset()
		})
	}
}

func TestGetFarmById(t *testing.T) {
	data := []struct {
		id   string
		farm domain.Farm
	}{
		{
			id: "1",
			farm: domain.Farm{
				Model: gorm.Model{
					ID: 1,
				},
				FarmName: "Farm1",
			},
		},
		{
			id: "2",
			farm: domain.Farm{
				Model: gorm.Model{
					ID: 2,
				},
				FarmName: "Farm2",
			},
		},
		{
			id: "3",
			farm: domain.Farm{
				Model: gorm.Model{
					ID: 3,
				},
				FarmName: "Farm3",
			},
		},
		{
			id: "4",
			farm: domain.Farm{
				Model: gorm.Model{
					ID: 4,
				},
				FarmName: "Farm4",
			},
		},
		{
			id: "5",
			farm: domain.Farm{
				Model: gorm.Model{
					ID: 5,
				},
				FarmName: "Farm5",
			},
		},
	}
	for _, d := range data {
		t.Run("get farm by id should success", func(t *testing.T) {
			call1 := farmUsecaseMock.Mock.On("GetFarmById", d.id).Return(d.farm, nil)
			router := gin.Default()
			router.GET("/farm/get/:farmId", farmHandler.GetFarmById)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/farm/get/%s", d.id), nil)
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			bodyContainer := container["data"].(map[string]any)
			assert.Equal(t, d.farm.FarmName, bodyContainer["name"], "farm name should be equal")
			assert.Equal(t, http.StatusOK, w.Code, "should return http status ok")
			assert.Equal(t, "success get farm", container["message"], "message should be equal")
			assert.Equal(t, true, container["success"], "success should be equal to true")
			call1.Unset()
		})
	}
}
