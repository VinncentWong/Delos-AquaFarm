package usecase

import (
	"errors"
	"testing"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	farmMock "github.com/VinncentWong/Delos-AquaFarm/domain/mock/farm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var fRepo farmMock.FarmRepositoryMock = farmMock.FarmRepositoryMock{
	Mock: mock.Mock{},
}

var fUsecase IFarmUsecase = NewFarmUsecase(&fRepo)

func TestCreateFarm(t *testing.T) {
	data := []domain.Farm{
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
			call1 := fRepo.Mock.On("GetFarmByName", d.FarmName).Return(d, errors.New("repository doesn't found farm"))
			call2 := fRepo.Mock.On("CreateFarm", &d).Return(d, nil)
			result, err := fUsecase.CreateFarm(&d)
			assert.Nil(t, err, "should not return an error")
			assert.Equal(t, d.FarmName, result.FarmName, "name should be equal")
			call1.Unset()
			call2.Unset()
		})
	}
}

func TestUpdateFarm(t *testing.T) {
	data := []domain.Farm{
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
			call1 := fRepo.Mock.On("UpdateFarm", &d).Return(nil)
			call2 := fRepo.Mock.On("GetFarmByName", d.FarmName).Return(d, errors.New("farm doesn't exist"))
			err := fUsecase.UpdateFarm(&d)
			assert.Nil(t, err, "err should be nil")
			call1.Unset()
			call2.Unset()
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
			call1 := fRepo.Mock.On("GetFarmById", d).Return(domain.Farm{}, nil)
			call2 := fRepo.Mock.On("DeleteFarm", d).Return(nil)
			err := fUsecase.DeleteFarm(d)
			assert.Nil(t, err, "should be nil")
			call1.Unset()
			call2.Unset()
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
			call1 := fRepo.Mock.On("GetAll").Return(d, nil)
			result, err := fUsecase.GetAll()
			assert.Nil(t, err, "should be return nil")
			assert.Equal(t, len(d), len(result), "length of result with expected should be same")
			j := 0
			for _, r := range result {
				assert.Equal(t, data[i][j].FarmName, r.FarmName, "farm name should be same")
				j++
			}
			call1.Unset()
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
			caller1 := fRepo.Mock.On("GetFarmById", d.id).Return(d.farm, nil)
			result, err := fUsecase.GetFarmById(d.id)
			assert.Nil(t, err, "should return nil")
			assert.Equal(t, d.farm.FarmName, result.FarmName, "farm name should be equal")
			caller1.Unset()
		})
	}
}
