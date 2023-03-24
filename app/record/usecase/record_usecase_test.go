package usecase

import (
	"testing"

	"github.com/VinncentWong/Delos-AquaFarm/domain/mock/record"
	"github.com/VinncentWong/Delos-AquaFarm/domain/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var recordRepository record.RecordRepositoryMock = record.RecordRepositoryMock{
	Mock: mock.Mock{},
}

var recordUsecase IRecordUsecase = NewRecordUsecase(&recordRepository)

func TestGetRecord(t *testing.T) {
	data := [][]response.RecordResponse{
		{
			{
				MethodName: "GET",
				Endpoint:   "/record",
			},
			{
				MethodName: "POST",
				Endpoint:   "/record/create",
			},
			{
				MethodName: "PUT",
				Endpoint:   "/record/update",
			},
			{
				MethodName: "DELETE",
				Endpoint:   "/record/delete",
			},
		},
		{
			{
				MethodName: "GET",
				Endpoint:   "/pond",
			},
			{
				MethodName: "POST",
				Endpoint:   "/pond/create",
			},
			{
				MethodName: "PUT",
				Endpoint:   "/pond/update",
			},
			{
				MethodName: "DELETE",
				Endpoint:   "/pond/delete",
			},
		},
		{
			{
				MethodName: "GET",
				Endpoint:   "/farm",
			},
			{
				MethodName: "POST",
				Endpoint:   "/farm/create",
			},
			{
				MethodName: "PUT",
				Endpoint:   "/farm/update",
			},
			{
				MethodName: "DELETE",
				Endpoint:   "/farm/delete",
			},
		},
	}
	for i, d := range data {
		t.Run("record should success get all data", func(t *testing.T) {
			call1 := recordRepository.Mock.On("GetRecord").Return(d, nil)
			result, err := recordUsecase.GetRecord()
			assert.Nil(t, err, "should return nil")
			assert.Len(t, result, 4, "length should be equal to 3")
			j := 0
			for _, r := range result {
				assert.Equal(t, data[i][j].Endpoint, r.Endpoint, "endpoint should be equal")
				j++
			}
			call1.Unset()
		})
	}
}

func TestGetUniqueAgent(t *testing.T) {
	data := []struct {
		endpoint     string
		nUniqueAgent int
	}{
		{
			endpoint:     "/record/get",
			nUniqueAgent: 1,
		},
		{
			endpoint:     "/record/create",
			nUniqueAgent: 2,
		},
		{
			endpoint:     "/record/update",
			nUniqueAgent: 3,
		},
		{
			endpoint:     "/record/delete",
			nUniqueAgent: 4,
		},
		{
			endpoint:     "/farm/get",
			nUniqueAgent: 5,
		},
	}
	for _, d := range data {
		t.Run("get unique agent should success return number of unique agent", func(t *testing.T) {
			call1 := recordRepository.Mock.On("GetUniqueAgent", d.endpoint).Return(d.nUniqueAgent, nil)
			result, err := recordUsecase.GetUniqueAgent(d.endpoint)
			assert.Nil(t, err, "should return nil")
			assert.Equal(t, d.nUniqueAgent, result, "number of unique agent should be equal")
			call1.Unset()
		})
	}
}
