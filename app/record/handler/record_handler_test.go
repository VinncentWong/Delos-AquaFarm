package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VinncentWong/Delos-AquaFarm/domain/mock/record"
	"github.com/VinncentWong/Delos-AquaFarm/domain/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var recordUsecaseMock record.RecordUsecaseMock = record.RecordUsecaseMock{
	Mock: mock.Mock{},
}

var recordHandler *RecordHandler = NewRecordHandler(&recordUsecaseMock)

func TestGetAll(t *testing.T) {
	data := [][]response.RecordResponse{
		[]response.RecordResponse{
			{
				MethodName: "POST",
				Endpoint:   "/farm/create",
				Count:      10,
			},
			{
				MethodName: "GET",
				Endpoint:   "/farm/get",
				Count:      20,
			},
			{
				MethodName: "DELETE",
				Endpoint:   "/farm/delete",
				Count:      10,
			},
		},
		[]response.RecordResponse{
			{
				MethodName: "POST",
				Endpoint:   "/pond/create",
				Count:      40,
			},
			{
				MethodName: "GET",
				Endpoint:   "/pond/get",
				Count:      50,
			},
			{
				MethodName: "DELETE",
				Endpoint:   "/pond/delete",
				Count:      60,
			},
		},
		[]response.RecordResponse{
			{
				MethodName: "POST",
				Endpoint:   "/record/create",
				Count:      70,
			},
			{
				MethodName: "GET",
				Endpoint:   "/record/get",
				Count:      80,
			},
			{
				MethodName: "DELETE",
				Endpoint:   "/record/delete",
				Count:      90,
			},
		},
	}
	for i, d := range data {
		t.Run("record handler should success return response", func(t *testing.T) {
			caller1 := recordUsecaseMock.Mock.On("GetRecord").Return(d, nil)
			for _, apiRecord := range d {
				recordUsecaseMock.Mock.On("GetUniqueAgent", apiRecord.Endpoint).Return(i, nil).Once()
			}
			router := gin.Default()
			router.GET("/record/get", recordHandler.GetAll)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/record/get", nil)
			router.ServeHTTP(w, req)
			var container map[string]any
			err := json.Unmarshal(w.Body.Bytes(), &container)
			if err != nil {
				t.Fatal("error occured ", err.Error())
			}
			/*
				body container would be like this
				[map[POST /record/create:map[count:70 unique_user_agent:2]]
				map[GET /record/get:map[count:80 unique_user_agent:2]]
				map[DELETE /record/delete:map[count:90 unique_user_agent:2]]]
			*/
			bodyContainer := container["data"].([]any)
			j := 0
			for _, tempBody := range bodyContainer {
				// b would be equal to this
				// b =  map[POST /record/create:map[count:70 unique_user_agent:2]]
				b := tempBody.(map[string]any)
				/*
					k would be equal to POST /record/create
					v would be equal to map[count:70 unique_user_agent:2]
				*/
				for k, tempV := range b {
					arr := strings.Split(k, " ")
					assert.Equal(t, data[i][j].MethodName, arr[0], "method name should be equal")
					assert.Equal(t, data[i][j].Endpoint, arr[1], "endpoint should be equal")
					v := tempV.(map[string]any)
					assert.Equal(t, float64(data[i][j].Count), v["count"], "count of record should be equal")
					assert.Equal(t, float64(i), v["unique_user_agent"], "number of unique user agent should be equal")
					j++
				}
			}
			assert.Equal(t, http.StatusOK, w.Code, "http response code should be equal")
			assert.Equal(t, true, container["success"], "success should be equal to true")
			assert.Equal(t, "success get records data", container["message"], "message should be equal")
			caller1.Unset()
		})
	}
}
