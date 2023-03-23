package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	farmMock "github.com/VinncentWong/Delos-AquaFarm/domain/mock/farm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var fRepo farmMock.FarmRepositoryMock = farmMock.FarmRepositoryMock{
	Mock: mock.Mock{},
}

var fUsecase IFarmUsecase = NewFarmUsecase(&fRepo)

func TestCreateFarm(t *testing.T) {
	data := []*domain.Farm{
		{
			FarmName: "farm1",
		},
		{
			FarmName: "farm2",
		},
		{
			FarmName: "farm3",
		},
		{
			FarmName: "farm4",
		},
		{
			FarmName: "farm5",
		},
	}
	for _, d := range data {
		t.Run("create farm method should success and return same entity", func(t *testing.T) {
			fRepo.Mock.On("GetFarmByName", d.FarmName).Return(d, errors.New("repository doesn't found farm"))
			fRepo.Mock.On("CreateFarm", d).Return(d, nil)
			result, err := fUsecase.CreateFarm(d)
			assert.Nil(t, err, "should not return an error")
			assert.Equal(t, d.FarmName, result.FarmName, "name should be equal")
		})
	}
}

func TestUpdateFarm(t *testing.T) {
	data := []*domain.Farm{
		{
			FarmName: "farm1",
		},
		{
			FarmName: "farm2",
		},
		{
			FarmName: "farm3",
		},
		{
			FarmName: "farm4",
		},
		{
			FarmName: "farm5",
		},
	}
	for _, d := range data {
		t.Run("update farm should be success", func(t *testing.T) {
			fRepo.Mock.On("UpdateFarm", d).Return(nil)
			err := fUsecase.UpdateFarm(d)
			assert.Nil(t, err, "err should be nil")
		})
	}
}

func TestDeleteFarm(t *testing.T) {
	data := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}
	for _, d := range data {
		t.Run("delete farm should be success", func(t *testing.T) {
			fRepo.Mock.On("GetFarmById", d).Return(domain.Farm{}, nil)
			fRepo.Mock.On("DeleteFarm", d).Return(nil)
			err := fUsecase.DeleteFarm(d)
			assert.Nil(t, err, "should be nil")
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
		t.Run("get all farm should be success", func(t *testing.T) {
			fRepo.Mock.On("GetAll").Return(&d, nil)
			result, err := fUsecase.GetAll()
			assert.Nil(t, err, "should be return nil")
			assert.Equal(t, len(d), len(result), "length of result with expected should be same")
			j := 0
			for _, r := range result {
				fmt.Println("r in test = ", r.FarmName)
				assert.Equal(t, data[i][j].FarmName, r.FarmName, "farm name should be same")
				j++
			}
		})
	}
}
