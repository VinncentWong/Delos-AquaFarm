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
	UpdatePond(pond *domain.Pond) (domain.Pond, error)
	DeletePond(id string) error
	GetAll() ([]domain.Pond, error)
	GetPondById(id string) (domain.Pond, error)
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

func (u *PondUsecase) UpdatePond(pond *domain.Pond) (domain.Pond, error) {
	_, err := u.pondRepository.GetPondByName(pond.Name)
	if err == nil {
		return domain.Pond{}, errors.New("pond name already exist in database, duplicate was denied")
	}
	result, err := u.pondRepository.GetPondById(fmt.Sprint(pond.ID))
	if err == nil {
		result.Name = pond.Name
		err = u.pondRepository.UpdatePond(&result)
	} else {
		err = errors.New("can't insert ponds without farm, farm_id must exist")
	}
	return result, err
}

func (u *PondUsecase) DeletePond(id string) error {
	_, err := u.pondRepository.GetPondById(id)
	if err != nil {
		return errors.New("pond doesn't exist in database")
	}
	err = u.pondRepository.DeletePond(id)
	return err
}

func (u *PondUsecase) GetAll() ([]domain.Pond, error) {
	result, err := u.pondRepository.GetAll()
	// if no entity is found, then return 404 Not Found
	if len(result) == 0 {
		return []domain.Pond{}, errors.New("ponds doesn't exist")
	}
	return result, err
}

func (u *PondUsecase) GetPondById(id string) (domain.Pond, error) {
	result, err := u.pondRepository.GetPondById(id)
	if err != nil {
		return domain.Pond{}, err
	}
	return result, nil
}
