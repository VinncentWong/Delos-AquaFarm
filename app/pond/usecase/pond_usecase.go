package usecase

import (
	"errors"
	"fmt"

	farmRepository "github.com/VinncentWong/Delos-AquaFarm/app/farm/repository"
	pondRepository "github.com/VinncentWong/Delos-AquaFarm/app/pond/repository"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
)

type IPondUsecase interface {
	CreatePond(id string, pond *domain.Pond) (domain.Pond, error)
	UpdatePond(pond *domain.Pond) error
}

type PondUsecase struct {
	pondRepository pondRepository.IPondRepository
	farmRepository farmRepository.IFarmRepository
}

func NewPondUsecase(pondRepository pondRepository.IPondRepository, farmRepository farmRepository.IFarmRepository) IPondUsecase {
	return &PondUsecase{
		pondRepository: pondRepository,
		farmRepository: farmRepository,
	}
}

func (u *PondUsecase) CreatePond(idFarm string, pond *domain.Pond) (domain.Pond, error) {
	_, err := u.farmRepository.GetFarmById(idFarm)
	if err != nil {
		return domain.Pond{}, errors.New("farm id doesn't exist")
	}
	_, err = u.pondRepository.GetPondByName(pond.Name)
	if err == nil {
		return domain.Pond{}, errors.New("pond already exist in database, duplicate data was denied")
	}
	_, err = u.pondRepository.CreatePond(idFarm, pond)
	if err != nil {
		return domain.Pond{}, err
	}
	ponds, err := u.pondRepository.GetPondByName(pond.Name)
	if err != nil {
		return domain.Pond{}, err
	}
	return ponds, nil
}

func (u *PondUsecase) UpdatePond(pond *domain.Pond) error {
	result, err := u.pondRepository.GetPondById(fmt.Sprint(pond.ID))
	if err == nil {
		result.Name = pond.Name
		err = u.pondRepository.UpdatePond(&result)
	} else {
		err = errors.New("can't insert child without parent, farm_id must exist")
	}
	return err
}
