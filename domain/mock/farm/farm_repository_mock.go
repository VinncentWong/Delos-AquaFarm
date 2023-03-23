package farm

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/stretchr/testify/mock"
)

type FarmRepositoryMock struct {
	Mock mock.Mock
}

func (m *FarmRepositoryMock) CreateFarm(farm *domain.Farm) (domain.Farm, error) {
	args := m.Mock.Called(farm)
	farms := args[0].(domain.Farm)
	err := args[1].(error)
	if err != nil {
		return domain.Farm{}, err
	}
	return farms, nil
}

func (m *FarmRepositoryMock) GetFarmByName(name string) (domain.Farm, error) {
	args := m.Mock.Called(name)
	farm := args[0].(domain.Farm)
	err := args[1].(error)
	if err != nil {
		return domain.Farm{}, err
	}
	return farm, nil
}

func (m *FarmRepositoryMock) GetFarmById(id string) (domain.Farm, error) {
	args := m.Mock.Called(id)
	farm := args[0].(domain.Farm)
	err := args[1].(error)
	if err != nil {
		return domain.Farm{}, err
	}
	return farm, err
}

func (m *FarmRepositoryMock) UpdateFarm(farm *domain.Farm) error {
	args := m.Mock.Called(farm)
	err := args[0].(error)
	return err
}

func (m *FarmRepositoryMock) DeleteFarm(id string) error {
	args := m.Mock.Called(id)
	err := args[0].(error)
	return err
}

func (m *FarmRepositoryMock) GetAll() ([]domain.Farm, error) {
	args := m.Mock.Called()
	farms := args[0].([]domain.Farm)
	err := args[1].(error)
	if err != nil {
		return []domain.Farm{}, err
	}
	return farms, nil
}
