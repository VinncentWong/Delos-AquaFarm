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
	if args[1] == nil {
		return farm, nil
	}
	return domain.Farm{}, args[1].(error)
}

func (m *FarmUsecaseMock) UpdateFarm(param *domain.Farm) error {
	args := m.Mock.Called(param)
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *FarmUsecaseMock) DeleteFarm(id string) error {
	args := m.Mock.Called(id)
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *FarmUsecaseMock) GetFarmById(id string) (domain.Farm, error) {
	args := m.Mock.Called(id)
	farm := args[0].(domain.Farm)
	if args[1] == nil {
		return farm, nil
	}
	return domain.Farm{}, args[1].(error)
}

func (m *FarmUsecaseMock) GetAll() ([]domain.Farm, error) {
	args := m.Mock.Called()
	if args[1] == nil {
		return args[0].([]domain.Farm), nil
	}
	return []domain.Farm{}, args[1].(error)
}
