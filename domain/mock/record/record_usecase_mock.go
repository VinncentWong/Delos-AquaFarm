package record

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain/response"
	"github.com/stretchr/testify/mock"
)

type RecordUsecaseMock struct {
	Mock mock.Mock
}

func (m *RecordUsecaseMock) GetRecord() ([]response.RecordResponse, error) {
	args := m.Mock.Called()
	records := args[0].([]response.RecordResponse)
	if args[1] == nil {
		return records, nil
	}
	return []response.RecordResponse{}, args[1].(error)
}

func (m *RecordUsecaseMock) GetUniqueAgent(endpoint string) (int, error) {
	args := m.Mock.Called(endpoint)
	numberUniqueAgent := args[0].(int)
	if args[1] == nil {
		return numberUniqueAgent, nil
	}
	return -1, args[1].(error)
}
