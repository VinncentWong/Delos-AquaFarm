package usecase

import (
	"github.com/VinncentWong/Delos-AquaFarm/app/record/repository"
	"github.com/VinncentWong/Delos-AquaFarm/domain/response"
)

type IRecordUsecase interface {
	GetRecord() ([]response.RecordResponse, error)
	GetUniqueAgent(endpoint string) (int, error)
}

type RecordUsecase struct {
	repository repository.IRecordRepository
}

func NewRecordUsecase(repository repository.IRecordRepository) IRecordUsecase {
	return &RecordUsecase{
		repository: repository,
	}
}

func (u *RecordUsecase) GetRecord() ([]response.RecordResponse, error) {
	result, err := u.repository.GetRecord()
	if err != nil {
		return []response.RecordResponse{}, err
	}
	return result, nil
}

func (u *RecordUsecase) GetUniqueAgent(endpoint string) (int, error) {
	result, err := u.repository.GetUniqueAgent(endpoint)
	return result, err
}
