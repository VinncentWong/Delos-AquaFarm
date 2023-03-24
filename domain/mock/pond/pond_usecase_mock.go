package pond

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/stretchr/testify/mock"
)

type PondUsecaseMock struct {
	Mock mock.Mock
}

func (m *PondUsecaseMock) CreatePond(id string, param *domain.Pond) (domain.Pond, error) {
	args := m.Mock.Called(id, param)
	pond := args[0].(domain.Pond)
	if args[1] == nil {
		return pond, nil
	}
	return domain.Pond{}, args[1].(error)
}

func (m *PondUsecaseMock) UpdatePond(param *domain.Pond) (domain.Pond, error) {
	args := m.Mock.Called(param)
	pond := args[0].(domain.Pond)
	if args[1] == nil {
		return pond, nil
	}
	return domain.Pond{}, args[1].(error)
}

func (m *PondUsecaseMock) DeletePond(id string) error {
	args := m.Mock.Called(id)
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *PondUsecaseMock) GetAll() ([]domain.Pond, error) {
	args := m.Mock.Called()
	pond := args[0].([]domain.Pond)
	if args[1] == nil {
		return pond, nil
	}
	return []domain.Pond{}, args[1].(error)
}

func (m *PondUsecaseMock) GetPondById(id string) (domain.Pond, error) {
	args := m.Mock.Called(id)
	pond := args[0].(domain.Pond)
	if args[1] == nil {
		return pond, nil
	}
	return domain.Pond{}, args[1].(error)
}
