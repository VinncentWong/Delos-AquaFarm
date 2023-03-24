package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/domain/mock/pond"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var pondUsecaseMock pond.PondUsecaseMock = pond.PondUsecaseMock{
	Mock: mock.Mock{},
}

var pondHandler *PondHandler = NewPondHandler(&pondUsecaseMock)

func TestCreatePond(t *testing.T) {
	data := []struct {
		farmId string
		pond   domain.Pond
	}{
		{
			farmId: "1",
			pond: domain.Pond{
				FarmID: "1",
				Name:   "Pond1",
				Model: gorm.Model{
					ID: 1,
				},
			},
		},
		{
			farmId: "2",
			pond: domain.Pond{
				FarmID: "2",
				Name:   "Pond2",
				Model: gorm.Model{
					ID: 2,
				},
			},
		},
		{
			farmId: "3",
			pond: domain.Pond{
				FarmID: "3",
				Name:   "Pond3",
				Model: gorm.Model{
					ID: 3,
				},
			},
		},
		{
			farmId: "4",
			pond: domain.Pond{
				FarmID: "4",
				Name:   "Pond4",
				Model: gorm.Model{
					ID: 4,
				},
			},
		},
		{
			farmId: "5",
			pond: domain.Pond{
				FarmID: "5",
				Name:   "Pond5",
				Model: gorm.Model{
					ID: 5,
				},
			},
		},
	}
	for _, d := range data {
		t.Run("create pond should success", func(t *testing.T) {
			caller1 := pondUsecaseMock.Mock.On("CreatePond", d.farmId, &(d.pond)).Return(d.pond, nil)
			router := gin.Default()
			router.POST("/pond/create/:farmId", pondHandler.CreatePond)
			w := httptest.NewRecorder()
			dBody, _ := json.Marshal(d.pond)
			req, _ := http.NewRequest("POST", fmt.Sprintf("/pond/create/%s", d.farmId), bytes.NewBuffer(dBody))
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			bodyContainer := container["data"].(map[string]any)
			assert.Equal(t, http.StatusCreated, w.Code, "status should be created(201)")
			assert.Equal(t, "success create pond", container["message"], "message should be equal")
			assert.Equal(t, true, container["success"], "success value should be equal to true")
			assert.Equal(t, float64(d.pond.ID), bodyContainer["ID"], "id of the pond shold be equal")
			caller1.Unset()
		})
	}
}

func TestUpdatePond(t *testing.T) {
	data := []struct {
		id   string
		pond domain.Pond
	}{
		{
			id: "1",
			pond: domain.Pond{
				Name: "Pond1",
				Model: gorm.Model{
					ID: 1,
				},
			},
		},
		{
			id: "2",
			pond: domain.Pond{
				Name: "Pond2",
				Model: gorm.Model{
					ID: 2,
				},
			},
		},
		{
			id: "3",
			pond: domain.Pond{
				Name: "Pond3",
				Model: gorm.Model{
					ID: 3,
				},
			},
		},
		{
			id: "4",
			pond: domain.Pond{
				Name: "Pond4",
				Model: gorm.Model{
					ID: 4,
				},
			},
		},
		{
			id: "5",
			pond: domain.Pond{
				Name: "Pond5",
				Model: gorm.Model{
					ID: 5,
				},
			},
		},
	}
	for _, d := range data {
		t.Run("update pond should be success", func(t *testing.T) {
			caller1 := pondUsecaseMock.Mock.On("UpdatePond", &(d.pond)).Return(d.pond, nil)
			router := gin.Default()
			router.PUT("/pond/update/:pondId", pondHandler.UpdatePond)
			w := httptest.NewRecorder()
			dBody, _ := json.Marshal(d.pond)
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/pond/update/%s", d.id), bytes.NewBuffer(dBody))
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			assert.Equal(t, http.StatusOK, w.Code, "response code should be equal to status ok")
			assert.Equal(t, true, container["success"], "success value should be equal to true")
			assert.Equal(t, "success update pond", container["message"], "message should be equal")
			assert.Nil(t, container["data"], "data should be nil")
			caller1.Unset()
		})
	}
}

