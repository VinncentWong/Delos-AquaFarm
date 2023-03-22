package usecase

import (
	"errors"

	"github.com/VinncentWong/Delos-AquaFarm/app/farm/repository"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
)

type IFarmUsecase interface {
	CreateFarm(farm *domain.Farm) (domain.Farm, error)
	UpdateFarm(farm *domain.Farm) error
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
	_, err := u.repo.GetFarmByName(farm.FarmName)
	if err == nil {
		return domain.Farm{}, errors.New("farm already exist in database, duplicate data was denied")
	}
	result, err := u.repo.CreateFarm(farm)
	if err != nil {
		return domain.Farm{}, err
	}
	return result, nil
}

func (u *FarmUsecase) UpdateFarm(farm *domain.Farm) error {
	err := u.repo.UpdateFarm(farm)
	return err
}
