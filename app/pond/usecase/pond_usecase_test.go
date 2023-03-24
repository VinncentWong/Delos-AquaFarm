package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/domain/mock/farm"
	"github.com/VinncentWong/Delos-AquaFarm/domain/mock/pond"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var pondRepositoryMock pond.PondRepositoryMock = pond.PondRepositoryMock{
	Mock: mock.Mock{},
}

var farmRepositoryMock farm.FarmRepositoryMock = farm.FarmRepositoryMock{
	Mock: mock.Mock{},
}

var pondUsecase IPondUsecase = NewPondUsecase(&pondRepositoryMock, &farmRepositoryMock)

func TestCreatePond(t *testing.T) {
	data := []struct {
		id   string
		pond *domain.Pond
	}{
		{
			id: "1",
			pond: &domain.Pond{
				FarmID: "1",
				Name:   "Pond1",
				Farm: domain.Farm{
					FarmName: "Farm1",
				},
			},
		},
		{
			id: "2",
			pond: &domain.Pond{
				FarmID: "2",
				Name:   "Pond2",
				Farm: domain.Farm{
					FarmName: "Farm2",
				},
			},
		},
		{
			id: "3",
			pond: &domain.Pond{
				FarmID: "3",
				Name:   "Pond3",
				Farm: domain.Farm{
					FarmName: "Farm3",
				},
			},
		},
		{
			id: "4",
			pond: &domain.Pond{
				FarmID: "4",
				Name:   "Pond4",
				Farm: domain.Farm{
					FarmName: "Farm4",
				},
			},
		},
		{
			id: "5",
			pond: &domain.Pond{
				FarmID: "5",
				Name:   "Pond5",
				Farm: domain.Farm{
					FarmName: "Farm5",
				},
			},
		},
	}
	for _, d := range data {
		t.Run("create pond should success", func(t *testing.T) {
			farmRepositoryMock.Mock.On("GetFarmById", d.id).Return(d.pond.Farm, nil)
			call1 := pondRepositoryMock.Mock.On("GetPondByName", d.pond.Name).Return(domain.Pond{}, errors.New("pond doesn't exist")).Once()
			call2 := pondRepositoryMock.Mock.On("CreatePond", d.id, d.pond).Return(d.pond, nil)
			pondRepositoryMock.Mock.On("GetPondByName", d.pond.Name).Return(d.pond, nil)
			result, err := pondUsecase.CreatePond(d.id, d.pond)
			assert.Nil(t, err, "should not return an error")
			assert.Equal(t, d.pond.Name, result.Name, "pond name should be equal")
			assert.Equal(t, d.pond.ID, result.ID, "pond id should be equal")
			assert.Equal(t, d.pond.Farm.FarmName, result.Farm.FarmName, "farm name should be equal")

			call1.Unset()
			call2.Unset()
		})
	}
}