func TestDeletePond(t *testing.T) {
	data := []string{"1", "2", "3", "4", "5"}
	for _, d := range data {
		t.Run("delete pond should be success", func(t *testing.T) {
			caller1 := pondUsecaseMock.Mock.On("DeletePond", d).Return(nil)
			router := gin.Default()
			router.DELETE("/pond/update/:pondId", pondHandler.DeletePond)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/pond/update/%s", d), nil)
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			assert.Equal(t, http.StatusOK, w.Code, "response code should be equal")
			assert.Equal(t, "success delete pond", container["message"], "message response should be equal")
			assert.Equal(t, true, container["success"], "success value should be equal")
			caller1.Unset()
		})
	}
}

func TestGetAll(t *testing.T) {
	data := [][]domain.Pond{
		[]domain.Pond{
			{
				Name: "Pond1",
			},
			{
				Name: "Pond2",
			},
			{
				Name: "Pond3",
			},
		},
		[]domain.Pond{
			{
				Name: "Pond4",
			},
			{
				Name: "Pond5",
			},
			{
				Name: "Pond6",
			},
		},
		[]domain.Pond{
			{
				Name: "Pond7",
			},
			{
				Name: "Pond8",
			},
			{
				Name: "Pond9",
			},
		},
		[]domain.Pond{
			{
				Name: "Pond10",
			},
			{
				Name: "Pond11",
			},
			{
				Name: "Pond12",
			},
		},
		[]domain.Pond{
			{
				Name: "Pond13",
			},
			{
				Name: "Pond14",
			},
			{
				Name: "Pond15",
			},
		},
	}
	for i, d := range data {
		t.Run("get all pond should be success", func(t *testing.T) {
			caller1 := pondUsecaseMock.Mock.On("GetAll").Return(d, nil)
			router := gin.Default()
			router.GET("/pond/gets", pondHandler.GetAll)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/pond/gets", nil)
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
				assert.Equal(t, data[i][j].Name, dataBody["name"], "pond name should be equal")
				j++
			}
			assert.Equal(t, http.StatusOK, w.Code, "http response should be equal")
			assert.Equal(t, "success get ponds", container["message"], "message should be equal")
			assert.Equal(t, true, container["success"], "success value should be equal to true")
			caller1.Unset()
		})
	}
}

func TestGetPondById(t *testing.T) {
	data := []struct {
		id   string
		pond domain.Pond
	}{
		{
			id: "1",
			pond: domain.Pond{
				Model: gorm.Model{
					ID: 1,
				},
				Name: "Pond1",
			},
		},
		{
			id: "2",
			pond: domain.Pond{
				Model: gorm.Model{
					ID: 2,
				},
				Name: "Pond2",
			},
		},
		{
			id: "3",
			pond: domain.Pond{
				Model: gorm.Model{
					ID: 3,
				},
				Name: "Pond3",
			},
		},
		{
			id: "4",
			pond: domain.Pond{
				Model: gorm.Model{
					ID: 4,
				},
				Name: "Pond4",
			},
		},
		{
			id: "5",
			pond: domain.Pond{
				Model: gorm.Model{
					ID: 5,
				},
				Name: "Pond5",
			},
		},
	}
	for _, d := range data {
		t.Run("get pond by id should be success", func(t *testing.T) {
			caller1 := pondUsecaseMock.Mock.On("GetPondById", d.id).Return(d.pond, nil)
			router := gin.Default()
			router.GET("/pond/get/:pondId", pondHandler.GetPondById)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/pond/get/%s", d.id), nil)
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			bodyContainer := container["data"].(map[string]any)
			assert.Equal(t, http.StatusOK, w.Code, "http response should be equal")
			assert.Equal(t, "success get pond", container["message"], "message should be equal")
			assert.Equal(t, true, container["success"], "success value should be equal to true")
			assert.Equal(t, d.pond.Name, bodyContainer["name"], "pond name should be equal")
			caller1.Unset()
		})
	}
}
