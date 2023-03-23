package farm

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/stretchr/testify/mock"
)

/*
CreateFarm(farm *domain.Farm) (domain.Farm, error)
UpdateFarm(farm *domain.Farm) error
DeleteFarm(id string) error
GetAll() ([]domain.Farm, error)
GetFarmById(id string) (domain.Farm, error)
*/
type FarmUsecaseMock struct {
	Mock mock.Mock
}

func (m *FarmUsecaseMock) CreateFarm(param *domain.Farm) (domain.Farm, error) {
	args := m.Mock.Called(param)
	farm := args[0].(domain.Farm)
	err := args[1].(error)
	if err != nil {
		return domain.Farm{}, err
	}
	return farm, nil
}

func (m *FarmUsecaseMock) UpdateFarm(param *domain.Farm) error {
	args := m.Mock.Called(param)
	err := args[0].(error)
	return err
}

func (m *FarmUsecaseMock) DeleteFarm(id string) error {
	args := m.Mock.Called(id)
	err := args[0].(error)
	return err
}

func (m *FarmUsecaseMock) GetFarmById(id string) (domain.Farm, error) {
	args := m.Mock.Called(id)
	farm := args[0].(domain.Farm)
	err := args[1].(error)
	if err != nil {
		return domain.Farm{}, err
	}
	return farm, nil
}
