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
	if args[1] == nil {
		return records, nil
	}
	return []response.RecordResponse{}, args[1].(error)
}

func (m *RecordRepositoryMock) GetUniqueAgent(endpoint string) (int, error) {
	args := m.Mock.Called(endpoint)
	numberUniqueAgnet := args[0].(int)
	if args[1] == nil {
		return numberUniqueAgnet, nil
	}
	return -1, args[1].(error)
}
