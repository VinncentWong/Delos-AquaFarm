package usecase

import (
	"github.com/VinncentWong/Delos-AquaFarm/app/farm/repository"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
)

type IFarmUsecase interface {
	CreateFarm(farm *domain.Farm) (domain.Farm, error)
}

type FarmUsecase struct {
	repo repository.IFarmRepository
}

func NewFarmUsecase(repo repository.IFarmRepository) IFarmUsecase {
	return &FarmUsecase{
		repo: repo,
	}
}

func (u *FarmUsecase) CreateFarm(farm *domain.Farm) (domain.Farm, error) {
	result, err := u.repo.CreateFarm(farm)
	if err != nil {
		return domain.Farm{}, err
	}
	return result, nil
}
