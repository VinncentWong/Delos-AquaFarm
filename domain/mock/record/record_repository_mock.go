package record

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain/response"
	"github.com/stretchr/testify/mock"
)

type RecordRepositoryMock struct {
	Mock mock.Mock
}

func (m *RecordRepositoryMock) GetRecord() ([]response.RecordResponse, error) {
	args := m.Mock.Called()
	records := args[0].([]response.RecordResponse)
	err := args[1].(error)
	if err != nil {
		return []response.RecordResponse{}, err
	}
	return records, nil
}

func (m *RecordRepositoryMock) GetUniqueAgent(endpoint string) (int, error) {
	args := m.Mock.Called(endpoint)
	numberUniqueAgnet := args[0].(int)
	err := args[1].(error)
	if err != nil {
		return -1, err
	}
	return numberUniqueAgnet, nil
}
