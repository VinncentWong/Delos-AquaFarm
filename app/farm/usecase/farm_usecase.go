package usecase

import (
	"errors"

	"github.com/VinncentWong/Delos-AquaFarm/app/farm/repository"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
)

type IFarmUsecase interface {
	CreateFarm(farm *domain.Farm) (domain.Farm, error)
	UpdateFarm(farm *domain.Farm) error
	DeleteFarm(id string) error
	GetAll() ([]domain.Farm, error)
	GetFarmById(id string) (domain.Farm, error)
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

func (u *FarmUsecase) DeleteFarm(id string) error {
	_, err := u.repo.GetFarmById(id)
	if err != nil {
		return errors.New("farm doesn't exist")
	}
	err = u.repo.DeleteFarm(id)
	return err
}

func (u *FarmUsecase) GetAll() ([]domain.Farm, error) {
	result, err := u.repo.GetAll()
	// if no entity is found then return 404 Http Not Found
	if len(result) == 0 {
		return []domain.Farm{}, errors.New("farms doesn't exist")
	}
	return result, err
}

func (u *FarmUsecase) GetFarmById(id string) (domain.Farm, error) {
	result, err := u.repo.GetFarmById(id)
	if err != nil {
		return domain.Farm{}, err
	}
	return result, nil
}
