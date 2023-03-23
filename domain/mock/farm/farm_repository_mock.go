package farm

import (
	"fmt"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/stretchr/testify/mock"
)

type FarmRepositoryMock struct {
	Mock mock.Mock
}

func (m *FarmRepositoryMock) CreateFarm(farm *domain.Farm) (domain.Farm, error) {
	args := m.Mock.Called(farm)
	farms := args[0].(*domain.Farm)
	if args[1] == nil {
		return *farms, nil
	}
	err := args[1].(error)
	return domain.Farm{}, err
}

func (m *FarmRepositoryMock) GetFarmByName(name string) (domain.Farm, error) {
	args := m.Mock.Called(name)
	farm := args[0].(*domain.Farm)
	if args[1] == nil {
		return *farm, nil
	} else {
		err := args[1].(error)
		return domain.Farm{}, err
	}
}

func (m *FarmRepositoryMock) GetFarmById(id string) (domain.Farm, error) {
	args := m.Mock.Called(id)
	farm := args[0].(domain.Farm)
	if args[1] == nil {
		return farm, nil
	}
	err := args[1].(error)
	return domain.Farm{}, err
}

func (m *FarmRepositoryMock) UpdateFarm(farm *domain.Farm) error {
	args := m.Mock.Called(farm)
	if args[0] == nil {
		return nil
	} else {
		err := args[0].(error)
		return err
	}
}

func (m *FarmRepositoryMock) DeleteFarm(id string) error {
	args := m.Mock.Called(id)
	if args[0] == nil {
		return nil
	} else {
		err := args[0].(error)
		return err
	}
}

func (m *FarmRepositoryMock) GetAll() ([]domain.Farm, error) {
	args := m.Mock.Called()
	farms := args[0].(*[]domain.Farm)
	for _, f := range *farms {
		fmt.Println("f in test = ", f.FarmName)
	}
	if args[1] == nil {
		return *farms, nil
	}
	err := args[0].(error)
	return []domain.Farm{}, err
}