func TestUpdatePond(t *testing.T) {
	data := []struct {
		UpdateData domain.Pond
		ResultData domain.Pond
	}{
		{
			UpdateData: domain.Pond{
				Model: gorm.Model{
					ID: 1,
				},
				Name: "UpdatePond1",
				Farm: domain.Farm{
					FarmName: "Farm1",
				},
			},
			ResultData: domain.Pond{
				Model: gorm.Model{
					ID: 1,
				},
				Name: "Pond1",
				Farm: domain.Farm{
					FarmName: "Farm1",
				},
			},
		},
		{
			UpdateData: domain.Pond{
				Model: gorm.Model{
					ID: 2,
				},
				Name: "UpdatePond2",
				Farm: domain.Farm{
					FarmName: "Farm2",
				},
			},
			ResultData: domain.Pond{
				Model: gorm.Model{
					ID: 2,
				},
				Name: "Pond2",
				Farm: domain.Farm{
					FarmName: "Farm2",
				},
			},
		},
		{
			UpdateData: domain.Pond{
				Model: gorm.Model{
					ID: 3,
				},
				Name: "UpdatePond3",
				Farm: domain.Farm{
					FarmName: "Farm3",
				},
			},
			ResultData: domain.Pond{
				Model: gorm.Model{
					ID: 3,
				},
				Name: "Pond3",
				Farm: domain.Farm{
					FarmName: "Farm3",
				},
			},
		},
		{
			UpdateData: domain.Pond{
				Model: gorm.Model{
					ID: 4,
				},
				Name: "UpdatePond4",
				Farm: domain.Farm{
					FarmName: "Farm4",
				},
			},
			ResultData: domain.Pond{
				Model: gorm.Model{
					ID: 4,
				},
				Name: "Pond4",
				Farm: domain.Farm{
					FarmName: "Farm4",
				},
			},
		},
		{
			UpdateData: domain.Pond{
				Model: gorm.Model{
					ID: 5,
				},
				Name: "UpdatePond5",
				Farm: domain.Farm{
					FarmName: "Farm5",
				},
			},
			ResultData: domain.Pond{
				Model: gorm.Model{
					ID: 5,
				},
				Name: "Pond5",
				Farm: domain.Farm{
					FarmName: "Farm5",
				},
			},
		},
	}
	for _, d := range data {
		t.Run("update pond should success", func(t *testing.T) {
			call1 := pondRepositoryMock.Mock.On("GetPondById", fmt.Sprint(d.UpdateData.ID)).Return(&d.ResultData, nil)
			call2 := pondRepositoryMock.Mock.On("UpdatePond", &d.UpdateData).Return(nil)
			result, err := pondUsecase.UpdatePond(&d.UpdateData)
			assert.Nil(t, err, "should not return an error")
			assert.Equal(t, d.UpdateData.Name, result.Name, "pond name should be updated")

			call1.Unset()
			call2.Unset()
		})
	}
}

func TestDeletePond(t *testing.T) {
	data := []struct {
		id   string
		pond domain.Pond
	}{
		{
			id: "1",
			pond: domain.Pond{
				Name: "Pond1",
			},
		},
		{
			id: "2",
			pond: domain.Pond{
				Name: "Pond2",
			},
		},
		{
			id: "3",
			pond: domain.Pond{
				Name: "Pond3",
			},
		},
		{
			id: "4",
			pond: domain.Pond{
				Name: "Pond4",
			},
		},
		{
			id: "5",
			pond: domain.Pond{
				Name: "Pond5",
			},
		},
	}
	for _, d := range data {
		t.Run("delete pond should be success", func(t *testing.T) {
			call1 := pondRepositoryMock.Mock.On("GetPondById", d.id).Return(d.pond, nil)
			call2 := pondRepositoryMock.Mock.On("DeletePond", d.id).Return(nil)
			err := pondUsecase.DeletePond(d.id)
			assert.Nil(t, err, "should return nil")

			call1.Unset()
			call2.Unset()
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
		t.Run("pond should success get all pond", func(t *testing.T) {
			call1 := pondRepositoryMock.Mock.On("GetAll").Return(&d, nil)
			result, err := pondUsecase.GetAll()
			assert.Nil(t, err, "should return nil")
			assert.Len(t, result, 3)
			j := 0
			for _, r := range result {
				assert.Equal(t, data[i][j].Name, r.Name, "pond name should be same")
				j++
			}
			call1.Unset()
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
		t.Run("get pond by id should success", func(t *testing.T) {
			call1 := pondRepositoryMock.Mock.On("GetPondById", d.id).Return(&(d.pond), nil)
			result, err := pondUsecase.GetPondById(d.id)
			assert.Nil(t, err, "should return nil")
			assert.Equal(t, d.pond.Name, result.Name, "pond name should be equal")
			assert.Equal(t, d.pond.ID, result.ID, "pond id should be equal")
			call1.Unset()
		})
	}
}
