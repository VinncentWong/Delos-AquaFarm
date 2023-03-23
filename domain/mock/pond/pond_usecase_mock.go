package pond

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/stretchr/testify/mock"
)

/*
	CreatePond(id string, pond *domain.Pond) (domain.Pond, error)
	UpdatePond(pond *domain.Pond) error
	DeletePond(id string) error
	GetAll() ([]domain.Pond, error)
	GetPondById(id string) (domain.Pond, error)
*/

type PondUsecaseMock struct {
	Mock mock.Mock
}

func (m *PondUsecaseMock) CreatePond(id string, param *domain.Pond) (domain.Pond, error) {
	args := m.Mock.Called(id, param)
	pond := args[0].(domain.Pond)
	err := args[1].(error)
	if err != nil {
		return domain.Pond{}, err
	}
	return pond, nil
}

func (m *PondUsecaseMock) UpdatePond(param *domain.Pond) error {
	args := m.Mock.Called(param)
	err := args[0].(error)
	return err
}

func (m *PondUsecaseMock) DeletePond(id string) error {
	args := m.Mock.Called(id)
	err := args[0].(error)
	return err
}

func (m *PondUsecaseMock) GetAll() ([]domain.Pond, error) {
	args := m.Mock.Called()
	ponds := args[0].([]domain.Pond)
	err := args[1].(error)
	if err != nil {
		return []domain.Pond{}, err
	}
	return ponds, nil
}

func (m *PondUsecaseMock) GetPondById(id string) (domain.Pond, error) {
	args := m.Mock.Called(id)
	pond := args[0].(domain.Pond)
	err := args[1].(error)
	if err != nil {
		return domain.Pond{}, err
	}
	return pond, nil
}
