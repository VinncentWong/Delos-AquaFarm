package usecase

import (
	"errors"

	farmRepository "github.com/VinncentWong/Delos-AquaFarm/app/farm/repository"
	pondRepository "github.com/VinncentWong/Delos-AquaFarm/app/pond/repository"
	"github.com/VinncentWong/Delos-AquaFarm/domain"
)

type IPondUsecase interface {
	CreatePond(id string, pond *domain.Pond) (domain.Pond, error)
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
