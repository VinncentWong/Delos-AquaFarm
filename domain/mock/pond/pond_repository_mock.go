package pond

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/stretchr/testify/mock"
)

type PondRepositoryMock struct {
	Mock mock.Mock
}

func (m *PondRepositoryMock) CreatePond(id string, param *domain.Pond) (domain.Pond, error) {
	args := m.Mock.Called(id, param)
	pond := args[0].(domain.Pond)
	err := args[1].(error)
	if err != nil {
		return domain.Pond{}, err
	}
	return pond, nil
}

func (m *PondRepositoryMock) GetPondByName(pondName string) (domain.Pond, error) {
	args := m.Mock.Called(pondName)
	pond := args[0].(domain.Pond)
	err := args[1].(error)
	if err != nil {
		return domain.Pond{}, err
	}
	return pond, nil
}

func (m *PondRepositoryMock) GetPondById(id string) (domain.Pond, error) {
	args := m.Mock.Called(id)
	pond := args[0].(domain.Pond)
	err := args[1].(error)
	if err != nil {
		return domain.Pond{}, err
	}
	return pond, nil
}

func (m *PondRepositoryMock) UpdatePond(pond *domain.Pond) error {
	args := m.Mock.Called(pond)
	err := args[0].(error)
	return err
}

func (m *PondRepositoryMock) DeletePond(id string) error {
	args := m.Mock.Called(id)
	err := args[0].(error)
	return err
}

func (m *PondRepositoryMock) GetAll() ([]domain.Pond, error) {
	args := m.Mock.Called()
	ponds := args[0].([]domain.Pond)
	err := args[1].(error)
	if err != nil {
		return []domain.Pond{}, err
	}
	return ponds, nil
}
