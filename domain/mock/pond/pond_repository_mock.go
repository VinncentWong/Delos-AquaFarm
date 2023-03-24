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
	pond := args[0].(*domain.Pond)
	if args[1] == nil {
		return *pond, nil
	}
	err := args[1].(error)
	return domain.Pond{}, err
}

func (m *PondRepositoryMock) GetPondByName(pondName string) (domain.Pond, error) {
	args := m.Mock.Called(pondName)
	pond, ok := args[0].(domain.Pond)
	if !ok {
		pond = *(args[0].(*domain.Pond))
	}
	if args[1] == nil {
		return pond, nil
	}
	return domain.Pond{}, args[1].(error)
}

func (m *PondRepositoryMock) GetPondById(id string) (domain.Pond, error) {
	args := m.Mock.Called(id)
	pond, ok := args[0].(domain.Pond)
	if !ok {
		pond = *(args[0].(*domain.Pond))
	}
	if args[1] == nil {
		return pond, nil
	}
	err := args[1].(error)
	return domain.Pond{}, err
}

func (m *PondRepositoryMock) UpdatePond(pond *domain.Pond) error {
	args := m.Mock.Called(pond)
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *PondRepositoryMock) DeletePond(id string) error {
	args := m.Mock.Called(id)
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *PondRepositoryMock) GetAll() ([]domain.Pond, error) {
	args := m.Mock.Called()
	ponds := args[0].(*[]domain.Pond)
	if args[1] == nil {
		return *ponds, nil
	}
	err := args[1].(error)
	return []domain.Pond{}, err
}
